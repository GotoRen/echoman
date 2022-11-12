package app

import (
	"encoding/binary"
	"fmt"

	"github.com/GotoRen/echoman/server/internal/logger"
	"github.com/GotoRen/echoman/server/layers"
	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
)

func GenerateUDPResponsePacket(req []byte) []byte {
	data := []byte("Pong")

	// ここで送信元と宛先の情報を入れ替える
	srcIPv4Addr := req[layers.IPv4offsetDst : layers.IPv4offsetDst+layers.DstIPv4Length]
	dstIPv4Addr := req[layers.IPv4offsetSrc : layers.IPv4offsetSrc+layers.SrcIPv4Length]
	srcPort := binary.BigEndian.Uint16(req[layers.DstUDPPortOffset : layers.DstUDPPortOffset+layers.DstUDPLength])
	dstPort := binary.BigEndian.Uint16(req[layers.SrcUDPPortOffset : layers.SrcUDPPortOffset+layers.SrcUDPLength])

	ip := golayers.IPv4{
		Version:  4,
		Protocol: golayers.IPProtocolUDP,
		SrcIP:    srcIPv4Addr,
		DstIP:    dstIPv4Addr,
	}

	fmt.Println("srcIPv4Addr:", srcIPv4Addr)
	fmt.Println("dstIPv4Addr:", dstIPv4Addr)
	fmt.Println("srcPort:", srcPort)
	fmt.Println("dstPort:", dstPort)

	udp := golayers.UDP{
		SrcPort: golayers.UDPPort(srcPort),
		DstPort: golayers.UDPPort(dstPort),
	}

	// calculating checksum
	if err := udp.SetNetworkLayerForChecksum(&ip); err != nil {
		logger.LogErr("udp set error", "error", err)
		return nil
	}

	options := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}

	buffer := gopacket.NewSerializeBuffer()

	if err := gopacket.SerializeLayers(buffer, options,
		&ip,
		&udp,
		gopacket.Payload(data),
	); err != nil {
		logger.LogErr("Serialize error", "error", err)
		return nil
	}

	outgoingPacket := buffer.Bytes()

	return outgoingPacket
}
