package internal

import (
	"github.com/GotoRen/echoman/server/internal/logger"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

// RoutineSequentialReceiver forwards the peer's packets obtained from
// the real interface to the virtual interface.
func (device *Device) RoutineSequentialReceiver() {
	for {
		buf := make([]byte, 1500)
		size, _, err := device.Peer.ConnUDP.ReadFrom(buf)
		if err != nil {
			logger.LogErr("Failed to receive UDP packet", "error", err)
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
			// dstIP := net.IP(buf[layers.IPv4offsetDst : layers.IPv4offsetDst+net.IPv4len]).To4()
			// dstPort := golayers.UDPPort(binary.BigEndian.Uint16(buf[layers.DstUDPPortOffset : layers.DstUDPPortOffset+layers.DstUDPLength]))

			if _, err := device.Tun.Device.Tun.Write(buf); err != nil {
				logger.LogErr("Failed to write to tun/tap interface", "error", err)
			}

		case ipv6.Version:
			if len(buf) < ipv6.HeaderLen {
				logger.LogErr("Received IPv6 packet is too small", "error", len(buf))
				buf = nil
				continue
			}
			// dst := buf[layers.IPv6offsetDst : layers.IPv6offsetDst+net.IPv6len]
			// fmt.Println("[INFO] Peer IPv6 Address", dst)

			device.ReadIPv6Packet(buf)

			if _, err := device.Tun.Device.Tun.Write(buf); err != nil {
				logger.LogErr("Failed to write to tun/tap interface", "error", err)
			}

		default:
			logger.LogErr("ip version error", "error", int(buf[0]>>4))
		}
	}
}
