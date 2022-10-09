package internal

import "net"

func SetAdapterInterface() (*net.Interface, error) {
	netInterface, err := net.InterfaceByName("eth0")
	if err != nil {
		return nil, err
	}

	return netInterface, err
}
