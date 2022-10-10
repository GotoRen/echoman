package internal

import (
	"fmt"
	"log"
	"net"
)

func portConf(lp uint16) (*net.UDPConn, error) {
	udpAddr := &net.UDPAddr{
		IP:   net.IPv4zero.To4(),
		Port: int(lp),
	}
	fmt.Println("lp:", lp)

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

func (device *Device) ListenServe() {
	conn, err := portConf(device.LocalUDPPort)
	if err != nil {
		log.Fatal(err)
	}

	go listenUDPPort(conn)
}
