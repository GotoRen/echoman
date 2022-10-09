package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/joho/godotenv"

	"github.com/GotoRen/echoman/server/internal"
)

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// var interf string
	// var netInterface *net.Interface
	// var err error

	// interf = "eth0"
	// // interf = "eth0"
	// netInterface, err = checkInterface(interf)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// getHardwareAddr(netInterface)

	// // interf = "en10"
	// // interf = "eth1"
	// netInterface, err = checkInterface(interf)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // fmt.Println(netInterface.Name)
	// // fmt.Println(netInterface.HardwareAddr)

	// getHardwareAddr(netInterface)
	EnvLoad()
	i := internal.GetServerInfo(os.Getenv("SERVER_INTERFACE"))
	fmt.Println("[INFO] Server Interface Information:", i.IfIndex)
	fmt.Println("[INFO] L3 Server IPv4Address:", i.ServerIPv4)
	fmt.Println("[INFO] L2 Server HardwareAddress:", i.ServerMAC)

	rv4soc, err := internal.RecvIPv4RawSocket(i.IfIndex)
	if err != nil {
		fmt.Println("[ERROR] Failed to open receive IPv6 Raw socket:", err)
	}
	defer syscall.Close(rv4soc)

	for {
		buf := make([]byte, 1500)
		num, _, err := syscall.Recvfrom(rv4soc, buf, 0)
		if err != nil {
			fmt.Println("[ERROR] Failed to read packet:", err)
		}
		fmt.Printf("GetPacket: %v\n", buf[:num])
		// internal.IPv4Packet(buf[:num])
		internal.PrintPacketInfo(buf[:num])
	}
}
