package stillepost

/* Implements the Member server code */

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	pb "pg_kommsys/stillepost/stillepost_grpc"
	"sync"
	"time"
)

const (
	MEMBER_UNCONFIGURED             = 0
	MEMBER_READY                    = 1
	MEMBER_STILLEPOST_IN_PROGRESS   = 2
	MEMBER_CONFIGURING              = 3
	MEMBER_WAITING_FOR_RUNDGANGCONF = 4
)

// Events
const (
	TIMEOUT_OCCURRED      = 0
	ALL_MESSAGES_RECEIVED = 1
)

func MemberStateToString(state int) string {
	switch state {
	case MEMBER_UNCONFIGURED: // := member has not yet received nextmember and rundgang parameters (messagecount, length..)
		return "MEMBER_UNCONFIGURED"
	case MEMBER_READY: // := member has received both the rundgang parameters and the rundgang configuration (the rundgang id)
		return "MEMBER_READY"
	case MEMBER_STILLEPOST_IN_PROGRESS:
		return "MEMBER_STILLEPOST_IN_PROGRESS"
	case MEMBER_CONFIGURING: // := Member is being configured by the coordinator already
		return "MEMBER_CONFIGURING"
	case MEMBER_WAITING_FOR_RUNDGANGCONF: // := Member is waiting for the rundgang id
		return "MEMBER_WAITING_FOR_RUNDGANGCONF"
	}
	return "UNDEFINED STATE"
}

type MemberEndpoints struct {
	PeerBind      net.UDPAddr
	ClientBind    net.TCPAddr // The Endpoint this Member uses for instructions from the Coordinator (which is the Member'c client)
	PeerBindStr   string      // The Endpoint this Member uses for peer communication (e.g. where the udp Server binds)
	ClientBindStr string
}

type Member struct {
	Endpoints MemberEndpoints // Endpoints of the Member for client (Coordinator) and peer Communication

	mu         sync.Mutex  // Mutex for reading state
	state      int         // status of the ring
	NextMember net.UDPAddr // Peer endpoint of the next Member in the ring

	RundgangId       string
	RundgangIdLength uint
	MessageLen       uint32
	MessageCount     uint32
	MessageCounter   uint32 // Counts how many messages this member has received
	MemberTimeout    uint32
	startTime        int64
	endTime          int64

	IsOrigin           bool
	ForwardImmediately bool
	UdpServer          net.PacketConn

	pb.UnimplementedStillepostMemberServer
	StillepostMemberContext      context.Context
	StillepostMemberCancellation context.CancelFunc
	StillepostOriginContext      context.Context
	StillepostOriginCancellation context.CancelFunc
	ErrorChannel                 chan error
	MessageChannel               chan Message
}

type Message struct {
	sender    net.Addr
	contents  []byte
	size      int
	exception error
}

func NewMember(endpoints MemberEndpoints, state int, messageLen uint32, messageCount uint32) *Member {
	member := Member{Endpoints: endpoints, state: state, MessageLen: messageLen, MessageCount: messageCount, MessageCounter: messageCount, MemberTimeout: 0, RundgangIdLength: 16, ErrorChannel: make(chan error)}

	return &member
}

func UDPReader(conn *net.PacketConn, messageChannel chan Message, ctx context.Context, packetLength uint32) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled. Stopping UDP read.")
			return
		default:
			buffer := make([]byte, packetLength)
			n, sender, err := (*conn).ReadFrom(buffer)
			messageChannel <- Message{sender: sender, size: n, exception: err, contents: buffer}
		}
	}
}

func (m *Member) HandleMessage(message Message) {
	n, _, err, buf := message.size, message.sender, message.exception, message.contents
	if err != nil {
		log.Printf("Error reading from server %v", err)
		if m.IsOrigin {
			m.StillepostOriginCancellation()
		} else {
			m.StillepostMemberCancellation()
		}
	}

	m.mu.Lock()
	if m.state == MEMBER_READY || m.state == MEMBER_STILLEPOST_IN_PROGRESS {
		if uint(n) >= m.RundgangIdLength && m.RundgangId == string(buf)[:m.RundgangIdLength] {
			if m.MessageCounter > 1 {
				m.state = MEMBER_STILLEPOST_IN_PROGRESS
				m.MessageCounter--

				if !m.IsOrigin && m.ForwardImmediately {
					m.ForwardMessageCountTimes(buf[:m.RundgangIdLength], buf[m.RundgangIdLength:n], 1)
				}
				m.mu.Unlock()
			} else if m.MessageCounter == 1 {
				if !m.IsOrigin {
					if m.ForwardImmediately {
						m.MessageCounter--
						m.ForwardMessageCountTimes(buf[:m.RundgangIdLength], buf[m.RundgangIdLength:], 1)
					} else {
						m.MessageCounter--
						m.ForwardMessageCountTimes(buf[:m.RundgangIdLength], buf[m.RundgangIdLength:], m.MessageCount)
					}
					m.MessageCounter = m.MessageCount
					m.state = MEMBER_WAITING_FOR_RUNDGANGCONF
					m.mu.Unlock()
					m.StillepostMemberCancellation()
				} else {
					m.MessageCounter = m.MessageCount
					m.state = MEMBER_WAITING_FOR_RUNDGANGCONF
					m.mu.Unlock()
					m.StillepostOriginCancellation()
				}
			}
		} else {
			log.Printf("Got message with wrong id")
			m.mu.Unlock()
		}
	} else {
		log.Printf("Am not READY")
		m.mu.Unlock()
	}
}

