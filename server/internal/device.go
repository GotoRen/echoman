package internal

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/GotoRen/echoman/server/internal/logger"
)

type Device struct {
	EnvIPver int

	IfIndex   *net.Interface
	LocalIPv4 net.IP
	LocalIPv6 net.IP
	LocalMAC  net.HardwareAddr

	LocalUDPPort uint16

	Sd4soc int
	Rv4soc int
	Sd6soc int
	Rv6soc int

	ConnUDP *net.UDPConn

	Tun struct {
		Device *TunInterface
		mtu    int32
	}

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

	envIPVer, err := strconv.Atoi(os.Getenv("VIRTUAL_IP_TYPE"))
	if err != nil {
		logger.LogErr("Unable to get VIRTUAL_IP_TYPE", "error", err)
	}

	device = &Device{
		EnvIPver:     envIPVer,
		IfIndex:      netInterface,
		LocalIPv4:    localIPv4addr,
		LocalMAC:     localMACaddr,
		LocalUDPPort: uint16(localUDPport),
		Peer:         GerPeerInfo(),
	}

	return device
}

func (device *Device) CreateTunInterface() {
	var err error

	switch device.EnvIPver {
	case 4:
		fmt.Println(os.Getenv("ECHOMAN_SERVER_IPV4_TUN"))
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
}
