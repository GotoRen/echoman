package internal

import (
	"fmt"
	"net"
	"syscall"
)

func htons(host uint16) uint16 {
	return (host&0xff)<<8 | (host >> 8)
}

// SendIPv4RawSocket creates a raw socket for sending IPv4 packet.
func SendIPv4RawSocket() (int, error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	if err != nil {
		fmt.Println("[DEBUG] error 1")
		return -1, err
	}

	ip := net.ParseIP("0.0.0.0").To4()
	addr := syscall.SockaddrInet4{
		Addr: [4]byte{ip[0], ip[1], ip[2], ip[3]},
	}

	if err = syscall.Bind(fd, &addr); err != nil {
		fmt.Println("[DEBUG] error 2")
		return -1, err
	}

	return fd, nil
}

// SendPacket4 sends IPv4 packet.
func SendPacket4(s int, b []byte, ip []byte) error {
	addr := syscall.SockaddrInet4{
		Addr: [4]byte{ip[0], ip[1], ip[2], ip[3]},
	}

	if err := syscall.Sendto(s, b, 0, &addr); err != nil {
		return err
	}

	return nil
}

func RecvIPv4RawSocket(intfIndex *net.Interface) (int, error) {
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

func SendEtherPacket(fd int, b []byte) error {
	if _, err := syscall.Write(fd, b); err != nil {
		return err
	}

	return nil
}
