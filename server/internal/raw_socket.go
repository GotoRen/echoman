package internal

import (
	"net"
	"os"
	"syscall"
)

func htons(host uint16) uint16 {
	return (host&0xff)<<8 | (host >> 8)
}

// -------------------------------------------------------------------- //
// Send sockets
// -------------------------------------------------------------------- //
func EtherSendSock(intfIndex *net.Interface) (int, error) {
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

// -------------------------------------------------------------------- //
// Receive sockets
// -------------------------------------------------------------------- //
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

	// Received in promiscuous mode
	if err := syscall.SetLsfPromisc(os.Getenv("SERVER_INTERFACE"), true); err != nil {
		return -1, err
	}

	return fd, nil
}
