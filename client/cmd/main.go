package main

import (
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

	fd, err := internal.EtherSendSock(intf)
	if err != nil {
		logger.LogErr("Failed to open send IPv4 raw socket", "error", err)
	}

	for {
		<-t.C
		// internal.GenerateICMPv4Packet(fd)
		internal.GenerateUDPPacket(fd)
	}
}
