package layers

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
)

// /***********************************
// [*] L3(IPv4) : 20 byte
// ex:)
// Version: 69
// Diffrentiated Services: 0
// Total Length: 72
// Identification: 35 73
// Offset: 0 0
// TTL: 255
// protocol: 17
// checksum: 151 92
// srcIP: 0 0 0 0
// dstIP: 255 255 255 255
// ***********************************/

// type IPv4Header struct {
// 	VersionAndHeaderLenght []byte // 1 byte
// 	ServiceType            []byte // 1 byte
// 	TotalPacketLength      []byte // 2 byte
// 	PacketIdentification   []byte // 2 byte
// 	FlagOffset             []byte // 2 byte
// 	TTL                    []byte // 1 byte
// 	Protocol               []byte // 1 byte
// 	HeaderCheckSum         []byte // 2 byte
// 	SourceIPAddr           []byte // 4 byte
// 	DstIPAddr              []byte // 4 byte
// }

const (
	SrcIPv4Length = 4
	DstIPv4Length = 4
)

// const (
// 	SrcIPv4AddrOffset = 26
// 	DstIPv4AddrOffset = 30
// )

// IPv4 offset length.
const (
	IPv4offsetTotalLength = 2                           // IPv4offsetPayloadLength is IPv4 offset payload length.
	IPv4offsetSrc         = 12                          // IPv4offsetSrc is IPv4 offset src length.
	IPv4offsetDst         = IPv4offsetSrc + net.IPv4len // IPv4offsetDst is IPv4 offset dst length.
)

func UnmarshalIPv4Packet(b []byte) {
	packet := gopacket.NewPacket(b, golayers.LayerTypeIPv4, gopacket.Default)
	ipv4Layer := packet.Layer(golayers.LayerTypeIPv4)
	if ipv4Layer != nil {
		ipv4, _ := ipv4Layer.(*golayers.IPv4)
		fmt.Println("---------------------------------------------")
		fmt.Println("[*] IPv4 - Version:", ipv4.Version)
		fmt.Println("[*] IPv4 - IHL:", ipv4.IHL)
		fmt.Println("[*] IPv4 - TypeOfService:", ipv4.TOS)
		fmt.Println("[*] IPv4 - TotalLength:", ipv4.Length)
		fmt.Println("[*] IPv4 - Identification:", ipv4.Id)
		fmt.Println("[*] IPv4 - Flags:", ipv4.Flags)
		fmt.Println("[*] IPv4 - FragmentOffset:", ipv4.FragOffset)
		fmt.Println("[*] IPv4 - TimeToLive:", ipv4.TTL)
		fmt.Println("[*] IPv4 - Protocol:", ipv4.Protocol)
		fmt.Println("[*] IPv4 - HeaderChecksum:", ipv4.Checksum)
		fmt.Println("[*] IPv4 - SrcIP:", ipv4.SrcIP)
		fmt.Println("[*] IPv4 - DstIP:", ipv4.DstIP)
		fmt.Println("[*] IPv4 - Options:", ipv4.Options)
		fmt.Println("---------------------------------------------")
	}
}

func DebugIPv4Packet(b []byte) {
	fmt.Println()
	fmt.Println("---------------------------------------------")
	fmt.Println("IPv4 Layer")
	fmt.Println("---------------------------------------------")
	fmt.Println("[*] Version:", b[0])
	fmt.Println("[*] Differentiated Services:", b[1])
	fmt.Println("[*] Total Length:", b[2])
	fmt.Println("[*] Identification:", b[3:5])
	fmt.Println("[*] Offset:", b[6:8])
	fmt.Println("[*] TTL:", b[9])
	fmt.Println("[*] Protocol:", b[10])
	fmt.Println("[*] Checksum:", b[11:12])
	fmt.Println("[*] SrcIP:", b[12:16])
	fmt.Println("[*] DstIP:", b[16:20])
	fmt.Println()
}
