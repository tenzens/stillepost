package main

/* Stillepost Client
 * defines the two methods Configure and Stillepost which both execute the appropriate grpc with the coordinator
 */
import (
	"context"
	"fmt"
	"github.com/alecthomas/kong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "pg_kommsys/stillepost/stillepost_grpc"
	"time"
)

var cli struct {
	Configure struct {
		Coordinator        string `default:"localhost:2380" help:"The Endpoint of the Coordinator of the form <ip>:<port>"`
		Cluster            string `required:"" help:"The list of client endpoints of the cluster"`
		Timeout            uint32 `default:"10" help:"Timout for configuration in seconds, default is 10"`
		Msgbytelen         uint32 `default:"256" help:"The size of the message that is to be send around"`
		Msgcount           uint32 `default:"1" help:"The number of times the message is to be send to the next member"`
		ForwardImmediately bool   `default:"false" help:"If the members should forward their messages immediately"`
		MemberTimeout      uint32 `default:"0" help:"Time in ms that the members wait for stillepost messages before resetting themselves. If 0 they wait infinitely."`
	} `cmd:"" help:"Configures the Ring"`

	Stillepost struct {
		Coordinator string `default:"localhost:2380" help:"The Endpoint of the Coordinator of the form <ip>:<port>"`
	} `cmd:"" help:"Starts a Stillepost round with the given message settings"`

	PingSystems struct {
		Coordinator string `default:"localhost:2380" help:"The Endpoint of the Coordinator of the form <ip>:<port>"`
		Cluster     string `default:"" help:"The list of client endpoints of the cluster"`
	} `cmd:"" help:"Coordinator will send ping requests to all systems running stillepost using their textual endpoints given in the configuration."`
}

func Configure(coordinatorEndpoint, cluster string, timeout, memberTimeout, messageLen, messageCount uint32, forwardImmediately bool) {
	conn, err := grpc.Dial(coordinatorEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewStillepostCoordinatorClient(conn)

	// Configuration should not take longer than ten seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(int64(timeout)))
	defer cancel()

	_, err = c.ConfigureRing(ctx, &pb.ClientRingConfiguration{Cluster: cluster, ClientMemberConfig: &pb.ClientMemberConfiguration{
		MessageLen:         messageLen,
		MessageCount:       messageCount,
		ForwardImmediately: forwardImmediately,
		MemberTimeout:      memberTimeout,
	}})

	if err != nil {
		log.Fatalf("Could not configure ring: %v", err)
	}

	fmt.Printf("OK")
}

func Stillepost(coordinatorEndpoint string) {
	conn, err := grpc.Dial(coordinatorEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewStillepostCoordinatorClient(conn)

	// Maybe use different ctx in the future, like one with a timeout
	// However I don't know a good limit for the roundtrip time
	ctx := context.Background()

	r, err := c.Stillepost(ctx, &pb.StillepostClientParams{})
	if err != nil {
		log.Fatalf("Could not execute Stillepost: %v", err)
	}
	fmt.Printf("OK, Roundtrip took %d ns or %f ms", r.Time, float64(r.Time)/float64(1e6))
}

func PingSystems(coordinatorEndpoint, cluster string) {
	conn, err := grpc.Dial(coordinatorEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewStillepostCoordinatorClient(conn)

	// Maybe use different ctx in the future, like one with a timeout
	// However I don't know a good limit for the roundtrip time
	ctx := context.Background()

	r, err := c.PingAllSystems(ctx, &pb.ClientPingParams{Cluster: cluster})
	if err != nil {
		log.Fatalf("Could not execute ping-systems: %v", err)
	}

	fmt.Printf("Ping results")
	fmt.Printf("%s", r.Results)
}

func main() {
	ctx := kong.Parse(&cli)
	switch ctx.Command() {
	case "configure":
		Configure(cli.Configure.Coordinator, cli.Configure.Cluster, cli.Configure.Timeout, cli.Configure.MemberTimeout, cli.Configure.Msgbytelen, cli.Configure.Msgcount, cli.Configure.ForwardImmediately)
	case "stillepost":
		Stillepost(cli.Stillepost.Coordinator)
	case "ping-systems":
		PingSystems(cli.PingSystems.Coordinator, cli.PingSystems.Cluster)
	default:
		panic(ctx.Command())
	}
}
