package internal

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	golayers "github.com/google/gopacket/layers"
)

type EtherHeader struct {
	DstMacAddr net.HardwareAddr
	SrcMacAddr net.HardwareAddr
	ProtoType  uint16
}

const (
	EthTypeArp  uint16 = 0x0806
	EthTypeIpv4 uint16 = 0x0800
	EthTypeIpv6 uint16 = 0x86dd
)

func IPv4Packet(data []byte) {
	dstMacAddr := data[:6]
	srcMacAddr := data[6:12]
	protoType := binary.BigEndian.Uint16(data[12:14])
	eh := &EtherHeader{
		DstMacAddr: dstMacAddr,
		SrcMacAddr: srcMacAddr,
		ProtoType:  protoType,
	}
	if eh.ProtoType == EthTypeIpv4 {
		packet := gopacket.NewPacket(data, golayers.LayerTypeEthernet, gopacket.Default)
		ipLayer := packet.Layer(golayers.LayerTypeIPv4)
		if ipLayer != nil {
			fmt.Println("[*] IPv4 layer")
			ipv4, _ := ipLayer.(*layers.IPv4)
			fmt.Println("BaseLayer", ipv4.BaseLayer)
			fmt.Println("Version", ipv4.Version)
			fmt.Println("IHL", ipv4.IHL)
			fmt.Println("TOS", ipv4.TOS)
			fmt.Println("Length", ipv4.Length)
			fmt.Println("Id", ipv4.Id)
			fmt.Println("Flags", ipv4.Flags)
			fmt.Println("FragOffset", ipv4.FragOffset)
			fmt.Println("TTl", ipv4.TTL)
			fmt.Println("Protocol", ipv4.Protocol)
			fmt.Println("Checksum", ipv4.Checksum)
			fmt.Println("SrcIP", ipv4.SrcIP)
			fmt.Println("DstIP", ipv4.DstIP)
			fmt.Println("Options", ipv4.Options)
			fmt.Println("Padding", ipv4.Padding)
		}
	}
}

func PrintPacketInfo(data []byte) {
	packet := gopacket.NewPacket(data, golayers.LayerTypeEthernet, gopacket.Default)
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("[*] L2 EtherLayer")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		// Ethernet type is typically IPv4 but could be ARP or other
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
		fmt.Println()
	}

	// IPパケットへキャスト
	// Let's see if the packet is IP (even though the ether type told us)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 layer detected.")
		ip, _ := ipLayer.(*layers.IPv4)

		// IP layer variables:
		// Version (Either 4 or 6)
		// IHL (IP Header Length in 32-bit words)
		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
		// Checksum, SrcIP, DstIP
		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Protocol: ", ip.Protocol)
		fmt.Println()
	}

	// IPパケットへキャスト
	// Let's see if the packet is TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected.")
		tcp, _ := tcpLayer.(*layers.TCP)

		// TCP layer variables:
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence number: ", tcp.Seq)
		fmt.Println()
	}
}
