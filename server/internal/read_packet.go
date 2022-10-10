package internal

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/GotoRen/echoman/server/internal/logger"
	"github.com/GotoRen/echoman/server/layers"
	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

func RoutineReceiveIncoming(buf []byte, size, sd4soc int) {
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

		udpLayer := packet.Layer(golayers.LayerTypeUDP)
		if udpLayer != nil {
			udp := udpLayer.(*golayers.UDP)
			switch udp.DstPort.LayerType() {
			case golayers.LayerTypeDHCPv4:
				fmt.Println("[INFO] Received DHCPv4 packet")
			case golayers.LayerTypeDNS:
				fmt.Println("[INFO] Received DNS A record packet")
			default:
				if udp.DstPort.String() == layers.EchomanServerPort {
					fmt.Println("[INFO] Received Original UDP packet")
					// ここで新しくパケットを生成して返す
					/**************************************************************/
					// ここでUnmarshal => UDP構造体
					udpres := NewUDPResponsePacket(buf)
					// fmt.Printf("GetPacket: %v\n", udpres)
					// layers.UnmarshalEtherPacket(udpres)
					// layers.UnmarshalIPv4Packet(udpres)
					// layers.UnmarshalUDPPacket(udpres)
					if err := SendEtherPacket(sd4soc, udpres); err != nil {
						log.Fatal(err)
					}
					/**************************************************************/
				} else if udp.DstPort.String() == layers.EchomanClientPort {
					// Do nothing.
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
				// Do nothing.
			case golayers.ICMPv6TypeEchoRequest:
				fmt.Println("[INFO] Received ICMPv6 echo request")
			case golayers.ICMPv6TypeEchoReply:
				fmt.Println("[INFO] Received ICMPv6 echo replay")
			case golayers.ICMPv6TypeRouterSolicitation:
				fmt.Println("[INFO] Received Router Solicitation")
			case golayers.ICMPv6TypeRouterAdvertisement:
				fmt.Println("[INFO] Received Router Advertisement")
			case golayers.ICMPv6TypeNeighborSolicitation:
				fmt.Println("[INFO] Received Neighbor Solicitation")
			case golayers.ICMPv6TypeNeighborAdvertisement:
				fmt.Println("[INFO] Received Neighbor Advertisement")
			case golayers.ICMPv6TypeMLDv2MulticastListenerReportMessageV2:
				fmt.Println("[INFO] Received Multicast ListenerReport MessageV2")
			default:
				logger.LogErr("Unknown ICMPv6 packet type", "error", icmpv6.TypeCode.Type())
			}
		}

		udpLayer := packet.Layer(golayers.LayerTypeUDP)
		if udpLayer != nil {
			udp := udpLayer.(*golayers.UDP)
			switch udp.DstPort.LayerType() {
			case golayers.LayerTypeDHCPv6:
				fmt.Println("[INFO] Received DHCPv6 packet")
			case golayers.LayerTypeDNS:
				fmt.Println("[INFO] Received DNS AAAA record packet")
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