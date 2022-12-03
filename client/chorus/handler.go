package chorus

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/GotoRen/echoman/client/internal/logger"
)

// GetChorusNetworkInfo returns the Peer's TUN interface information.
func GetChorusNetworkInfo() (net.IP, int) {
	peerIP := net.ParseIP(os.Getenv("ECHOMAN_SERVER_IPV4_TUN"))

	peerPort, err := strconv.Atoi(os.Getenv("CHROUS_PORT"))
	if err != nil {
		logger.LogErr("Type conversion failed", "error", err)
	}

	return peerIP, peerPort
}

// HandlePacket handles chorus packets.
func HandlePacket(localVIP, remoteVIP net.IP, appPort int) {
	conn, err := bindTun(localVIP, appPort)
	if err != nil {
		logger.LogErr("Failed to create connection", "error", err)
	}

	go routineChorusPacketReceiver(conn)
	go routineChorusPacketSender(conn, remoteVIP, appPort)
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

// routineChorusPacketSender sends Chorus (UDP) Packet to Peer.
func routineChorusPacketSender(c *net.UDPConn, remoteIP net.IP, appPort int) {
	t := time.NewTicker(time.Second * 1)

	udpAddr := &net.UDPAddr{
		IP:   remoteIP,
		Port: appPort,
	}
	fmt.Println("[INFO] chorus remote info:", udpAddr)

	for {
		<-t.C

		request_message := []byte("Ping")

		// Write to TUN internface.
		if _, err := c.WriteToUDP(request_message, udpAddr); err != nil {
			logger.LogErr("Failed to send Chorus packet", "error", err)
		} else {
			fmt.Println("[INFO] Send chrous request packet:", string(request_message))
		}
	}
}

// routineChorusPacketReceiver receives Chorus (UDP) Packet from Peer.
func routineChorusPacketReceiver(c *net.UDPConn) {
	for {
		buf := make([]byte, 1500)

		// Read from TUN interface.
		size, _, err := c.ReadFrom(buf)
		if err != nil {
			logger.LogErr("Failed to receive Chorus packet", "error", err)
		} else {
			fmt.Println("[INFO] Receive chrous response packet:", string(buf[:size]))
		}
	}
}
