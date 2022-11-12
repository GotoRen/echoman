package internal

import (
	"fmt"
	"net"
	"syscall"

	"github.com/GotoRen/echoman/client/internal/logger"
)

func htons(host uint16) uint16 {
	return (host&0xff)<<8 | (host >> 8)
}

// etherSendSock creates a new send socket for IPv4 ethernet packet.
func etherSendSock(intfIndex *net.Interface) (int, error) {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_IP)))
	if err != nil {
		return -1, err
	}

	addr := syscall.SockaddrLinklayer{
		Protocol: htons(syscall.ETH_P_ALL),
		Ifindex:  intfIndex.Index,
	}

	if err := syscall.Bind(fd, &addr); err != nil {
		return -1, err
	}

	return fd, nil
}

// vipSendSock creates a new send socket for Virtual IPv4 packet.
func vipSendSock(sip string) (int, error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	if err != nil {
		return -1, err
	}

	fmt.Println("sip:", sip)

	ip := net.ParseIP(sip)
	addr := syscall.SockaddrInet4{
		Addr: [4]byte{ip[0], ip[1], ip[2], ip[3]},
	}

	if err = syscall.Bind(fd, &addr); err != nil {
		return -1, err
	}

	return fd, nil
}

// etherRecvSock creates a new receive socket for IPv4 ethernet packet.
func etherRecvSock(intfIndex *net.Interface) (int, error) {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_IP)))
	if err != nil {
		return -1, err
	}

	addr := syscall.SockaddrLinklayer{
		Protocol: htons(syscall.ETH_P_ALL),
		Ifindex:  intfIndex.Index,
	}
	if err := syscall.Bind(fd, &addr); err != nil {
		return -1, err
	}

	// Received in promiscuous mode
	if err := syscall.SetLsfPromisc(intfIndex.Name, true); err != nil {
		return -1, err
	}

	return fd, nil
}

// SendEtherPacket uses a send socket to send an ether packet.
func SendEtherPacket(fd int, b []byte) error {
	if _, err := syscall.Write(fd, b); err != nil {
		return err
	}

	return nil
}

// SendPacket4 sends IPv4 packet.
func SendPacket4(fd int, b []byte, dip []byte) error {
	addr := syscall.SockaddrInet4{
		Addr: [4]byte{dip[0], dip[1], dip[2], dip[3]},
	}

	if err := syscall.Sendto(fd, b, 0, &addr); err != nil {
		return err
	}

	return nil
}

// CreateDescriptor creates socket descriptor.
func (device *Device) CreateDescriptor() {
	var err error

	// send socket
	device.Sd4soc, err = vipSendSock(device.Tun.VIP)
	if err != nil {
		logger.LogErr("Failed to open send IPv4 raw socket", "error", err)
	}
}

// CloseRawSocket closes opening file descriptors.
func (device *Device) CloseRawSocket() {
	closeRawSocket(device.Sd4soc, "send IPv4")
}

// closeRawSocket closes opening file descriptor.
func closeRawSocket(fd int, fdType string) {
	if fd == -1 {
		return
	}

	if err := syscall.Close(fd); err != nil {
		message := "Failed to close the " + fdType + " Raw socket"
		logger.LogErr(message, "error", err)
	}
}
