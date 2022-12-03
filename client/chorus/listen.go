package chorus

import (
	"fmt"
	"net"
	"time"

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
	for {
		buf := make([]byte, 1500)
		size, err := c.Read(buf) // TUN read
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("[-] Received chrous packet:", string(buf[:size]))

	}
}

func Writer(c *net.UDPConn) {
	t := time.NewTicker(time.Second * 1)
	for {
		<-t.C
		// ここでWriteする
		rep := []byte{
			0x48, 0x4f, 0x47, 0x45,
		}

		dstAddr := &net.UDPAddr{
			IP:   net.ParseIP("198.18.9.10"),
			Port: 30910,
		}

		// TUN write
		if _, err := c.WriteToUDP(rep, dstAddr); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("write")
		}
	}
}

func Listen(appIP string, appPort int) {
	conn, err := portConf(appIP, appPort)
	if err != nil {
		logger.LogErr("Failed to create connection", "error", err)
	}

	go listenUDPPort(conn)
	go Writer(conn)
}
