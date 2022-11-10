package internal

// func htons(host uint16) uint16 {
// 	return (host&0xff)<<8 | (host >> 8)
// }

// // etherSendSock creates a new send socket for IPv4 ethernet packet.
// func etherSendSock(intfIndex *net.Interface) (int, error) {
// 	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_IP)))
// 	if err != nil {
// 		return -1, err
// 	}

// 	addr := syscall.SockaddrLinklayer{
// 		Protocol: htons(syscall.ETH_P_ALL),
// 		Ifindex:  intfIndex.Index,
// 	}

// 	if err := syscall.Bind(fd, &addr); err != nil {
// 		return -1, err
// 	}

// 	return fd, nil
// }

// // etherRecvSock creates a new receive socket for IPv4 ethernet packet.
// func etherRecvSock(intfIndex *net.Interface) (int, error) {
// 	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_IP)))
// 	if err != nil {
// 		return -1, err
// 	}

// 	addr := syscall.SockaddrLinklayer{
// 		Protocol: htons(syscall.ETH_P_ALL),
// 		Ifindex:  intfIndex.Index,
// 	}
// 	if err := syscall.Bind(fd, &addr); err != nil {
// 		return -1, err
// 	}

// 	// Received in promiscuous mode
// 	if err := syscall.SetLsfPromisc(intfIndex.Name, true); err != nil {
// 		return -1, err
// 	}

// 	return fd, nil
// }

// // SendEtherPacket uses a send socket to send an Ethernete packet.
// func SendEtherPacket(fd int, b []byte) error {
// 	if _, err := syscall.Write(fd, b); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // CreateDescriptor creates socket descriptor.
// func (device *Device) CreateDescriptor() {
// 	var err error

// 	// send socket
// 	device.Sd4soc, err = etherSendSock(device.IfIndex)
// 	if err != nil {
// 		logger.LogErr("Failed to open send IPv4 raw socket", "error", err)
// 	}

// 	// receive socket
// 	device.Rv4soc, err = etherRecvSock(device.IfIndex)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
