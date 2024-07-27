package main

/* Stillepost Daemon
 * Works in two modes: Member mode and Coordinator mode
 * Member Mode starts a Server as grpc StillepostMember and an Udp Server on which the messages of a Stillepost
 * are sent and received
 *
 * Coordinator Mode starts a Server as grpc StillepostCoordinator
 */
import (
	"errors"
	"fmt"
	"github.com/alecthomas/kong"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"pg_kommsys/stillepost"
	pb "pg_kommsys/stillepost/stillepost_grpc"
	"syscall"
)

var cli struct {
	Coordinator struct {
		ClientBind string `default:"localhost:2379" required:""`
	} `cmd:""`

	Member struct {
		CoordBind string `required:"" default:"localhost:2378"`
		PeerBind  string `required:"" default:"localhost:2377"`
	} `cmd:""`
}

func Coordinate(clientBind string) {
	server := stillepost.NewCoordinatorServer(stillepost.RING_UNCONFIGURED, 16)

	if clientBindEndpoint, err := net.ResolveTCPAddr("tcp", clientBind); err == nil {
		server.ClientBind = *clientBindEndpoint
		server.ClientBindStr = clientBind
	} else {
		log.Fatalf("Daemon cannot listen on Endpoint %s, format isn't of <ip>:<port>", clientBind)
	}

	lis, err := net.ListenTCP("tcp", &server.ClientBind)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStillepostCoordinatorServer(s, server)
	log.Printf("server listening at %v", lis.Addr())

	stopChannel := make(chan os.Signal)
	errorChannel := make(chan error)

	signal.Notify(stopChannel, syscall.SIGTERM, syscall.SIGINT)

	defer func() {
		s.GracefulStop()
	}()

	go Serve(errorChannel, lis, s)

	select {
	case err := <-errorChannel:
		fmt.Println(err)
	case <-stopChannel:
		fmt.Println("Server received Signal, stopping Server")
	}

}

func Serve(errorChannel chan error, lis net.Listener, server *grpc.Server) {
	if err := server.Serve(lis); err != nil {
		errorChannel <- errors.New(fmt.Sprintf("failed to serve: %v", err))
	}
}

// clientBind is the endpoint the Member listens on for instructions from the Coordinator
func Member(clientBind string, memberBind string) {
	clientBindEndpoint, err := net.ResolveTCPAddr("tcp", clientBind)
	if err != nil {
		log.Fatalf("Daemon cannot listen on Endpoint %s, %v", clientBind, err)
	}

	memberBindEndpoint, err := net.ResolveUDPAddr("udp", memberBind)
	if err != nil {
		log.Fatalf("Daemon cannot listen on Endpoint %s %v", memberBind, err)
	}

	lis, err := net.ListenTCP("tcp", clientBindEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	memberServer := stillepost.NewMember(stillepost.MemberEndpoints{
		PeerBind:      *memberBindEndpoint,
		PeerBindStr:   memberBind,
		ClientBind:    *clientBindEndpoint,
		ClientBindStr: clientBind,
	}, stillepost.MEMBER_UNCONFIGURED, 256, 1)

	memberServer.UdpServer, err = net.ListenUDP("udp", &memberServer.Endpoints.PeerBind)
	if err != nil {
		memberServer.UdpServer.Close()
		log.Fatalf("failed to listen for udp: %v", err)
	}

	stopChannel := make(chan os.Signal)

	signal.Notify(stopChannel, syscall.SIGTERM, syscall.SIGINT)
	s := grpc.NewServer()
	pb.RegisterStillepostMemberServer(s, memberServer)
	log.Printf("server listening at %v", lis.Addr())

	defer func() {
		if memberServer.StillepostMemberCancellation != nil {
			memberServer.StillepostMemberCancellation()
		}
		if memberServer.StillepostOriginCancellation != nil {
			memberServer.StillepostOriginCancellation()
		}
		s.GracefulStop()
		memberServer.UdpServer.Close()
	}()

	memberServer.MessageChannel = make(chan stillepost.Message)
	go Serve(memberServer.ErrorChannel, lis, s)

	select {
	case err := <-memberServer.ErrorChannel:
		fmt.Println(err)
	case <-stopChannel:
		fmt.Println("Server received Signal, stopping Server")
	}

}

func main() {
	ctx := kong.Parse(&cli)
	switch ctx.Command() {
	case "coordinator":
		Coordinate(cli.Coordinator.ClientBind)
	case "member":
		Member(cli.Member.CoordBind, cli.Member.PeerBind)
	default:
		panic(ctx.Command())
	}
}
