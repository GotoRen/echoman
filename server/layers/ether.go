package layers

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
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

const (
	DstMacLength       = 6
	SrcMacLength       = 6
	ProtocolTypeLength = 2
	Etherlen           = DstMacLength + SrcMacLength + ProtocolTypeLength
)

const (
	DstMACAddrOffset = 0
	SrcMACAddrOffset = 6
	Protocoloffset   = 12
)

func UnmarshalEtherPacket(b []byte) {
	packet := gopacket.NewPacket(b, golayers.LayerTypeEthernet, gopacket.Default)
	ehterLayer := packet.Layer(golayers.LayerTypeEthernet)
	if ehterLayer != nil {
		ether, _ := ehterLayer.(*golayers.Ethernet)
		fmt.Println("---------------------------------------------")
		fmt.Println("[*] Ether - DstMAC:", ether.DstMAC.String())
		fmt.Println("[*] Ether - SrcMAC:", ether.SrcMAC.String())
		fmt.Println("[*] Ether - EthernetType:", ether.EthernetType)
		fmt.Println("---------------------------------------------")
	}
}
