package main

import (
	"fmt"
	"os"

	"github.com/GotoRen/echoman/client/internal"
	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.LogErr("Error loading .env file", "error", err)
	}
}

func main() {
	logger.InitZap()
	// t := time.NewTicker(time.Second * 1)

	// paket_type := os.Getenv("PACKET_TYPE")

	device := internal.GetDeviceInfo(os.Getenv("LOCAL_INTERFACE"))
	fmt.Println("[INFO] Local Interface Information:", device.IfIndex)
	fmt.Println("[INFO] Local IPv4 Address:", device.LocalIPv4)
	fmt.Println("[INFO] Local Hardware Address:", device.LocalMAC)
	// fmt.Println("[INFO] Peer IPv4 Address:", device.Peer.PeerIPv4)

	device.CreateTunInterface() // create tun interface
	// device.ListenClient()        // echoman client listen
	device.CreateUDPConnection() // Create connection with Peer node (echoman server)

	// device.CreateDescriptor()
	// defer syscall.Close(device.Sd4soc)
	// defer syscall.Close(device.Rv4soc)

	for {
		// <-t.C

		// switch paket_type {
		// case "ICMPV4":
		// 	device.GenerateICMPv4Packet(device.Sd4soc)
		// case "UDPV4":
		// 	device.GenerateUDPPacket(device.Sd4soc)
		// default:
		// 	logger.LogErr("You chose a mode that doesn't exist", "error", paket_type)
		// 	os.Exit(1)
		// }

		// buf := make([]byte, 1500)
		// size, _, err := syscall.Recvfrom(device.Rv4soc, buf, 0)
		// if err != nil {
		// 	fmt.Println("[ERROR] Failed to read packet:", err)
		// }
		// if size < 8 {
		// 	fmt.Println("error")
		// 	continue
		// }

		// device.RoutineReceiveIncoming(buf, size, device.Sd4soc)
		device.RoutineReadFromTUN()
	}
}

// // execute command.
// func run() {
// }
