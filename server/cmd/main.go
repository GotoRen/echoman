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
		DebugIPv4Packet(buf[14:num])
		DebugICMPv4Packet(buf[14:num])
		// internal.IPv4Packet(buf[:num])
		// internal.PrintPacketInfo(buf[:num])

	}
}

func DebugIPv4Packet(b []byte) {
	fmt.Println()
	fmt.Println("---------------------------------------------")
	fmt.Println("IPv4 Layer")
	fmt.Println("---------------------------------------------")
	fmt.Println("[*] Version:", b[0])
	fmt.Println("[*] Differentiated Services:", b[1])
	fmt.Println("[*] Total Length:", b[2])
	fmt.Println("[*] Identification:", b[3:5])
	fmt.Println("[*] Offset:", b[6:8])
	fmt.Println("[*] TTL:", b[9])
	fmt.Println("[*] Protocol:", b[10])
	fmt.Println("[*] Checksum:", b[11:12])
	fmt.Println("[*] SrcIP:", b[12:16])
	fmt.Println("[*] DstIP:", b[16:20])
	fmt.Println()
}

func DebugICMPv4Packet(b []byte) {
	fmt.Println("---------------------------------------------")
	fmt.Println("ICMP Layer")
	fmt.Println("---------------------------------------------")
	fmt.Println("[*] Type:", b[20])
	fmt.Println("[*] Code:", b[21])
	fmt.Println("[*] Checksum:", b[22:24])
	fmt.Println("[*] Identifier:", b[24:26])
	fmt.Println("[*] SequenceNumber:", b[26:28])
	fmt.Println("[*] TimeStamp:", b[28:36])
	fmt.Println("[*] Data:", b[36:84])
}
