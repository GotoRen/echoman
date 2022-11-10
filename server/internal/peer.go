package internal

import (
	"net"
	"os"
	"strconv"

	"github.com/GotoRen/echoman/server/internal/logger"
)

type Peer struct {
	// Network information
	PeerEndPoint net.UDPAddr

	// UDP connection
	ConnUDP *net.UDPConn
}

// NewPeer creates peer network information and UDP connections.
func (device *Device) NewPeer() {
	peerIPv4addr := net.ParseIP(os.Getenv("PEER_IPV4_ADDRESS"))

	peerUDPport, err := strconv.Atoi(os.Getenv("PEER_UDP_PORT"))
	if err != nil {
		logger.LogErr("Type conversion failed", "error", err)
	}

	connUDP, err := CreateUDPConnection(peerUDPport)
	if err != nil {
		logger.LogErr("Unable to establish UDP connection", "error", err)
	}

	device.Peer = &Peer{
		PeerEndPoint: net.UDPAddr{
			IP:   peerIPv4addr,
			Port: peerUDPport,
		},
		ConnUDP: connUDP,
	}
}
