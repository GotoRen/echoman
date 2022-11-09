// Package internal contains the TCP/UDP connection,
// setups TUN/TAP Device, handles DNS packets.
package internal

import (
	"fmt"
	"net"

	"github.com/GotoRen/echoman/client/internal/logger"
	"github.com/GotoRen/echoman/client/layers"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

// RoutineReadFromTUN reads buffers from TUNnel device to use go routine.
func (device *Device) RoutineReadFromTUN() {
	for {
		buf := make([]byte, 1500)
		size, err := device.Tun.Device.Read(buf)
		if err != nil {
			fmt.Println(err)
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
			fmt.Println("[INFO] Received IPv4 packet from TUN/TAP", buf[:size])
			dst := buf[layers.IPv4offsetDst : layers.IPv4offsetDst+net.IPv4len]
			fmt.Println("[INFO] Peer IPv4 Address", dst)

			// -------------------------------------------------- //
			// Add udp header
			// udp := golayers.UDP{
			// 	SrcPort: golayers.UDPPort(device.LocalUDPPort),
			// 	DstPort: golayers.UDPPort(device.Peer.PeerUDPPort),
			// }
			// if err := udp.SetNetworkLayerForChecksum(&ip); err != nil {
			// 	logger.LogErr("udp set error", "error", err)
			// 	return nil
			// }

			// options := gopacket.SerializeOptions{
			// 	ComputeChecksums: true,
			// 	FixLengths:       true,
			// }

			// buffer := gopacket.NewSerializeBuffer()

			// if err := gopacket.SerializeLayers(buffer, options,
			// 	&ether,
			// 	&ip,
			// 	&udp,
			// 	gopacket.Payload(data),
			// ); err != nil {
			// 	logger.LogErr("Serialize error", "error", err)
			// 	return nil
			// }

			// outgoingPacket := buffer.Bytes()
			// -------------------------------------------------- //

			// _, err := peer.connUDP.WriteToUDP(elem.packet, &peer.endpointv4)
			if _, err := device.ConnUDP.WriteToUDP(buf, &device.Peer.PeerEndPoint); err != nil {
				logger.LogErr("Failed to write virtual IPv4 Packet", "error", err)
			} else {
				fmt.Println("OKOKOKOKOKOKOKOKOK")
			}

		case ipv6.Version:
			if len(buf) < ipv6.HeaderLen {
				logger.LogErr("Received IPv6 packet is too small", "error", len(buf))
				buf = nil
				continue
			}
			fmt.Println("[INFO] Received IPv6 packet from TUN/TAP", buf[:size])
			dst := buf[layers.IPv6offsetDst : layers.IPv6offsetDst+net.IPv6len]
			fmt.Println("[INFO] Peer IPv6 Address", dst)

		default:
			fmt.Println("ip version error")
		}
	}
}

// TUN から吸い上げたパケットをReal I/Fに書き込む
func (peer *Peer) RoutineSequentialSender() {
	// device := peer.device
	// defer func() {
	// 	defer fmt.Printf("%v - Routine: sequential sender - stopped", peer)
	// 	peer.stopping.Done()
	// }()
	// // fmt.Printf("%v - Routine: sequential sender - started", peer)

	// for elem := range peer.queue.outbound.c {
	// 	if elem == nil {
	// 		return
	// 	}
	// 	elem.Lock()
	// 	if !peer.isRunning.Get() {
	// 		device.PutMessageBuffer(elem.buffer)
	// 		device.PutOutboundElement(elem)
	// 		continue
	// 	}

	// 	// send message and return buffer to pool
	// 	if peer.isTRS.Get() {
	// 		switch device.IPVersion {
	// 		case layers.TypeLocalIPVersion4:
	// 			_, err := peer.connUDP.WriteToUDP(elem.packet, &peer.endpointTRSv4)
	// 			device.PutMessageBuffer(elem.buffer)
	// 			device.PutOutboundElement(elem)
	// 			if err != nil {
	// 				fmt.Printf("%v - Failed to send data packet: %v", peer, err)
	// 				continue
	// 			}
	// 		case layers.TypeLocalIPVersion6, layers.TypeLocalDualStackNetwork:
	// 			_, err := peer.connUDP.WriteToUDP(elem.packet, &peer.endpointTRSv6)
	// 			device.PutMessageBuffer(elem.buffer)
	// 			device.PutOutboundElement(elem)
	// 			if err != nil {
	// 				fmt.Printf("%v - Failed to send data packet: %v", peer, err)
	// 				continue
	// 			}
	// 		}
	// 	} else {
	// 		switch device.IPVersion {
	// 		case layers.TypeLocalIPVersion4:
	// 			_, err := peer.connUDP.WriteToUDP(elem.packet, &peer.endpointv4)
	// 			device.PutMessageBuffer(elem.buffer)
	// 			device.PutOutboundElement(elem)
	// 			if err != nil {
	// 				fmt.Printf("%v - Failed to send data packet: %v", peer, err)
	// 				continue
	// 			}
	// 		case layers.TypeLocalIPVersion6, layers.TypeLocalDualStackNetwork:
	// 			_, err := peer.connUDP.WriteToUDP(elem.packet, &peer.endpointv6)
	// 			device.PutMessageBuffer(elem.buffer)
	// 			device.PutOutboundElement(elem)
	// 			if err != nil {
	// 				fmt.Printf("%v - Failed to send data packet: %v", peer, err)
	// 				continue
	// 			}
	// 		}
	// 	}
	// }
}
