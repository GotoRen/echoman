package main

import (
	"fmt"
	"net"
)

func main() {
	var interf string
	var netInterface *net.Interface
	var err error

	// interf = "en0"
	interf = "eth0"
	netInterface, err = checkInterface(interf)
	if err != nil {
		fmt.Println(err)
	}
	getHardwareAddr(netInterface)

	// interf = "en10"
	interf = "eth1"
	netInterface, err = checkInterface(interf)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(netInterface.Name)
	// fmt.Println(netInterface.HardwareAddr)

	getHardwareAddr(netInterface)
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
