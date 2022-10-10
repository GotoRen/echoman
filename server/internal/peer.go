package internal

import (
	"net"
	"os"
	"strconv"

	"github.com/GotoRen/echoman/server/internal/logger"
)

type Peer struct {
	PeerIPv4    net.IP
	PeerMAC     net.HardwareAddr
	PeerUDPPort uint16
}

func GerPeerInfo() (peer *Peer) {
	peerIPv4addr := net.ParseIP(os.Getenv("PEER_IPV4_ADDRESS"))

	peerMACaddr, err := net.ParseMAC(os.Getenv("PEER_MAC_ADDRESS"))
	if err != nil {
		logger.LogErr("MAC address parse error", "error", err)
	}

	peerUDPport, err := strconv.Atoi(os.Getenv("PEER_UDP_PORT"))
	if err != nil {
		logger.LogErr("Type conversion failed", "error", err)
	}

	peer = &Peer{
		PeerIPv4:    peerIPv4addr,
		PeerMAC:     peerMACaddr,
		PeerUDPPort: uint16(peerUDPport),
	}

	return peer
}
