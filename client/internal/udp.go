package internal

import (
	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/google/gopacket"
	golayers "github.com/google/gopacket/layers"
)

func (device *Device) NewUDPPacket() []byte {
	// dstMacAddr, err := net.ParseMAC("02:42:0a:00:03:5f")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// srcMacAddr, err := net.ParseMAC("02:42:0a:00:03:60")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	ether := golayers.Ethernet{
		DstMAC:       device.Peer.PeerMAC,
		SrcMAC:       device.LocalMAC,
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
		SrcIP:      device.LocalIPv4,
		DstIP:      device.Peer.PeerIPv4,
	}

	udp := golayers.UDP{
		SrcPort: golayers.UDPPort(device.LocalUDPPort),
		DstPort: golayers.UDPPort(device.Peer.PeerUDPPort),
	}
	data := []byte("Ping")

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
		gopacket.Payload(data),
	); err != nil {
		logger.LogErr("Serialize error", "error", err)
		return nil
	}

	outgoingPacket := buffer.Bytes()

	return outgoingPacket
}