/* Is the goroutine that listens on the peer bind socket for messages of the other members
 * Checks if the member is in a good state and an incoming paket has the correct ringId then counts down
 * if its not the origin it will either wait for all packets to be received and then sends all of them to the next
 * member or it will immediately forward them to the next member
 * if the last packet arrives and this member is the origin it unlocks the fin mutex so that the process can
 * continue in the stillepost function to return the coordinator the time it took for the rundgang
 * it also puts this member back in the MEMBER_WAITING_FOR_RUNDGANGCONF so that a next round of stillepost
 * can be executed
 */
func (m *Member) HandleUDPServer(channel chan error, stillepostContext context.Context) {

out:
	for {
		log.Printf("Handling udp")
		select {
		case message := <-m.MessageChannel:
			m.HandleMessage(message)
		case <-channel:
			log.Printf("Error occurred")
			break out
		case <-stillepostContext.Done():
			switch stillepostContext.Err() {
			case context.DeadlineExceeded:
				log.Println("context timeout exceeded")
				m.mu.Lock()
				m.state = MEMBER_WAITING_FOR_RUNDGANGCONF
				m.MessageCounter = m.MessageCount
				m.mu.Unlock()
				break out
			case context.Canceled:
				log.Println("All Messages received, and forwarded")
				break out
			}
		}
	}

	log.Printf("returned go routine")
}

/* Forwards a Message count times to the next member
 * Resolves it first and then writes it to the connection
 */
func (m *Member) ForwardMessageCountTimes(buf []byte, message []byte, count uint32) {
	log.Printf("Sending Pakets %d times to %s", count, m.NextMember.String())

	conn, err := net.DialUDP("udp", nil, &m.NextMember)
	if err != nil {
		log.Printf("Listen failed: %v", err)
	}
	defer conn.Close()

	for i := uint32(0); i < count; i++ {
		_, err = conn.Write(append(buf, message...))
		if err != nil {
			log.Printf("Error writing: %v", err)
		}

		log.Printf("Sent Paket No. %d", i)
	}

}

func MemberConfigToString(mc *pb.MemberConfiguration) string {
	return fmt.Sprintf(
		"IsOrigin: %t\n"+
			"Nextmember: %s\n"+
			"ForwardImmediately: %t\n"+
			"MessageCount: %d\n"+
			"MessageLen: %d\n"+
			"MemberTimeout: %d\n",
		mc.IsOrigin,
		mc.NextMember,
		mc.ClientMemberConfiguration.ForwardImmediately,
		mc.ClientMemberConfiguration.MessageCount,
		mc.ClientMemberConfiguration.MessageLen,
		mc.ClientMemberConfiguration.MemberTimeout,
	)
}

/* Configures this Member, is a grpc and is called by the Coordinator
 * Sets the ring properties
 */
func (m *Member) ConfigureMember(ctx context.Context, in *pb.MemberConfiguration) (*pb.MemberConfigurationResult, error) {
	m.mu.Lock()
	if m.state != MEMBER_READY && m.state != MEMBER_UNCONFIGURED && m.state != MEMBER_WAITING_FOR_RUNDGANGCONF {
		m.mu.Unlock()
		return &pb.MemberConfigurationResult{}, status.Error(codes.FailedPrecondition, "Member is not ready")
	}

	log.Printf("Member is being configured with\n" + MemberConfigToString(in))
	m.state = MEMBER_CONFIGURING
	m.mu.Unlock()

	nextMember, err := net.ResolveUDPAddr("udp", in.NextMember)

	if err != nil {
		m.state = MEMBER_UNCONFIGURED
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Invalid Next Member Address %v", err))
	}

	m.NextMember = *nextMember
	m.IsOrigin = in.IsOrigin

	m.ForwardImmediately = in.ClientMemberConfiguration.ForwardImmediately
	m.MessageCount = in.ClientMemberConfiguration.MessageCount
	m.MessageCounter = m.MessageCount
	m.MessageLen = in.ClientMemberConfiguration.MessageLen
	m.MemberTimeout = in.ClientMemberConfiguration.MemberTimeout

	m.state = MEMBER_WAITING_FOR_RUNDGANGCONF
	return &pb.MemberConfigurationResult{}, nil
}

