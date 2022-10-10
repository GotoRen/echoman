package internal

import (
	"fmt"
	"net"
	"strings"

	"github.com/GotoRen/echoman/client/internal/logger"
)

// getServerIPv4 returns the IPv4 address obtained from the interface.
func getServerIPv4(intfIndex string, interfaces []net.Interface) (ipv4 net.IP) {
	for _, intf := range interfaces {
		if addrs, err := intf.Addrs(); err == nil {
			for _, addr := range addrs {
				if intf.Name == intfIndex {
					strAddr := addr.String()
					ipv4 = net.ParseIP(strAddr[:strings.Index(strAddr, "/")])
					if ipv4.To4() != nil {
						return ipv4
					}
				}
			}
		}
	}
	return nil
}

// checkInterface returns the obtained interface information
func checkInterface(interf string) (netInterface *net.Interface, err error) {
	for {
		netInterface, err = net.InterfaceByName(interf)
		if err != nil {
			logger.LogErr("Can't find interface", "error", err)
		} else {
			break
		}
	}
	fmt.Println("Connected.")

	return netInterface, nil
}

// getHardwareAddr retruns extract the hardware information base on the interface name
func getHardwareAddr(netInterface *net.Interface) {
	name := netInterface.Name
	macAddress := netInterface.HardwareAddr

	fmt.Println("Hardware name :", name)
	fmt.Println("MAC address :", macAddress)

	// verify if the MAC address can be parsed properly
	hwAddr, err := net.ParseMAC(macAddress.String())
	if err != nil {
		logger.LogErr("No able to parse MAC address", "error", err)
	}
	fmt.Printf("Physical hardware address :%s \n", hwAddr.String())
}
