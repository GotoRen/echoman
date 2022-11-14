package exec

import (
	"fmt"
	"os"
	"strconv"

	"github.com/GotoRen/echoman/server/chorus"
	"github.com/GotoRen/echoman/server/internal"
	"github.com/GotoRen/echoman/server/internal/logger"
	"github.com/joho/godotenv"
)

func Run() {
	var err error
	loadConf()

	device := internal.NewDevice(os.Getenv("LOCAL_INTERFACE"))
	fmt.Println("[INFO] IP version:", device.EnvIPver)
	fmt.Println("[INFO] Local IPv4:", device.LocalIPv4)

	device.NewPeer()
	fmt.Println("[INFO] Peer IPv4:", device.Peer.PeerEndPoint.IP)

	device.CreateTunInterface()
	fmt.Println("[INFO] TUN IPv4:", device.Tun.VIP)

	// Create Application
	//=======================================================================//
	device.ChorusPort, err = strconv.Atoi(os.Getenv("CHROUS_PORT"))
	if err != nil {
		fmt.Println(err)
	}

	// Listens on ports used by applications (chorus) that use the overlay network.
	chorus.Listen(device.Tun.VIP, device.ChorusPort)
	//=======================================================================//

	go device.RoutineSequentialReceiver()
	go device.RoutineSequentialSender()

	for {
		// make the main routine wait
	}
}

func loadConf() {
	err := godotenv.Load()
	if err != nil {
		logger.LogErr("Error loading .env file", "error", err)
	}

	logger.InitZap()
}
