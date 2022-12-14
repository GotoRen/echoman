package chorus

import (
	"net"

	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
)

// GenerateUDPRequestPacket generates udp request packet.
func GenerateUDPRequestPacket(srcVIPv4, dstVIPv4 net.IP, chorusPort golayers.UDPPort) ([]byte, error) {
	data := []byte("Ping")

	ip := golayers.IPv4{
		Version:  4,
		Protocol: golayers.IPProtocolUDP,
		SrcIP:    srcVIPv4,
		DstIP:    dstVIPv4,
	}

	udp := golayers.UDP{
		SrcPort: chorusPort,
		DstPort: chorusPort,
	}

	// calculating checksum
	if err := udp.SetNetworkLayerForChecksum(&ip); err != nil {
		logger.LogErr("udp set error", "error", err)
		return nil, err
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
		return nil, err
	}

	// _Test
	// outgoing := []byte{
	// 	0x45, 0x00, 0x00, 0x20,
	// 	0x0d, 0x73, 0x00, 0x00,
	// 	0x00, 0x11, 0x3c, 0xf2,
	// 	0xc6, 0x12, 0xdb, 0x39,
	// 	0xc6, 0x12, 0x09, 0x0a,
	// 	0x75, 0x30, 0x75, 0x30,
	// 	0x00, 0x0c, 0xe6, 0x3b,
	// 	0x50, 0x69, 0x6e, 0x67,
	// }

	// return outgoing, nil
	return buffer.Bytes(), nil
}
