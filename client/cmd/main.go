package main

import (
	"fmt"
	"os"
	"time"

	"github.com/GotoRen/echoman/client/exec"
	"github.com/GotoRen/echoman/client/internal"
	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/GotoRen/echoman/client/mock"
)

func main() {
	exec.Run()
	t := time.NewTicker(time.Second * 1)
	paket_type := os.Getenv("PACKET_TYPE")

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

	mock.Listen() // Listens on ports used by applications that use the overlay network

	go device.RoutineSequentialReceiver()
	go device.RoutineSequentialSender()

	for {
		// generates mock packet
		<-t.C
		switch paket_type {
		case "ICMPV4":
			// ICMPv4 for mock
		case "UDPV4":
			mock.NewUDPPacket(device)
		default:
			logger.LogErr("You chose a mode that doesn't exist", "error", paket_type)
			os.Exit(1)
		}
	}
}
