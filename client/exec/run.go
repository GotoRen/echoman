package exec

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/GotoRen/echoman/client/chorus"
	"github.com/GotoRen/echoman/client/internal"
	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/joho/godotenv"
)

func Run() {
	var err error
	t := time.NewTicker(time.Second * 1)
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
	device.CreateDescriptor()
	defer device.Close()

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
		<-t.C
		device.NewChorusUDPPacket() // If you want to generate UDP packets, please uncomment here.
	}
}

func loadConf() {
	err := godotenv.Load()
	if err != nil {
		logger.LogErr("Error loading .env file", "error", err)
	}

	logger.InitZap()
}
