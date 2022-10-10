package internal

import (
	"log"
	"net"

	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
)

func NewICMPv4Packet() []byte {
	dstMacAddr, err := net.ParseMAC("02:42:0a:00:03:5f")
	if err != nil {
		log.Fatal(err)
	}
	srcMacAddr, err := net.ParseMAC("02:42:0a:00:03:60")
	if err != nil {
		log.Fatal(err)
	}

	ether := golayers.Ethernet{
		DstMAC:       dstMacAddr,
		SrcMAC:       srcMacAddr,
		EthernetType: golayers.EthernetTypeIPv4,
	}

	ip := golayers.IPv4{
		Version:    4,
		TOS:        0,   // Type of Service
		Length:     0,   // Total Length
		Id:         0,   // Identification
		FragOffset: 0,   // FragOffset
		TTL:        255, // Time To Live
		Protocol:   golayers.IPProtocolICMPv4,
		Checksum:   0,
		SrcIP:      net.ParseIP("10.0.3.96"),
		DstIP:      net.ParseIP("10.0.3.95"),
	}

	icmpv4 := golayers.ICMPv4{
		TypeCode: golayers.CreateICMPv4TypeCode(8, 0),
		Checksum: 0,
		Id:       10,
		Seq:      1,
	}

	// ICMPv4 Max payload length: 48 byte
	req_data := []byte{
		34, 43, 67, 99,
		0, 0, 0, 0,
		49, 129, 5, 0,
		0, 0, 0, 0,
		16, 17, 18, 19,
		21, 22, 23, 24,
		25, 26, 27, 28,
		29, 30, 31, 32,
		33, 34, 35, 36,
		37, 38, 39, 40,
		41, 42, 43, 44,
		45, 46, 47, 48,
		49, 50, 51, 52,
		53, 54, 55, 20,
	}

	options := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}

	buffer := gopacket.NewSerializeBuffer()

	if err := gopacket.SerializeLayers(buffer, options,
		&ether,
		&ip,
		&icmpv4,
		gopacket.Payload(req_data),
	); err != nil {
		logger.LogErr("Serialize error", "error", err)
		return nil
	}

	outgoingPacket := buffer.Bytes()

	return outgoingPacket
}
