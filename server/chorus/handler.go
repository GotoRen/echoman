package chorus

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/GotoRen/echoman/server/internal/logger"
)

// GetChorusNetworkInfo returns the Peer's TUN interface information.
func GetChorusNetworkInfo() int {
	peerPort, err := strconv.Atoi(os.Getenv("CHROUS_PORT"))
	if err != nil {
		logger.LogErr("Type conversion failed", "error", err)
	}

	return peerPort
}

// HandlePacket handles chorus packets.
func HandlePacket(localVIP net.IP, appPort int) {
	conn, err := bindTun(localVIP, appPort)
	if err != nil {
		logger.LogErr("Failed to create connection", "error", err)
	}

	go routineServe(conn)
}

// bindTun binds UDPconf to the TUN interface.
func bindTun(li net.IP, lp int) (*net.UDPConn, error) {
	udpAddr := &net.UDPAddr{
		IP:   li,
		Port: lp,
	}
	fmt.Println("[INFO] chorus local info:", udpAddr)

	c, err := net.ListenUDP("udp4", udpAddr)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// routineServe receives request packet, when it returns response packet.
func routineServe(c *net.UDPConn) {
	for {
		buf := make([]byte, 1500)

		// Read from TUN interface.
		size, addr, err := c.ReadFrom(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("[INFO] chorus remote info:", addr.(*net.UDPAddr))
		fmt.Println("[INFO] Receive chrous request packet:", string(buf[:size]))

		response_message := []byte("Pong")

		// Write to TUN internface.
		if _, err = c.WriteToUDP(response_message, addr.(*net.UDPAddr)); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("[INFO] Send chrous response packet:", string(response_message))
		}
	}
}
