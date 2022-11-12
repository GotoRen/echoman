package internal

import (
	"github.com/GotoRen/echoman/server/app"
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
		// fmt.Println("[-] size:", size)

		if size == 0 {
			logger.LogErr("Received packet is too small", "error", size)
			continue
		}

		// fmt.Println("buf:", buf)

		// fmt.Println("+++++++++++++++++++++++++++++++++++")
		// layers.UnmarshalUDPPacket(buf)  // DEBUG
		// layers.UnmarshalIPv4Packet(buf) // DEBUG
		// fmt.Println("+++++++++++++++++++++++++++++++++++")

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

			// ============================== //
			// TODO => mock: Port 番号が 30005の時
			res := app.GenerateUDPResponsePacket(buf)
			if _, err := device.Peer.ConnUDP.WriteToUDP(res, &device.Peer.PeerEndPoint); err != nil {
				logger.LogErr("Failed to write to real interface", "error", err)
			}
			// ============================== //

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
			// os.Exit(1)
		}
	}
}

// func (device *Device) RoutineSequentialApplicationReceiver() {
// 	for {
// 		buf := make([]byte, 1500)
// 		size, _, err := syscall.Recvfrom(device.Rv4soc, buf, 0)
// 		// size, err := syscall.Read(device.Rv4soc, buf)
// 		if err != nil {
// 			logger.LogErr("Failed to read packet", "error", err)
// 		}
// 		// fmt.Println("[-] size:", size)

// 		if size == 0 {
// 			logger.LogErr("Received packet is too small", "error", size)
// 			continue
// 		}

// 		// fmt.Println(buf)

// 		// if size == 1500 {
// 		// 	continue
// 		// }

// 		// fmt.Println("RAW", size, buf)
// 		// layers.DebugUDPMessage(buf) // DEBUG
// 		// layers.DebugIPv4Message(buf)
// 		// layers.UnmarshalIPv4Packet(buf)
// 		// layers.UnmarshalUDPPacket(buf)
// 		// fmt.Println("OK")
// 		// fmt.Println("buf2:", buf)
// 		// layers.DebugUDPMessage(buf)

// 		if err != nil {
// 			logger.LogErr("Failed to receive UDP packet", "error", err)
// 		}

// 		fmt.Println("buf2:", buf)

// 		// layers.UnmarshalUDPPacket(buf)
// 		// fmt.Println("+++++++++++++++++++++++++++")

// 		// ============================================================== //
// 		// ここで情報を引っくり返す
// 		// _ = app.GenerateUDPResponsePacket(buf)
// 		// fmt.Println("+++++++++++++++++++++++++++")
// 		// // layers.UnmarshalIPv4Packet(udpres)
// 		// // layers.UnmarshalUDPPacket(udpres)
// 		// fmt.Println("+++++++++++++++++++++++++++")

// 		// ============================================================== //

// 		// switch buf[0] >> 4 {
// 		// case ipv4.Version:
// 		// 	if len(buf) < ipv4.HeaderLen {
// 		// 		logger.LogErr("Received IPv4 packet is too small", "error", len(buf))
// 		// 		buf = nil
// 		// 		continue
// 		// 	}
// 		// 	// dst := buf[layers.IPv4offsetDst : layers.IPv4offsetDst+net.IPv4len]
// 		// 	// fmt.Println("[INFO] Peer IPv4 Address", dst)

// 		// 	if _, err := device.Tun.Device.Tun.Write(buf); err != nil {
// 		// 		logger.LogErr("Failed to write to tun/tap interface", "error", err)
// 		// 	}

// 		// case ipv6.Version:
// 		// 	if len(buf) < ipv6.HeaderLen {
// 		// 		logger.LogErr("Received IPv6 packet is too small", "error", len(buf))
// 		// 		buf = nil
// 		// 		continue
// 		// 	}
// 		// 	// dst := buf[layers.IPv6offsetDst : layers.IPv6offsetDst+net.IPv6len]
// 		// 	// fmt.Println("[INFO] Peer IPv6 Address", dst)

// 		// 	device.ReadIPv6Packet(buf)

// 		// 	if _, err := device.Tun.Device.Tun.Write(buf); err != nil {
// 		// 		logger.LogErr("Failed to write to tun/tap interface", "error", err)
// 		// 	}

// 		// default:
// 		// 	logger.LogErr("ip version error", "error", int(buf[0]>>4))
// 		// 	// os.Exit(1)
// 		// }
// 		buf = nil
// 	}
// }
