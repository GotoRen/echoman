package mock

import (
	"net"

	"github.com/GotoRen/echoman/client/internal"
	"github.com/GotoRen/echoman/client/internal/logger"
)

func NewUDPPacket(device *internal.Device) {
	udpPacket, err := GenerateUDPPacket(device)
	if err != nil {
		logger.LogErr("Failed to generate udp packet", "error", err)
	}

	if err := internal.SendPacket4(device.Sd4soc, udpPacket, net.ParseIP(ECHOMAN_SERVER_IPV4_TUN).To4()); err != nil {
		logger.LogErr("Failed to write to tun/tap interface", "error", err)
	}
}
