package internal

import (
	"encoding/binary"
	"net"

	"github.com/GotoRen/echoman/server/chorus"
	"github.com/GotoRen/echoman/server/internal/logger"
	"github.com/GotoRen/echoman/server/layers"
	golayers "github.com/google/gopacket/layers"
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
			dstIP := net.IP(buf[layers.IPv4offsetDst : layers.IPv4offsetDst+net.IPv4len]).To4()
			dstPort := golayers.UDPPort(binary.BigEndian.Uint16(buf[layers.DstUDPPortOffset : layers.DstUDPPortOffset+layers.DstUDPLength]))

			// fmt.Println("[DEBUG] Peer IPv4 Address", dstIP)
			// fmt.Println("[DEBUG] Peer IPv4 Port:", dstPort)
			// fmt.Println("[DEBUG] Chrous app IPv4:", net.ParseIP(device.Tun.VIP).To4())
			// fmt.Println("[DEBUG] Chrous app Port:", golayers.UDPPort(uint16(device.ChrousPort)))

			if _, err := device.Tun.Device.Tun.Write(buf); err != nil {
				logger.LogErr("Failed to write to tun/tap interface", "error", err)
			} else {
				/*************************************************************************************
				 * README: description for Chorus application *
				*************************************************************************************
				 * Checking the TUN -> Application packet flow using source code is complicated.
				 *   - For the time being, I will check with wireshark.app.
				 * Thefore, if the write to TUN succeeds, we generate and return a response message.
				 *   - Write the message generated at this time directly to the Real interface.
				 * ### Judgment method ###
				 *   - If the destination is "198.18.9.10:30910", judge it as chorus.app and return the message.
				 *   - And, return a response to the UDP packet received from the client.
				*************************************************************************************/
				if net.ParseIP(device.Tun.VIP).To4().Equal(dstIP) && golayers.UDPPort(uint16(device.ChorusPort)) == dstPort {
					logger.LogDebug("Receive chorus message", "chrous", "success")
					res := chorus.GenerateUDPResponsePacket(buf)
					if _, err := device.Peer.ConnUDP.WriteToUDP(res, &device.Peer.PeerEndPoint); err != nil {
						logger.LogErr("[Failed] Send chorus message", "error", err)
					} else {
						logger.LogDebug("Send chorus message", "chrous", "success")
					}
				}
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
