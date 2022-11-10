package internal

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/GotoRen/echoman/client/internal/logger"
)

type Device struct {
	EnvIPver int

	// Real Interface
	IfIndex   *net.Interface
	LocalIPv4 net.IP
	LocalIPv6 net.IP

	// TUN/TAP Interface
	Tun struct {
		Device *TunInterface
		mtu    int32
	}

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

	vipv4, vipv4prefix := generateVirtualIPv4()
	vipv6, vipv6prefix := generateVirtualIPv6()

	switch device.EnvIPver {
	case 4:
		device.Tun.Device, err = NewTunInterface(os.Getenv("TUN_INTERFACE_NAME"), vipv4.String(), vipv4prefix)
		if err != nil {
			logger.LogErr("Unable to create IPv4 TUN/TAP interface", "error", err)
		}
	case 6:
		device.Tun.Device, err = NewTunInterface(os.Getenv("TUN_INTERFACE_NAME"), vipv6.String(), vipv6prefix)
		if err != nil {
			logger.LogErr("Unable to create IPv6 TUN/TAP interface", "error", err)
		}
	}

	if err := device.Tun.Device.Up(); err != nil {
		logger.LogErr("Failed to Tunnel up", "error", err)
	}

	fmt.Println("[INFO] TUN IPv4:", device.Tun.Device.address)
}
