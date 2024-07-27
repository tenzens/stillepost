package stillepost

import (
	"net"
	"strings"
)

func ParseCluster(cluster string) ([]net.TCPAddr, []string) {
	endpointsAsStrings := strings.Split(cluster, ",")
	var endpoints []net.TCPAddr
	for _, endpointAsString := range endpointsAsStrings {
		endpoint, err := net.ResolveTCPAddr("tcp", endpointAsString)

		if err != nil {
			return nil, nil
		} else {
			endpoints = append(endpoints, *endpoint)
		}
	}

	return endpoints, endpointsAsStrings
}
