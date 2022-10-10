package internal

import (
	"log"

	"github.com/GotoRen/echoman/client/layers"
)

func GenerateICMPv4Packet(fd int) {
	icmpv4Packet := NewICMPv4Packet()
	layers.UnmarshalIPv4Packet(icmpv4Packet)
	layers.UnmarshalICMPv4Packet(icmpv4Packet)
	if err := SendEtherPacket(fd, icmpv4Packet); err != nil {
		log.Fatal(err)
	}
}

func GenerateUDPPacket(fd int) {
	udpPacket := NewUDPPacket()
	// layers.UnmarshalIPv4Packet(udpPacket)
	// layers.UnmarshalICMPv4Packet(udpPacket)
	if err := SendEtherPacket(fd, udpPacket); err != nil {
		log.Fatal(err)
	}
}