func (m *Member) GetPeerEndpoint(ctx context.Context, in *pb.GetPeerEndpointParams) (*pb.GetPeerEndpointResult, error) {
	return &pb.GetPeerEndpointResult{PeerEndpoint: m.Endpoints.PeerBindStr}, nil
}

/* Configures the Rundgang, is a grpc and called by the Coordinator
 * Basically sets the rundgangId and the state to ready
 */
func (m *Member) ConfigureRundgang(ctx context.Context, in *pb.RundgangConfiguration) (*pb.RundgangConfigurationResult, error) {
	m.mu.Lock()
	if m.state != MEMBER_WAITING_FOR_RUNDGANGCONF {
		m.mu.Unlock()
		return &pb.RundgangConfigurationResult{}, status.Error(codes.FailedPrecondition, "Member is not waiting for rundgang conf")
	}

	m.RundgangId = in.RundgangId

	m.state = MEMBER_READY
	log.Printf("Configured Rundgang %s, am ready for stillepost.", in.RundgangId)

	if m.IsOrigin {
		if m.MemberTimeout > 0 {
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(m.MemberTimeout))
			m.StillepostOriginContext = ctx
			m.StillepostOriginCancellation = cancel
		} else {
			ctx, cancel := context.WithCancel(context.Background())
			m.StillepostOriginContext = ctx
			m.StillepostOriginCancellation = cancel
		}
	} else {
		if m.MemberTimeout > 0 {
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(m.MemberTimeout))
			m.StillepostMemberContext = ctx
			m.StillepostMemberCancellation = cancel
		} else {
			ctx, cancel := context.WithCancel(context.Background())
			m.StillepostMemberContext = ctx
			m.StillepostMemberCancellation = cancel
		}
	}

	m.mu.Unlock()

	if m.IsOrigin {
		go UDPReader(&m.UdpServer, m.MessageChannel, m.StillepostOriginContext, m.MessageLen+uint32(m.RundgangIdLength))
		go m.HandleUDPServer(m.ErrorChannel, m.StillepostOriginContext)
	} else {
		go UDPReader(&m.UdpServer, m.MessageChannel, m.StillepostMemberContext, m.MessageLen+uint32(m.RundgangIdLength))
		go m.HandleUDPServer(m.ErrorChannel, m.StillepostMemberContext)
	}

	return &pb.RundgangConfigurationResult{}, nil
}

/* Starts a Stillepost, is a grpc and called by the Coordinator
 * Only ever called on Stillepost Origin
 * Creates the message of specified length
 * starts the time and starts forwarding the message count times
 */
func (m *Member) Stillepost(ctx context.Context, conf *pb.StillepostCoordinatorParams) (*pb.StillepostOriginResult, error) {
	m.mu.Lock()
	if m.state != MEMBER_READY {
		log.Printf("Member is not ready")
		m.mu.Unlock()
		return &pb.StillepostOriginResult{}, status.Error(codes.FailedPrecondition, "Member is not ready")
	}

	m.state = MEMBER_STILLEPOST_IN_PROGRESS

	log.Printf("Member is Ready, Stillepost is beginning")

	message := make([]byte, m.MessageLen)

	for index := range message {
		message[index] = ([]byte("4"))[0]
	}

	m.startTime = time.Now().UnixNano()
	m.mu.Unlock()

	go m.ForwardMessageCountTimes([]byte(m.RundgangId), message, m.MessageCount)

	/* We lock fin mutex. It is already locked by default so this function
	 * will have to wait for the ServeUdp function to unlock it. That only happens if this member (the origin)
	 * has received all messages it sent out in the first place

	m.Fin.Lock()*/

	select {
	case <-m.StillepostOriginContext.Done():
		switch m.StillepostOriginContext.Err() {
		case context.DeadlineExceeded:
			m.mu.Lock()
			m.state = MEMBER_WAITING_FOR_RUNDGANGCONF
			m.MessageCounter = m.MessageCount
			m.mu.Unlock()

			return nil, status.Error(codes.Aborted, fmt.Sprintf("Timeout occured"))
		}

		m.endTime = time.Now().UnixNano()

		// return with the time it took
		return &pb.StillepostOriginResult{Time: uint64(m.endTime - m.startTime)}, nil
	}
}
