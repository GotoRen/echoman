package internal

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/GotoRen/echoman/server/internal/logger"
)

type Device struct {
	EnvIPver int

	// Real Interface
	IfIndex      *net.Interface
	LocalIPv4    net.IP
	LocalIPv6    net.IP
	LocalUDPPort uint16

	socket struct {
		sd4soc int
		rv4soc int
	}

	// TUN/TAP Interface
	Tun struct {
		Device *TunInterface
		VIP    string
		mtu    int32
	}

	ChorusPort int

	Peer *Peer
}

// NewDevice defines device information.
func NewDevice(intf string) (device *Device) {
	envIPVer, err := strconv.Atoi(os.Getenv("VIRTUAL_IP_TYPE"))
	if err != nil {
		logger.LogErr("Unable to get VIRTUAL_IP_TYPE", "error", err)
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	localIPv4addr := getServerIPv4(intf, interfaces)
	if localIPv4addr == nil {
		logger.LogErr("IPv4 address is empty", "error", err)
	}

	device = &Device{
		EnvIPver:  envIPVer,
		LocalIPv4: localIPv4addr,
	}

	return device
}

// CreateTunInterface creates a TUN/TAP interface.
func (device *Device) CreateTunInterface() {
	var err error

	switch device.EnvIPver {
	case 4:
		device.Tun.Device, err = NewTunInterface(os.Getenv("TUN_INTERFACE_NAME"), os.Getenv("ECHOMAN_SERVER_IPV4_TUN"), "/16")
		if err != nil {
			logger.LogErr("Failed to create Tunnel Interface Virtual IPv4", "error", err)
		}
	case 6:
		device.Tun.Device, err = NewTunInterface(os.Getenv("TUN_INTERFACE_NAME"), os.Getenv("ECHOMAN_SERVER_IPV6_TUN"), "/64")
		if err != nil {
			logger.LogErr("Failed to create Tunnel Interface Virtual IPv6", "error", err)
		}
	}

	if err := device.Tun.Device.Up(); err != nil {
		logger.LogErr("Failed to Tunnel up", "error", err)
	}

	device.Tun.VIP = device.Tun.Device.address[:strings.Index(device.Tun.Device.address, "/")]
}
