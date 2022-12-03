package internal

import (
	"fmt"

	"github.com/GotoRen/echoman/client/internal/logger"
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
			// dst := buf[layers.IPv4offsetDst : layers.IPv4offsetDst+net.IPv4len]
			// fmt.Println("[INFO] Peer IPv4 Address", dst)

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

			if _, err := device.Tun.Device.Tun.Write(buf); err != nil {
				logger.LogErr("Failed to write to tun/tap interface", "error", err)
			}

		default:
			fmt.Println("ip version error")
		}
	}
}
