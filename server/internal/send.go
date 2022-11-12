package internal

import (
	"fmt"

	"github.com/GotoRen/echoman/server/internal/logger"
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

		fmt.Println("OK")

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
				logger.LogErr("Failed to write to real interface", "error", err)
			}

		default:
			fmt.Println("ip version error")
		}
	}
}
