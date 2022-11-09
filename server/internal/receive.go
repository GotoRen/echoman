package internal

import (
	"fmt"
	"net"

	"github.com/GotoRen/echoman/server/internal/logger"
	"github.com/GotoRen/echoman/server/layers"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

func (device *Device) RoutineReceiveIncoming() {
	for {
		buf := make([]byte, 1500)
		size, addr, err := device.ConnUDP.ReadFrom(buf)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("[DEBUG] Peer EndPoint:", addr)
		fmt.Println("[DEBUG] Packet size:", size)

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
			fmt.Println("[INFO] Received IPv4 packet from Real I/F", buf[:len(buf)])
			dst := buf[layers.IPv4offsetDst : layers.IPv4offsetDst+net.IPv4len]
			fmt.Println("[INFO] Peer IPv4 Address", dst)

			// if _, err = device.ConnUDP.WriteToUDP(buf[:size], &device.Peer.PeerIPv4); err != nil {
			// 	logger.LogErr("Failed to write virtual IPv4 Packet", "error", err)
			// } else {
			// 	fmt.Println("OKOKOKOKOKOKOKOKOK")
			// }

		case ipv6.Version:
			if len(buf) < ipv6.HeaderLen {
				logger.LogErr("Received IPv6 packet is too small", "error", len(buf))
				buf = nil
				continue
			}
			fmt.Println("[INFO] Received IPv6 packet from Real I/F", buf[:size])
			dst := buf[layers.IPv6offsetDst : layers.IPv6offsetDst+net.IPv6len]
			fmt.Println("[INFO] Peer IPv6 Address", dst)

		default:
			fmt.Println("ip version error")
		}
	}
}
