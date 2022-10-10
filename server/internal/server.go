package internal

import (
	"log"
	"net"
)

func portConf() (*net.UDPConn, error) {
	udpAddr := &net.UDPAddr{
		IP:   net.IPv4zero.To4(),
		Port: 3333,
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

func ServeListen() {
	conn, err := portConf()
	if err != nil {
		log.Fatal(err)
	}

	listenUDPPort(conn)
}
