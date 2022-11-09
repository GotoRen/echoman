package internal

import (
	"fmt"
	"net"
	"strconv"

	"github.com/GotoRen/echoman/client/internal/logger"
)

// makeUDPConnection makes UDP Connection.
func makeUDPConnection(dstPort uint16) (connUDP *net.UDPConn, err error) {
	fmt.Println("[DEBUG] Peer UDP port:", strconv.Itoa(int(dstPort)))
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(int(dstPort)))
	if err != nil {
		logger.LogErr("Unable to return the address of the UDP endpoint", "error", err)
		return nil, err
	}

	connUDP, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		logger.LogErr("Unable to establish UDP connection", "error", err)
		return nil, err
	}

	return connUDP, nil
}

// CreateUDPConnection creates a UDP connection with the peer node.
func (device *Device) CreateUDPConnection() {
	connUDP, err := makeUDPConnection(device.Peer.PeerEndPoint.AddrPort().Port())
	if err != nil {
		logger.LogErr("Unable to establish UDP connection", "error", err)
	}
	device.ConnUDP = connUDP
}
