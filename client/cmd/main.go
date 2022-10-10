package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/GotoRen/echoman/client/internal"
	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/GotoRen/echoman/client/layers"
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

	logger.InitZap()
	i := internal.GetServerInfo(os.Getenv("CLIENT_INTERFACE"))
	fmt.Println("[INFO] Server Interface Information:", i.IfIndex)
	fmt.Println("[INFO] L3 Server IPv4Address:", i.ServerIPv4)
	fmt.Println("[INFO] L2 Server HardwareAddress:", i.ServerMAC)

	// 送信ソケット
	fd, err := internal.EtherSendSock(i.IfIndex)
	if err != nil {
		logger.LogErr("Failed to open send IPv4 raw socket", "error", err)
	}

	// 受信ソケット
	rv4soc, err := internal.RecvIPv4RawSocket(i.IfIndex)
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Close(rv4soc)

	internal.ListenServe() // linten: udp-> 30006

	for {
		<-t.C
		// internal.GenerateICMPv4Packet(fd)
		internal.GenerateUDPPacket(fd)

		buf := make([]byte, 1500)
		size, _, err := syscall.Recvfrom(rv4soc, buf, 0)
		if err != nil {
			fmt.Println("[ERROR] Failed to read packet:", err)
		}
		if size < 8 {
			fmt.Println("error")
			continue
		}

		layers.UnmarshalEtherPacket(buf)
		layers.UnmarshalIPv4Packet(buf)
		layers.UnmarshalUDPPacket(buf)
	}
}
