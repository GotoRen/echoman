package exec

import (
	"fmt"
	"net"
	"os"

	"github.com/GotoRen/echoman/server/chorus"
	"github.com/GotoRen/echoman/server/internal"
	"github.com/GotoRen/echoman/server/internal/logger"
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

	// create the Chorus application.
	device.Chorus.PeerPort = chorus.GetChorusNetworkInfo()
	chorus.HandlePacket(net.ParseIP(device.Tun.VIP), device.Chorus.PeerPort)

	go device.RoutineSequentialReceiver()
	go device.RoutineSequentialSender()

	// make the main routine wait.
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
