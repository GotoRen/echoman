package main

import (
	"fmt"
	"os"

	"github.com/GotoRen/echoman/server/app"
	"github.com/GotoRen/echoman/server/exec"
	"github.com/GotoRen/echoman/server/internal"
)

func main() {
	exec.Run()

	device := internal.NewDevice(os.Getenv("LOCAL_INTERFACE"))
	fmt.Println("[INFO] IP version:", device.EnvIPver)
	fmt.Println("[INFO] Local IPv4:", device.LocalIPv4)

	device.NewPeer()
	fmt.Println("[INFO] Peer IPv4:", device.Peer.PeerEndPoint.IP)
	fmt.Println("[INFO] Peer UDP port:", device.Peer.PeerEndPoint.Port)

	device.CreateTunInterface()
	fmt.Println("[INFO] TUN IPv4:", device.Tun.VIP)

	device.CreateDescriptor()
	defer device.Close()

	// mock.CreateApplicaton(device)
	app.Listen() // Listens on ports used by applications that use the overlay network

	// go device.RoutineSequentialApplicationReceiver()
	go device.RoutineSequentialReceiver()
	go device.RoutineSequentialSender()

	for {
		// make the main routine wait
	}
}
