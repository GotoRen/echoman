package internal

import (
	"net"

	"github.com/GotoRen/echoman/server/internal/logger"
)

func portConf(lp uint16) (*net.UDPConn, error) {
	udpAddr := &net.UDPAddr{
		IP:   net.IPv4zero.To4(),
		Port: int(lp),
	}

	c, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func listenUDPPort(c *net.UDPConn) {
	buf := make([]byte, 1500)
	for {
		_, _ = c.Read(buf)
	}
}

func (device *Device) ListenServer() {
	conn, err := portConf(device.LocalUDPPort)
	if err != nil {
		logger.LogErr("Failed to create connection", "error", err)
	}

	go listenUDPPort(conn)
}
