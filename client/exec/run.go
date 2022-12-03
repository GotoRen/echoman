package exec

import (
	"fmt"
	"net"
	"os"

	"github.com/GotoRen/echoman/client/chorus"
	"github.com/GotoRen/echoman/client/internal"
	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/joho/godotenv"
)

func Run() {
	loadConf()

	device := internal.NewDevice(os.Getenv("LOCAL_INTERFACE"))
	fmt.Println("[INFO] IP version:", device.EnvIPver)
	fmt.Println("[INFO] Local IPv4:", device.LocalIPv4)

	device.NewPeer()
	fmt.Println("[INFO] Peer IPv4:", device.Peer.PeerEndPoint.IP)

	device.CreateTunInterface()
	fmt.Println("[INFO] TUN IPv4:", device.Tun.VIP)

	// create the Chorus Application.
	device.Chorus.PeerIP, device.Chorus.PeerPort = chorus.GetChorusNetworkInfo()
	chorus.HandlePacket(net.ParseIP(device.Tun.VIP), device.Chorus.PeerIP, device.Chorus.PeerPort)

	go device.RoutineSequentialReceiver()
	go device.RoutineSequentialSender()

	// continue the main routine.
	for {
	}
}

func loadConf() {
	err := godotenv.Load()
	if err != nil {
		logger.LogErr("Error loading .env file", "error", err)
	}

	logger.InitZap()
}
