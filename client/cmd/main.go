package main

import (
	"fmt"
	"log"
	"time"

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
	t := time.NewTicker(time.Second * 1)

	intf, err := internal.SetAdapterInterface()
	if err != nil {
		logger.LogErr("Unable to get interface information", "error", err)
	}

	fd, err := internal.RecvIPv4RawSocket(intf)
	// fd, err := internal.SendIPv4RawSocket()
	if err != nil {
		logger.LogErr("Failed to open send IPv4 raw socket", "error", err)
	}

	for {
		<-t.C
		icmpv4Packet := internal.GenerateICMPv4()
		fmt.Println(icmpv4Packet)
		DebugIPv4Packet(icmpv4Packet[14:])
		DebugICMPv4Packet(icmpv4Packet[14:])
		if err := internal.SendEtherPacket(fd, icmpv4Packet); err != nil {
			log.Fatal(err)
		} else {
			logger.LogDebug("Generate ARP Response", icmpv4Packet)
		}
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
