package chorus

import (
	"encoding/binary"

	"github.com/GotoRen/echoman/server/internal/logger"
	"github.com/GotoRen/echoman/server/layers"
	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
)

// GenerateUDPResponsePacket generates udp response packet.
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

// func NewICMPv4ReplayPacket(req []byte) []byte {
// 	// swap source and destination information
// 	dstMacAddr := net.HardwareAddr(req[layers.SrcMACAddrOffset : layers.SrcMACAddrOffset+layers.SrcMacLength])
// 	srcMacAddr := net.HardwareAddr(req[layers.DstMACAddrOffset : layers.DstMACAddrOffset+layers.DstMacLength])
// 	srcIPv4Addr := req[layers.DstIPv4AddrOffset : layers.DstIPv4AddrOffset+layers.DstIPv4Length]
// 	dstIPv4Addr := req[layers.SrcIPv4AddrOffset : layers.SrcIPv4AddrOffset+layers.SrcIPv4Length]

// 	res_data := req[layers.ICMPv4Dataoffset : layers.ICMPv4Dataoffset+layers.ICMPv4DataLength+layers.ICMPv4TimeStampLength]

// 	ether := golayers.Ethernet{
// 		DstMAC:       dstMacAddr,
// 		SrcMAC:       srcMacAddr,
// 		EthernetType: golayers.EthernetTypeIPv4,
// 	}

// 	ip := golayers.IPv4{
// 		Version:    4,
// 		TOS:        0,
// 		Length:     0,
// 		Id:         0,
// 		FragOffset: 0,
// 		TTL:        255,
// 		Protocol:   golayers.IPProtocolICMPv4,
// 		Checksum:   0,
// 		SrcIP:      srcIPv4Addr,
// 		DstIP:      dstIPv4Addr,
// 	}

// 	icmpv4 := golayers.ICMPv4{
// 		TypeCode: golayers.CreateICMPv4TypeCode(0, 0),
// 		Checksum: 0,
// 		Id:       10,
// 		Seq:      1,
// 	}

// 	options := gopacket.SerializeOptions{
// 		ComputeChecksums: true,
// 		FixLengths:       true,
// 	}

// 	buffer := gopacket.NewSerializeBuffer()

// 	if err := gopacket.SerializeLayers(buffer, options,
// 		&ether,
// 		&ip,
// 		&icmpv4,
// 		gopacket.Payload(res_data),
// 	); err != nil {
// 		logger.LogErr("Serialize error", "error", err)
// 		return nil
// 	}

// 	outgoingPacket := buffer.Bytes()

// 	return outgoingPacket
// }
