package internal

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type InterfaceInfo struct {
	IfIndex    *net.Interface
	ServerIPv4 net.IP
	ServerMAC  net.HardwareAddr
}

// GetServerInfo returns the server's L3 and L2 address information.
func GetServerInfo(intf string) (i *InterfaceInfo) {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	netInterface, err := net.InterfaceByName(intf)
	if err != nil {
		log.Fatal(err)
	}

	ipv4 := getServerIPv4(intf, interfaces)
	if ipv4 == nil {
		log.Fatal("[ERROR]: ipv4 address is empty.")
	}

	i = &InterfaceInfo{
		IfIndex:    netInterface,
		ServerIPv4: ipv4,
		ServerMAC:  netInterface.HardwareAddr,
	}

	return i
}

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
			fmt.Println("Can't find interface")
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
		fmt.Println("No able to parse MAC address :", err)
	}
	fmt.Printf("Physical hardware address :%s \n", hwAddr.String())
}

func SetAdapterInterface() (*net.Interface, error) {
	netInterface, err := net.InterfaceByName("eth0")
	if err != nil {
		return nil, err
	}

	return netInterface, err
}
