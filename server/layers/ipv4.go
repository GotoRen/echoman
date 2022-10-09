package layers

import (
	"encoding/binary"
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	golayers "github.com/google/gopacket/layers"
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
