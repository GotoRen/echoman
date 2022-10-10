package internal

import (
	"encoding/binary"
	"net"

	"github.com/GotoRen/echoman/server/internal/logger"
	"github.com/GotoRen/echoman/server/layers"
	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
)

func NewUDPResponsePacket(req []byte) []byte {
	// ここで送信元と宛先の情報を入れ替える
	dstMacAddr := net.HardwareAddr(req[layers.SrcMACAddrOffset : layers.SrcMACAddrOffset+layers.SrcMacLength])
	srcMacAddr := net.HardwareAddr(req[layers.DstMACAddrOffset : layers.DstMACAddrOffset+layers.DstMacLength])
	srcIPv4Addr := req[layers.DstIPv4AddrOffset : layers.DstIPv4AddrOffset+layers.DstIPv4Length]
	dstIPv4Addr := req[layers.SrcIPv4AddrOffset : layers.SrcIPv4AddrOffset+layers.SrcIPv4Length]
	srcPort := binary.BigEndian.Uint16(req[layers.DstUDPPortOffset : layers.DstUDPPortOffset+layers.DstUDPLength])
	dstPort := binary.BigEndian.Uint16(req[layers.SrcUDPPortOffset : layers.SrcUDPPortOffset+layers.SrcUDPLength])

	ether := golayers.Ethernet{
		DstMAC:       dstMacAddr,
		SrcMAC:       srcMacAddr,
		EthernetType: golayers.EthernetTypeIPv4,
	}

	ip := golayers.IPv4{
		Version:    4,
		TOS:        0,
		Length:     0,
		Id:         0,
		FragOffset: 0,
		TTL:        255,
		Protocol:   golayers.IPProtocolUDP,
		Checksum:   0,
		SrcIP:      srcIPv4Addr,
		DstIP:      dstIPv4Addr,
	}

	udp := golayers.UDP{
		SrcPort: golayers.UDPPort(srcPort),
		DstPort: golayers.UDPPort(dstPort),
	}
	rdm_data := []byte("Pong")

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
		&ether,
		&ip,
		&udp,
		gopacket.Payload(rdm_data),
	); err != nil {
		logger.LogErr("Serialize error", "error", err)
		return nil
	}

	outgoingPacket := buffer.Bytes()

	return outgoingPacket
}
