package internal

import (
	"log"
	"net"
	"os"
	"strconv"

	"github.com/GotoRen/echoman/client/internal/logger"
)

type Device struct {
	IfIndex      *net.Interface
	LocalIPv4    net.IP
	LocalMAC     net.HardwareAddr
	LocalUDPPort uint16
	Sd4soc       int
	Rv4soc       int

	Peer *Peer
}

func GetDeviceInfo(intf string) (device *Device) {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	netInterface, err := net.InterfaceByName(intf)
	if err != nil {
		log.Fatal(err)
	}

	localIPv4addr := getServerIPv4(intf, interfaces)
	if localIPv4addr == nil {
		log.Fatal("[ERROR]: ipv4 address is empty.")
	}

	localMACaddr := netInterface.HardwareAddr

	localUDPport, err := strconv.Atoi(os.Getenv("LOCAL_UDP_PORT"))
	if err != nil {
		logger.LogErr("Type conversion failed", "error", err)
	}

	device = &Device{
		IfIndex:      netInterface,
		LocalIPv4:    localIPv4addr,
		LocalMAC:     localMACaddr,
		LocalUDPPort: uint16(localUDPport),
		Peer:         GerPeerInfo(),
	}

	return device
}
