package internal

import (
	"fmt"
	"log"

	"github.com/GotoRen/echoman/client/layers"
)

func GenerateICMPv4Packet(fd int) {
	icmpv4Packet := NewICMPv4Packet()
	// layers.UnmarshalIPv4Packet(icmpv4Packet)
	// layers.UnmarshalICMPv4Packet(icmpv4Packet)
	if err := SendEtherPacket(fd, icmpv4Packet); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("[INFO] Received ICMPv4 echo request")
	}
}

func GenerateUDPPacket(fd int) {
	udpPacket := NewUDPPacket()
	// layers.UnmarshalIPv4Packet(udpPacket)
	// layers.UnmarshalICMPv4Packet(udpPacket)
	layers.DebugUDPMessage(udpPacket)
	if err := SendEtherPacket(fd, udpPacket); err != nil {
		log.Fatal(err)
	}
}
