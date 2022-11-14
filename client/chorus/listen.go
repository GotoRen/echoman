package chorus

import (
	"net"

	"github.com/GotoRen/echoman/client/internal/logger"
)

func portConf(li string, lp int) (*net.UDPConn, error) {
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP(li),
		Port: lp,
	}

	c, err := net.ListenUDP("udp", udpAddr)
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

func Listen(appIP string, appPort int) {
	conn, err := portConf(appIP, appPort)
	if err != nil {
		logger.LogErr("Failed to create connection", "error", err)
	}

	go listenUDPPort(conn)
}
