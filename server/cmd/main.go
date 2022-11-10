package main

import (
	"fmt"
	"os"

	"github.com/GotoRen/echoman/server/exec"
	"github.com/GotoRen/echoman/server/internal"
)

func main() {
	exec.Run()

	device := internal.NewDevice(os.Getenv("LOCAL_INTERFACE"))
	fmt.Println("[INFO] IP version:", device.EnvIPver)
	fmt.Println("[INFO] Local IPv4:", device.LocalIPv4)

	device.NewPeer()

	device.CreateTunInterface()
	fmt.Println("[INFO] Peer IPv4:", device.Peer.PeerEndPoint.IP)
	fmt.Println("[INFO] Peer UDP port:", device.Peer.PeerEndPoint.Port)

	go device.RoutineSequentialReceiver()
	go device.RoutineSequentialSender()

	for {
		// make the main routine wait
	}
}
