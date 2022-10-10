package internal

import (
	"log"

	"github.com/GotoRen/echoman/client/internal/logger"
)

func GenerateICMPv4Packet(fd int) {
	icmpv4Packet := NewICMPv4Packet()
	// layers.IPv4PacketUnmarshal(icmpv4Packet)
	// layers.ICMPv4PacketUnmarshal(icmpv4Packet)
	if err := SendEtherPacket(fd, icmpv4Packet); err != nil {
		log.Fatal(err)
	} else {
		logger.LogDebug("Generate ICMPv4 Packet", icmpv4Packet)
	}
}

func GenerateUDPPacket(fd int) {
	udpPacket := NewUDPPacket()
	// layers.IPv4PacketUnmarshal(udpPacket)
	// layers.UDPPacketUnmarshal(udpPacket)
	if err := SendEtherPacket(fd, udpPacket); err != nil {
		log.Fatal(err)
	}
}
