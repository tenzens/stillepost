# Stillepost
## Overview
Stillepost allows you to send a udp message of specified length count times in a ring around member nodes and returns 
the time it took for a tour of the ring.

Stillepost consists of a daemon executable which can be started in either member mode or coordinator mode and a client
executable. The coordinator coordinates and manages the ring of members while the members send messages to each other in
a ring. In member mode you specifiy two ports, one on which the coordinator will talk to the member (coord endpoint) and
one on which the member will receive forwarded messages from the other members.

To start a *stillepost* meaning a tour of messages around the ring you first `configure` the ring properties: 
you give the coordinator the coord endpoints for all members of the ring, you specify the message length, how many 
messages are to be sent and if the members should forward a received message immediately or wait for all count messages
before sending count messages themselves. You can also specify how long each member should wait for all messages before
they stop waiting and reset themselves. Additionally you can specifiy a timeout for the configuration itself. 
With the given configuration you can call `stillepost` which will block on the 
client until all the messages have gone around and then return with the time it took.

You can also send echo requests (similar to a ping) to all members with `echo-members`. This effectively checks
if the physical member nodes are still running the stillepost application. For that the coordinator first resolves 
the endpoint for a member and sends that member a udp paket with an echo id, waits for a reply and repeats this for
every member.

Additionally, if you don't want to echo the members but ping the underlying systems you can do so with the 
`ping-systems` subcommand. The coordinator will use the ping tool to ping each member host and send them each 5 ping 
requests. The coordinator returns with a complete output of all ping requests. The ping tool is expected to be at `/bin/ping`.

The application draws its name and function on the children's game "Chinese whispers" which in german is called "Stille Post".

## Usage

With the source at hand in stillepost directory

Coordinator:
```bash
go run daemon/stillepost_daemon.go coordinator --client-bind=<endpoint>
```

Member:
```bash
go run daemon/stillepost_daemon.go member --coord-bind=<endpoint>  --peer-bind=<endpoint>
```

Client:

To configure the Ring:
```bash
go run client/stillepost_client.go configure \
	--coordinator=<endpoint>  \
	--cluster=<cluster> \
	--msgbytelen=<num> \
	--msgcount=<num> \
	--timeout=<num> \
	--member-timeout=<num> \
	--forward-immediately=<bool>
```
The following commands require the ring to be configured.

To execute stillepost:
```bash
go run client/stillepost_client.go stillepost [--coordinator=<coord-endpoint>]
```

To echo all members:
```bash
go run client/stillepost_client.go echo-members [--coordinator=<coord-endpoint>]
```

To ping all systems:
```bash
go run client/stillepost_client.go ping-systems [--coordinator=<coord-endpoint>]
```

An endpoint is either a `<hostname>:<port>` or `<ip>:<port>`. A cluster is a comma separated list of endpoints (the client endpoints).

Procedure is equivalent when using the binaries.

## Docker

Two docker containers are already supplied, one for the coordinator and one for the member. 
Inside the `stillepost_coordinator` directory:

```bash
$ docker build -t stillepost_coordinator .
```

Inside the `stillepost_member` directory:

```bash
$ docker build -t stillepost_member .
```

Run the coordinator:
```bash
$ docker run -e coord_bind=<endpoint> stillepost_coordinator
```

Run a member:
```bash
$ docker run -e coord_bind=<endpoint> -e peer_bind=<endpoint> stillepost_member_1
```

The coordinator listens on `0.0.0.0:2379` for calls by the client by default. The members listen on coord_bind 
`0.0.0.0:2379` and peer_bind `0.0.0.0:2380` by default. 

## Building from source

To build the executables from inside stillepost:

```bash
$ (cd client && go build -o ../build/<architecture>/stillepost_client)
$ (cd daemon && go build -o ../build/<architecture>/stillepost_daemon)
```

stillepost uses the net package which makes go output dynamically linked binaries by default. To compile to a static binary use the following command. Go will then use golangs implementations of networking stuff instead of libcs. Etcd is build the same way.

```bash
$ (cd client && CGO_ENABLED=0 go build -o ../build/<architecture>/stillepost_client)
$ (cd daemon && CGO_ENABLED=0 go build -o ../build/<architecture>/stillepost_daemon)
```

## Notes

- hasn't been tested thoroughly
- no security considerations were made
- emulating on a single linux system can have different effects for example large message sizes are not well-supported on linux. If the stillepost command doesn't return, it might just be because
the kernel started dropping the messages and not because of stillepost

## How it works

The communication between the client and the coordinator and the coordinator and the members uses grpcs. The messages
that are sent for a tour and the echo requests use udp. 

Calling the `configure` command will execute one grpc on the coordinator, who in turn sequentially issues one grpc call 
on each member for configuration. This grpc will set the next member in the ring for every member, set one member
as the origin of the ring and so on. 

Calling the `stillepost` command will first create a stillepost id that is unique for one tour of messages. This
id is communicated to every member through grpcs. Once every member has received the stillepost id the coordinator
issues a stillepost grpc on the origin of the ring. The origin of the ring then starts a timer and sends UDP packages 
as specified under the configuration to the next member in the ring. Once the messages have made a tour or the timer on 
the member runs out the coordinator receives a result for that tour of messages. 

Multiple `stillepost` commands can be issued back to back for one configuration. Having one coordinator allows for a 
central point that manages the ring. No two clients can interfere with configuring the ring, as these configurations need
to go through the coordinator. Additionally only one client can call `stillepost` on the ring at a time as the coordinator
manages access to it.
