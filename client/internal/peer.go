package internal

import (
	"net"
	"os"
	"strconv"

	"github.com/GotoRen/echoman/client/internal/logger"
)

type Peer struct {
	PeerEndPoint net.UDPAddr
}

func GerPeerInfo() (peer *Peer) {
	peerIPv4addr := net.ParseIP(os.Getenv("PEER_IPV4_ADDRESS"))

	peerUDPport, err := strconv.Atoi(os.Getenv("PEER_UDP_PORT"))
	if err != nil {
		logger.LogErr("Type conversion failed", "error", err)
	}

	peer = &Peer{
		PeerEndPoint: net.UDPAddr{
			IP:   peerIPv4addr,
			Port: peerUDPport,
		},
	}

	return peer
}
