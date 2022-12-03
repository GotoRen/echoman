package internal

import (
	"net"
	"strconv"
)

// makeUDPConnection makes UDP Connection.
func makeUDPConnection(dstPort int) (connUDP *net.UDPConn, err error) {
	udpAddr, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(dstPort))
	if err != nil {
		return nil, err
	}

	connUDP, err = net.ListenUDP("udp4", udpAddr)
	if err != nil {
		return nil, err
	}

	return connUDP, nil
}

// CreateUDPConnection creates a UDP connection with the peer node.
func CreateUDPConnection(peerUDPport int) (*net.UDPConn, error) {
	connUDP, err := makeUDPConnection(peerUDPport)
	if err != nil {
		return nil, err
	}

	return connUDP, nil
}
