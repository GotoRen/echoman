package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/joho/godotenv"

	"github.com/GotoRen/echoman/server/internal"
	"github.com/GotoRen/echoman/server/internal/logger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.LogErr("Error loading .env file", "error", err)
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
	// }]
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
	logger.InitZap()
	i := internal.GetServerInfo(os.Getenv("SERVER_INTERFACE"))
	fmt.Println("[INFO] Server Interface Information:", i.IfIndex)
	fmt.Println("[INFO] L3 Server IPv4Address:", i.ServerIPv4)
	fmt.Println("[INFO] L2 Server HardwareAddress:", i.ServerMAC)

	// 受信ソケット
	rv4soc, err := internal.RecvIPv4RawSocket(i.IfIndex)
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Close(rv4soc)

	// 送信ソケット
	sd4soc, err := internal.EtherSendSock(i.IfIndex)
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Close(sd4soc)

	internal.ListenServe() // linten: udp-> 30005

	for {
		buf := make([]byte, 1500)
		size, _, err := syscall.Recvfrom(rv4soc, buf, 0)
		if err != nil {
			fmt.Println("[ERROR] Failed to read packet:", err)
		}
		if size < 8 {
			fmt.Println("error")
			continue
		}

		// Packetを判断する
		internal.RoutineReceiveIncoming(buf, size, sd4soc)
	}

	// Pakcet を取得できたかどうか: recvSock
	// PacketTypeを判別
	// Unmarshal -> Generate
	// Send: sendSock (echo)
}
