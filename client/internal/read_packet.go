package internal

import (
	"encoding/binary"
	"fmt"

	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/GotoRen/echoman/client/layers"
	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

func (device *Device) RoutineReceiveIncoming(buf []byte, size, sd4soc int) {
	packet := gopacket.NewPacket(buf[:size], golayers.LayerTypeEthernet, gopacket.Default)
	eh := &layers.EtherHeader{
		DstMacAddr: buf[layers.DstMACAddrOffset : layers.DstMACAddrOffset+layers.DstMacLength],
		SrcMacAddr: buf[layers.SrcMACAddrOffset : layers.SrcMACAddrOffset+layers.SrcMacLength],
		ProtoType:  binary.BigEndian.Uint16(buf[layers.Protocoloffset : layers.Protocoloffset+layers.ProtocolTypeLength]),
	}

	switch eh.ProtoType {
	case layers.EthTypeIpv4:
		if len(buf[layers.Etherlen:size]) < ipv4.HeaderLen {
			logger.LogErr("Received IPv4 packet is too small", "error", len(buf[layers.Etherlen:size]))
		}

		icmpv4Layer := packet.Layer(golayers.LayerTypeICMPv4)
		if icmpv4Layer != nil {
			icmpv4 := icmpv4Layer.(*golayers.ICMPv4)
			switch icmpv4.TypeCode.Type() {
			case golayers.ICMPv4TypeDestinationUnreachable:
				fmt.Println("[ERROR] Received ICMPv4 Destination Unreachable")
			case golayers.ICMPv4TypeEchoRequest:
				// Do nothing.
			case golayers.ICMPv4TypeEchoReply:
				fmt.Println("[INFO] Received ICMPv4 echo replay")
			default:
				logger.LogErr("Unknown ICMPv4 packet type", "error", icmpv4.TypeCode.Type())
			}
		}

		udpLayer := packet.Layer(golayers.LayerTypeUDP)
		if udpLayer != nil {
			udp := udpLayer.(*golayers.UDP)
			switch udp.DstPort.LayerType() {
			case golayers.LayerTypeDHCPv4:
				logger.LogDebug("Received DHCPv4 packet", "DHCPv4", udp.DstPort.LayerType())
			case golayers.LayerTypeDNS:
				logger.LogDebug("Received DNS A record packet", "DNS", udp.DstPort.LayerType())
			default:
				if udp.DstPort == golayers.UDPPort(device.Peer.PeerUDPPort) {
					// Echoman requst
					// Do nothing.
				} else if udp.DstPort == golayers.UDPPort(device.LocalUDPPort) {
					// Echoman response
					layers.DebugUDPMessage(buf)
				} else {
					logger.LogErr("Unknown IPv4 UDP packet type", "error", udp.DstPort)
				}
			}
		}
	case layers.EthTypeIpv6:
		if len(buf[layers.Etherlen:size]) < ipv6.HeaderLen {
			logger.LogErr("Received IPv6 packet is too small", "error", len(buf[layers.Etherlen:size]))
		}

		icmpv6Layer := packet.Layer(golayers.LayerTypeICMPv6)
		if icmpv6Layer != nil {
			icmpv6 := icmpv6Layer.(*golayers.ICMPv6)
			switch icmpv6.TypeCode.Type() {
			case golayers.ICMPv6TypeDestinationUnreachable:
				fmt.Println("[ERROR] Received ICMPv6 Destination Unreachable")
			case golayers.ICMPv6TypeEchoRequest:
				logger.LogDebug("Received ICMPv6 echo request", "ICMPv6", icmpv6.TypeCode.Type())
			case golayers.ICMPv6TypeEchoReply:
				logger.LogDebug("Received ICMPv6 echo replay", "ICMPv6", icmpv6.TypeCode.Type())
			case golayers.ICMPv6TypeRouterSolicitation:
				logger.LogDebug("Received Router Solicitation", "ICMPv6", icmpv6.TypeCode.Type())
			case golayers.ICMPv6TypeRouterAdvertisement:
				logger.LogDebug("Received Router Advertisement", "ICMPv6", icmpv6.TypeCode.Type())
			case golayers.ICMPv6TypeNeighborSolicitation:
				logger.LogDebug("Received Neighbor Solicitation", "ICMPv6", icmpv6.TypeCode.Type())
			case golayers.ICMPv6TypeNeighborAdvertisement:
				logger.LogDebug("Received Neighbor Advertisement", "ICMPv6", icmpv6.TypeCode.Type())
			case golayers.ICMPv6TypeMLDv2MulticastListenerReportMessageV2:
				logger.LogDebug("Received Multicast ListenerReport MessageV2", "ICMPv6", icmpv6.TypeCode.Type())
			default:
				logger.LogErr("Unknown ICMPv6 packet type", "error", icmpv6.TypeCode.Type())
			}
		}

		udpLayer := packet.Layer(golayers.LayerTypeUDP)
		if udpLayer != nil {
			udp := udpLayer.(*golayers.UDP)
			switch udp.DstPort.LayerType() {
			case golayers.LayerTypeDHCPv6:
				logger.LogDebug("Received DHCPv6 packet", "DHCPv6", udp.DstPort.LayerType())
			case golayers.LayerTypeDNS:
				logger.LogDebug("Received DNS AAAA record packet", "DNS", udp.DstPort.LayerType())
			default:
				logger.LogErr("Unknown IPv6 UDP packet type", "error", udp.DstPort.LayerType())
			}
		}
	case layers.EthTypeArp:
		fmt.Println("[INFO] Received ARP packet")
	default:
		logger.LogErr("Detect unknown protocol version", "error", eh.ProtoType)
	}
}
