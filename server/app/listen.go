package app

import (
	"net"
	"os"
	"strconv"

	"github.com/GotoRen/echoman/server/internal/logger"
)

func portConf(lp int) (*net.UDPConn, error) {
	udpAddr := &net.UDPAddr{
		IP: net.ParseIP(os.Getenv("ECHOMAN_SERVER_IPV4_TUN")),
		// IP:   net.IPv4zero.To4(),
		Port: lp,
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

func Listen() {
	appPort, err := strconv.Atoi(os.Getenv("TEST_APP_PORT"))
	if err != nil {
		logger.LogErr("Unable to get TEST_APP_PORT", "error", err)
	}

	conn, err := portConf(appPort)
	if err != nil {
		logger.LogErr("Failed to create connection", "error", err)
	}

	go listenUDPPort(conn)
}
