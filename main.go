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
