package internal

import (
	"net"
	"os"

	"github.com/GotoRen/echoman/client/chorus"
	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/GotoRen/echoman/client/layers"
	golayers "github.com/google/gopacket/layers"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

// RoutineSequentialSender sends packets obtained from a virtual interface to the peer.
func (device *Device) RoutineSequentialSender() {
	for {
		buf := make([]byte, 1500)
		size, err := device.Tun.Device.Read(buf)
		if err != nil {
			logger.LogErr("Failed to receive virtual IP packet", "error", err)
		}

		if size == 0 {
			logger.LogErr("Received packet is too small", "error", size)
			continue
		}

		switch buf[0] >> 4 {
		case ipv4.Version:
			if len(buf) < ipv4.HeaderLen {
				logger.LogErr("Received IPv4 packet is too small", "error", len(buf))
				buf = nil
				continue
			}
			// fmt.Println("[DEBUG] Received IPv4 packet from TUN/TAP", buf[:size])
			// dst := buf[layers.IPv4offsetDst : layers.IPv4offsetDst+net.IPv4len]

			layers.DebugUDPMessage(buf) // Send debug

			if _, err := device.Peer.ConnUDP.WriteToUDP(buf, &device.Peer.PeerEndPoint); err != nil {
				logger.LogErr("Failed to write to real interface", "error", err)
			}

		case ipv6.Version:
			if len(buf) < ipv6.HeaderLen {
				logger.LogErr("Received IPv6 packet is too small", "error", len(buf))
				buf = nil
				continue
			}
			// fmt.Println("[DEBUG] Received IPv6 packet from TUN/TAP", buf[:size])
			// dst := buf[layers.IPv6offsetDst : layers.IPv6offsetDst+net.IPv6len]

			if _, err := device.Peer.ConnUDP.WriteToUDP(buf, &device.Peer.PeerEndPoint); err != nil {
				logger.LogErr("Failed to write virtual IPv6 Packet", "error", err)
			}

		default:
			logger.LogErr("ip version error", "error", buf[0]>>4)
		}
	}
}

// NewChorusUDPPacket sends the generated UDP packet.
func (device *Device) NewChorusUDPPacket() {
	srcVIPv4 := net.ParseIP(device.Tun.VIP).To4()
	dstVIPv4 := net.ParseIP(os.Getenv("ECHOMAN_SERVER_IPV4_TUN")).To4()
	chorusPort := golayers.UDPPort(device.ChorusPort)

	udpPacket, err := chorus.GenerateUDPRequestPacket(srcVIPv4, dstVIPv4, chorusPort)
	if err != nil {
		logger.LogErr("Failed to generate udp packet", "error", err)
	}

	// write to the tun/tap interface
	if err := SendPacket4(device.socket.sd4soc, udpPacket, dstVIPv4); err != nil {
		logger.LogErr("Failed to write to tun/tap interface", "error", err)
	}
}
