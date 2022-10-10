package layers

import (
	"fmt"

	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
)

// /****************
// [*] ICMPv4 : 64 byte
// ex.)
// Type : 8
// Code : 0
// Checksum : 136 236
// Identifier : 33 186
// SequenceNumber : 0 0
// TimeStamp : 97 174 80 213 0 12 175 198
// Data : 48
// ****************/

// type ICMPv4 struct {
// 	Type           []byte // 1 byte
// 	Code           []byte // 1 byte
// 	CheckSum       []byte // 2 byte
// 	Identification []byte // 2 byte
// 	SequenceNumber []byte // 2 byte
// 	TimeStamp      []byte // 8 byte
// 	Data           []byte // 48 byte
// }

func UnmarshalICMPv4Packet(b []byte) {
	packet := gopacket.NewPacket(b, golayers.LayerTypeEthernet, gopacket.Default)
	icmpv4Layer := packet.Layer(golayers.LayerTypeICMPv4)
	if icmpv4Layer != nil {
		icmpv4, _ := icmpv4Layer.(*golayers.ICMPv4)
		fmt.Println("---------------------------------------------")
		fmt.Println("[*] ICMPv4 - TypeCode:", icmpv4.TypeCode)
		fmt.Println("[*] ICMPv4 - Chechsum:", icmpv4.Checksum)
		fmt.Println("[*] ICMPv4 - Identifier:", icmpv4.Id)
		fmt.Println("[*] ICMPv4 - SequenceNumber:", icmpv4.Seq)
		fmt.Println("---------------------------------------------")
	}
}

func DebugICMPv4Packet(b []byte) {
	fmt.Println("---------------------------------------------")
	fmt.Println("ICMP Layer")
	fmt.Println("---------------------------------------------")
	fmt.Println("[*] Type:", b[20])
	fmt.Println("[*] Code:", b[21])
	fmt.Println("[*] Checksum:", b[22:24])
	fmt.Println("[*] Identifier:", b[24:26])
	fmt.Println("[*] SequenceNumber:", b[26:28])
	fmt.Println("[*] TimeStamp:", b[28:36])
	fmt.Println("[*] Data:", b[36:84])
}
