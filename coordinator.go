package stillepost

import (
	"context"
	"fmt"
	"github.com/thanhpk/randstr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os/exec"
	pb "pg_kommsys/stillepost/stillepost_grpc"
	"strings"
	"sync"
	"time"
)

/* Coordinator States
 *
 */
const (
	RING_UNCONFIGURED           = 0
	RING_READY                  = 1
	RING_STILLEPOST_IN_PROGRESS = 2
	CONFIGURING                 = 3
	ECHOING_RING                = 4
)

func CoordinatorStateToString(state int) string {
	switch state {
	case RING_UNCONFIGURED:
		return "RING_UNCONFIGURED"
	case RING_READY:
		return "RING_READY"
	case RING_STILLEPOST_IN_PROGRESS:
		return "RING_STILLEPOST_IN_PROGRESS"
	case CONFIGURING:
		return "CONFIGURING"
	case ECHOING_RING:
		return "ECHOING_RING"
	}
	return "UNDEFINED STATE"
}

type MemberEndpointsCoord struct {
	ClientBind    net.TCPAddr // The Endpoint this Member uses for instructions from the Coordinator (which is the Member's client)
	ClientBindStr string
	PeerBindStr   string // The Endpoint this Member uses for peer communication (e.g. where the udp Server binds)
}

type CoordinatorServer struct {
	RingSize         uint
	Members          []MemberEndpointsCoord // Members in Ring containing udp and tcp addresses
	ClientBind       net.TCPAddr            // Endpoint the daemon binds to for client communication by stillepost_client
	ClientBindStr    string
	mu               sync.Mutex // Mutex for reading state
	state            int
	RundgangId       string // every Stillepost is associated with an id so that Members can discard messages of older Stillepost rounds
	RundgangIdLength uint

	pb.UnimplementedStillepostCoordinatorServer
}

func NewCoordinatorServer(state int, rundgangIdLength uint) *CoordinatorServer {
	return &CoordinatorServer{state: state, RundgangIdLength: rundgangIdLength}
}

/* Configures the Ring
 * Is a grpc that the client calls on the coordinator
 * if state of Ring is not UNCONFIGURED or exits with error
 * contacts every member and sets its next member in the ring
 *
 * if any member is in state RING_STILLEPOST_IN_PROGRESS exits with error
 */
func (cs *CoordinatorServer) ConfigureRing(ctx context.Context, in *pb.ClientRingConfiguration) (*pb.ClientRingConfigurationResult, error) {
	// check status of ring, if it's busy (e.g. already being configured, doing an echo-members), don't allow configuration
	// else configure ring
	cs.mu.Lock()
	if cs.state != RING_UNCONFIGURED && cs.state != RING_READY {
		log.Printf("Ring is busy %s", CoordinatorStateToString(cs.state))
		cs.mu.Unlock()
		return &pb.ClientRingConfigurationResult{}, status.Error(codes.FailedPrecondition, "Ring is not ready")
	} else {
		log.Printf("Configuring Ring")
		cs.state = CONFIGURING
		cs.mu.Unlock()
	}

	// if the input cluster is empty and the coordinator hasn't yet been configured with members
	// return with an error
	if in.Cluster == "" && cs.Members == nil {
		log.Printf("Given Cluster is empty and Coordinator was never configured before")
		cs.state = RING_UNCONFIGURED
		return &pb.ClientRingConfigurationResult{}, status.Error(codes.InvalidArgument, "No cluster given and never configured before")
	}

	var memberEndpoints []MemberEndpointsCoord

	// if the given cluster isn't empty, parse it and fill the above list of memberEndpoints
	// else just use the one the coordinator has been configured with before
	if in.Cluster != "" {
		memberClientEndpoints, memberClientEndpointsStr := ParseCluster(in.Cluster)

		if memberClientEndpoints == nil {
			log.Printf("Cluster isnt formatted correctly.")
			cs.state = RING_UNCONFIGURED
			return &pb.ClientRingConfigurationResult{}, status.Error(codes.InvalidArgument, fmt.Sprintf("Cluster isn't formatted correctly."))
		}

		// Asks every member for their peer address
		// Necessary to know to set up the next members of every member
		memberPeerEndpointsStr, err := cs.QueryPeerEndpoints(memberClientEndpoints)
		if err != nil {
			log.Printf("Failed to get Peer Endpoints of every member %v", err)
			cs.state = RING_UNCONFIGURED
			return &pb.ClientRingConfigurationResult{}, err
		}

		for i := 0; i < len(memberClientEndpoints); i++ {
			memberEndpoints = append(memberEndpoints, MemberEndpointsCoord{
				PeerBindStr:   memberPeerEndpointsStr[i],
				ClientBind:    memberClientEndpoints[i],
				ClientBindStr: memberClientEndpointsStr[i],
			})
		}
	} else {
		memberEndpoints = cs.Members
	}

	// Configures all the members by contacting each member and setting the rundgang properties (message count, length, ...)
	// returns a list of indices with whom configuration resulted in an error (maybe because they are already busy)
	_, errorIndices := cs.ConfigureRingMembers(memberEndpoints, in)

	if errorIndices != nil {
		//maybe unconfigure the specifc nodes, might be a TODO
		log.Printf("Not every Member could be configured")
		cs.state = RING_UNCONFIGURED
		cs.Members = nil
		cs.RingSize = uint(0)
		return &pb.ClientRingConfigurationResult{}, status.Error(codes.Unavailable, "Some Members could not be configured, bad ring state")
	}

	log.Printf("Ring is configured and ready")
	// if all goes well, the ring is ready for a stillepost
	cs.state = RING_READY
	cs.Members = memberEndpoints
	cs.RingSize = uint(len(memberEndpoints))
	return &pb.ClientRingConfigurationResult{}, nil
}

