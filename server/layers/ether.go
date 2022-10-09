package layers

import "net"

type EtherHeader struct {
	DstMacAddr net.HardwareAddr
	SrcMacAddr net.HardwareAddr
	ProtoType  uint16
}

const (
	EthTypeArp  uint16 = 0x0806
	EthTypeIpv4 uint16 = 0x0800
	EthTypeIpv6 uint16 = 0x86dd
)
