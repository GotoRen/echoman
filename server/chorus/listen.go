package chorus

import (
	"fmt"
	"net"

	"github.com/GotoRen/echoman/server/internal/logger"
)

// bindTun binds UDPconf to the TUN interface.
func bindTun(li string, lp int) (*net.UDPConn, error) {
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP(li),
		Port: lp,
	}

	c, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func listenUDPPort(c *net.UDPConn) {
	for {
		buf := make([]byte, 1500)
		size, err := c.Read(buf) // TUN read
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("[+] Received chrous packet:", string(buf[:size]))

		rep := []byte{
			0x50, 0x6e, 0x6e, 0x67,
		}

		dstAddr := &net.UDPAddr{
			IP:   net.ParseIP("198.18.155.10"),
			Port: 30910,
		}

		// TUN write
		if _, err = c.WriteToUDP(rep, dstAddr); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("write")
		}
	}
}

func Listen(appIP string, appPort int, peerEndPoint *net.UDPAddr) {
	conn, err := bindTun(appIP, appPort)
	if err != nil {
		logger.LogErr("Failed to create connection", "error", err)
	}

	go listenUDPPort(conn)
}