/* Configures a list of Members
 * Returns a list of indices of members who were configured successfully, and a list of indices of members who weren't
 */
func (cs *CoordinatorServer) ConfigureRingMembers(memberEndpoints []MemberEndpointsCoord, in *pb.ClientRingConfiguration) ([]uint, []uint) {
	var ringSize uint = uint(len(memberEndpoints))

	var configuredMembers []uint
	var errorMembers []uint

	for i := uint(0); i < ringSize; i++ {
		member := memberEndpoints[i]
		nextMember := memberEndpoints[(i+1)%ringSize]

		_, err := cs.ConfigureRingMember(i, in, member, nextMember)

		if err != nil {
			errorMembers = append(errorMembers, i)
			log.Printf("Got Error configuring MemberEndpoints %d %v", i, err)
		} else {
			configuredMembers = append(configuredMembers, i)
		}
	}

	return configuredMembers, errorMembers
}

/* Configures a Single Member
 * Calls the ConfigureMember grpc on the member
 * Sets the first Member in the list of members as origin
 * and sets the rest of the rundgang properties on the member (except the rundgangId)
 *
 * An already configured Ring can execute multiple stillepost rounds without reconfiguration with the same properties
 */
func (cs *CoordinatorServer) ConfigureRingMember(i uint, in *pb.ClientRingConfiguration, member MemberEndpointsCoord, nextMember MemberEndpointsCoord) (*pb.MemberConfigurationResult, error) {
	conn, err := grpc.Dial(member.ClientBind.String(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	defer conn.Close()
	c := pb.NewStillepostMemberClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var isOrigin bool = false
	if i == 0 {
		isOrigin = true
	}

	configurationResult, err := c.ConfigureMember(ctx, &pb.MemberConfiguration{
		IsOrigin:   isOrigin,
		NextMember: nextMember.PeerBindStr,
		ClientMemberConfiguration: &pb.ClientMemberConfiguration{
			MessageLen:         in.ClientMemberConfig.MessageLen,
			MessageCount:       in.ClientMemberConfig.MessageCount,
			ForwardImmediately: in.ClientMemberConfig.ForwardImmediately,
			MemberTimeout:      in.ClientMemberConfig.MemberTimeout,
		},
	})

	if err != nil {
		return nil, err
	}

	return configurationResult, nil
}

func (cs *CoordinatorServer) QueryPeerEndpoints(memberClientEndpoints []net.TCPAddr) ([]string, error) {

	var peerEndpoints []string

	for i := 0; i < len(memberClientEndpoints); i++ {
		peerEndpoint, err := cs.QueryPeerEndpoint(memberClientEndpoints[i])

		if err != nil {
			return nil, err
		}

		peerEndpoints = append(peerEndpoints, peerEndpoint)
	}

	return peerEndpoints, nil
}

func (cs *CoordinatorServer) QueryPeerEndpoint(memberClientEndpoint net.TCPAddr) (string, error) {
	conn, err := grpc.Dial(memberClientEndpoint.String(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return "", err
	}

	defer conn.Close()
	c := pb.NewStillepostMemberClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	peerEndpointResult, err := c.GetPeerEndpoint(ctx, &pb.GetPeerEndpointParams{})

	if err != nil {
		return "", err
	}

	return peerEndpointResult.PeerEndpoint, nil
}

/* Configures a specific rundgang initiated by stillepost
 * Before a stillepost can successfully execute all members need to be given a rundgangId specific for that turn of stillepost
 * This function calls configureRundgangForMember of every member and returns an error if any of the members returned with an error
 */
func (cs *CoordinatorServer) configureRundgang(rundgangId string) error {
	for _, me := range cs.Members {
		_, err := cs.configureRundgangForMember(me, rundgangId)

		if err != nil {
			return err
		}
	}

	return nil
}

/*
 * Configures the specific Rundgang for a single Member
 * Sets the rundgangId on that Member
 */
func (cs *CoordinatorServer) configureRundgangForMember(memberEndpoints MemberEndpointsCoord, rundgangId string) (*pb.RundgangConfigurationResult, error) {
	conn, err := grpc.Dial(memberEndpoints.ClientBind.String(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {

		return nil, err
	}

	defer conn.Close()
	c := pb.NewStillepostMemberClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rundgangConfigurationResult, err := c.ConfigureRundgang(ctx, &pb.RundgangConfiguration{
		RundgangId: rundgangId,
	})

	if err != nil {
		return nil, err
	}

	return rundgangConfigurationResult, nil
}

/* Executes one Stillepost Rundgang
 * Is a grpc the client calls on the coordinator
 * First configures the Rundgang meaning it distributes the rundgangId to every member of the ring
 * then it executes the stillepost grpc on the ring origin
 */
func (cs *CoordinatorServer) Stillepost(ctx context.Context, in *pb.StillepostClientParams) (*pb.StillepostResult, error) {
	cs.mu.Lock()
	if cs.state != RING_READY {
		log.Printf("Ring is not Ready %s", CoordinatorStateToString(cs.state))
		cs.mu.Unlock()
		return &pb.StillepostResult{}, status.Error(codes.FailedPrecondition, "Ring is not Ready")
	}

	cs.state = RING_STILLEPOST_IN_PROGRESS
	cs.mu.Unlock()

	err := cs.configureRundgang(randstr.Hex(int(cs.RundgangIdLength)))

	if err != nil {
		log.Printf("Rundgang couldnt be configured %v", err)
		cs.mu.Lock()
		cs.state = RING_UNCONFIGURED
		cs.mu.Unlock()
		return nil, err
	}

	conn, err := grpc.Dial(cs.Members[0].ClientBind.String(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("Could not reach ring origin %v", err)
		cs.mu.Lock()
		cs.state = RING_UNCONFIGURED
		cs.mu.Unlock()
		return nil, err
	}

	defer conn.Close()
	c := pb.NewStillepostMemberClient(conn)

	ctx = context.Background()

	stillepostResult, err := c.Stillepost(ctx, &pb.StillepostCoordinatorParams{})
	if err != nil {
		log.Printf("Error executing stillepost on rundgang origin")
		if grpcStatus, ok := status.FromError(err); ok {
			if grpcStatus.Code() == codes.Aborted {
				cs.mu.Lock()
				cs.state = RING_READY
				cs.mu.Unlock()
				return nil, err
			} else {
				cs.mu.Lock()
				cs.state = RING_UNCONFIGURED
				cs.mu.Unlock()
				return nil, err
			}
		}
	}
	cs.mu.Lock()
	cs.state = RING_READY
	cs.mu.Unlock()

	return &pb.StillepostResult{Time: stillepostResult.Time}, nil
}

func (cs *CoordinatorServer) UnconfigureMembers(members []uint) {
	//do nothing at the moment, TODO maybe think about if its necessary at all
}

func (cs *CoordinatorServer) MembersPeerEndpointsToString() []string {

	peerEndpointsString := make([]string, len(cs.Members))

	for i, member := range cs.Members {
		peerEndpointsString[i] = member.PeerBindStr
	}
	return peerEndpointsString
}

func (cs *CoordinatorServer) PingAllSystems(ctx context.Context, in *pb.ClientPingParams) (*pb.ClientPingResult, error) {
	var members []MemberEndpointsCoord

	log.Printf("Got PingSystems Request with cluster %s", in.Cluster)
	if in.Cluster == "" {
		cs.mu.Lock()
		if cs.state != RING_READY {
			log.Printf("Ring is not Ready %s", CoordinatorStateToString(cs.state))
			cs.mu.Unlock()
			return &pb.ClientPingResult{}, status.Error(codes.FailedPrecondition, "Ring is not Ready")
		}
		cs.mu.Unlock()

		members = cs.Members
	} else {
		endpointsStr := strings.Split(in.Cluster, ",")
		for _, endpointStr := range endpointsStr {
			members = append(members, MemberEndpointsCoord{
				ClientBind:    net.TCPAddr{},
				PeerBindStr:   "",
				ClientBindStr: endpointStr,
			})
		}
	}

	var results string

	for _, member := range members {
		log.Printf("Pinging with command: %s %s", "/bin/ping -c 5 -i 0.2 -q", member.PeerBindStr[:strings.Index(member.PeerBindStr, ":")])
		cmd := exec.Command("/bin/ping", "-c", "5", "-i", "0.2", "-q", member.ClientBindStr[:strings.Index(member.ClientBindStr, ":")])

		var out strings.Builder
		cmd.Stdout = &out
		err := cmd.Run()

		if err != nil {
			log.Printf("Error executing ping %v", err)
		} else {
			log.Printf("%s", out.String())

			results = results + "Pinging member: " + member.ClientBindStr + "\n" + out.String() + "\n"
		}
	}

	return &pb.ClientPingResult{
		Results: results,
	}, nil
}
