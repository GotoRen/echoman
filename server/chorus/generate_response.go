package chorus

import (
	"encoding/binary"

	"github.com/GotoRen/echoman/server/internal/logger"
	"github.com/GotoRen/echoman/server/layers"
	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
)

// GenerateUDPResponsePacket generate response packet for received UDP packets.
func GenerateUDPResponsePacket(req []byte) []byte {
	data := []byte("Pong")

	// swap source and destination information
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

	udp := golayers.UDP{
		SrcPort: golayers.UDPPort(srcPort),
		DstPort: golayers.UDPPort(dstPort),
	}

	// fmt.Println("[DEBUG] srcIPv4Addr:", srcIPv4Addr)
	// fmt.Println("[DEBUG] dstIPv4Addr:", dstIPv4Addr)
	// fmt.Println("[DEBUG] srcPort:", srcPort)
	// fmt.Println("[DEBUG] dstPort:", dstPort)

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
