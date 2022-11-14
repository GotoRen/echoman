package internal

import (
	"fmt"
	"os"

	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
)

func (device *Device) ReadIPv4Packet(b []byte) {
	packet := gopacket.NewPacket(b, golayers.LayerTypeIPv4, gopacket.Default)

	icmpv4Layer := packet.Layer(golayers.LayerTypeICMPv4)
	if icmpv4Layer != nil {
		icmpv4 := icmpv4Layer.(*golayers.ICMPv4)
		switch icmpv4.TypeCode.Type() {
		case golayers.ICMPv4TypeEchoRequest:
			fmt.Println("[DEBUG] Received ICMPv4 echo request")
		case golayers.ICMPv4TypeEchoReply:
			fmt.Println("[DEBUG] Received ICMPv4 echo replay")
		default:
			logger.LogErr("Unknown ICMPv4 packet type", "error", icmpv4.TypeCode.Type())
			os.Exit(1)
		}
	}

	udpLayer := packet.Layer(golayers.LayerTypeUDP)
	if udpLayer != nil {
		udp := udpLayer.(*golayers.UDP)
		switch udp.DstPort.LayerType() {
		case golayers.LayerTypeDHCPv4:
			fmt.Println("[DEBUG] Received DHCPv4 packet")
		case golayers.LayerTypeDNS:
			fmt.Println(("[DEBUG] Received DNS A record packet"))
		default:
			logger.LogErr("Unknown IPv4 UDP packet type", "error", udp.DstPort.LayerType())
			os.Exit(1)
		}
	}
}

func (device *Device) ReadIPv6Packet(b []byte) {
	packet := gopacket.NewPacket(b, golayers.LayerTypeIPv6, gopacket.Default)

	icmpv6Layer := packet.Layer(golayers.LayerTypeICMPv6)
	if icmpv6Layer != nil {
		icmpv6 := icmpv6Layer.(*golayers.ICMPv6)
		switch icmpv6.TypeCode.Type() {
		case golayers.ICMPv6TypeDestinationUnreachable:
			fmt.Println("[DEBUG] Received ICMPv6 Destination Unreachable")
		case golayers.ICMPv6TypeEchoRequest:
			fmt.Println("[DEBUG] Received ICMPv6 echo request")
		case golayers.ICMPv6TypeEchoReply:
			fmt.Println("[DEBUG] Received ICMPv6 echo replay")
		case golayers.ICMPv6TypeRouterSolicitation:
			fmt.Println("[DEBUG] Received Router Solicitation")
		case golayers.ICMPv6TypeRouterAdvertisement:
			fmt.Println("[DEBUG] Received Router Advertisement")
		case golayers.ICMPv6TypeNeighborSolicitation:
			fmt.Println("[DEBUG] Received Neighbor Solicitation")
		case golayers.ICMPv6TypeNeighborAdvertisement:
			fmt.Println("[DEBUG] Received Neighbor Advertisement")
		case golayers.ICMPv6TypeMLDv2MulticastListenerReportMessageV2:
			fmt.Println("[DEBUG] Received Multicast ListenerReport MessageV2")
		default:
			logger.LogErr("Unknown ICMPv6 packet type", "error", icmpv6.TypeCode.Type())
			os.Exit(1)
		}
	}

	udpLayer := packet.Layer(golayers.LayerTypeUDP)
	if udpLayer != nil {
		udp := udpLayer.(*golayers.UDP)
		switch udp.DstPort.LayerType() {
		case golayers.LayerTypeDHCPv6:
			fmt.Println(("[DEBUG] Received DHCPv6 packet"))
		case golayers.LayerTypeDNS:
			fmt.Println(("[DEBUG] Received DNS AAAA record packet"))
		default:
			logger.LogErr("Unknown IPv6 UDP packet type", "error", udp.DstPort.LayerType())
			os.Exit(1)
		}
	}
}
