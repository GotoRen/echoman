package layers

import (
	"fmt"

	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
)

func UDPPacketUnmarshal(b []byte) {
	packet := gopacket.NewPacket(b, golayers.LayerTypeEthernet, gopacket.Default)
	udpLayer := packet.Layer(golayers.LayerTypeUDP)
	if udpLayer != nil {
		udp, _ := udpLayer.(*golayers.UDP)
		fmt.Println("---------------------------------------------")
		fmt.Println("[*] UDP - SrcPort:", udp.SrcPort.String())
		fmt.Println("[*] UDP - DstPort:", udp.DstPort.String())
		fmt.Println("[*] UDP - Length:", udp.Length)
		fmt.Println("[*] UDP - Checksum:", udp.Checksum)
		fmt.Println("[*] UDP - Payload:", udp.BaseLayer.Payload)
		fmt.Println("---------------------------------------------")
	}
}
