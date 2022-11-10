package internal

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"net"
	"strings"

	"github.com/GotoRen/echoman/server/internal/logger"
	"github.com/google/uuid"
)

// rfc3330, rfc2544.
// Using the Documentation addresses and prefixes.
var VIPv4NetworkAddress string = "198.18.0.0/16"
var VIPv6NetworkAddress string = "2001:0db8:c0ff:ee00::/64"

func getVIPv4NetworkAddressInfo() (vipv4nwaddress net.IP, vipv4netmask net.IPMask, vipv4prefix string) {
	_, ipnet, err := net.ParseCIDR(VIPv4NetworkAddress)
	if err != nil {
		logger.LogErr("Unable to get Virtual IPv4 network address", "error", err)
	}

	vipv4nwaddress = ipnet.IP
	vipv4netmask = ipnet.Mask
	vipv4prefix = VIPv4NetworkAddress[strings.Index(VIPv4NetworkAddress, "/"):]

	return
}

func getVIPv6NetworkAddressInfo() (vipv6nwaddress net.IP, vipv6netmask net.IPMask, vipv6prefix string) {
	_, ipnet, err := net.ParseCIDR(VIPv6NetworkAddress)
	if err != nil {
		logger.LogErr("Unable to get Virtual IPv4 network address", "error", err)
	}

	vipv6nwaddress = ipnet.IP
	vipv6netmask = ipnet.Mask
	vipv6prefix = VIPv6NetworkAddress[strings.Index(VIPv6NetworkAddress, "/"):]

	return
}

// generateVirtualIPv4 is to generate virtual IPv4 address.
func generateVirtualIPv4() (net.IP, string) {
	vipv4nwaddress, _, vipv4prefix := getVIPv4NetworkAddressInfo()
	b := make([]byte, 2)

createVIPv4:
	err := binary.Read(rand.Reader, binary.BigEndian, &b)
	if err != nil {
		logger.LogErr("Failed to generate random 2 bytes for VirtualIPv4", "error", err)
	}

	vipv4 := net.IPv4(vipv4nwaddress[0], vipv4nwaddress[1], b[0], b[1]).To4()

	// TODO: If you use netip.Addr type you can use switch-case.
	// 198.18.0.0/16 => network address or 198.18.255.255/16 => broadcast address
	if vipv4.Equal(net.IPv4(vipv4nwaddress[0], vipv4nwaddress[1], 0x00, 0x00).To4()) || vipv4.Equal(net.IPv4(vipv4nwaddress[0], vipv4nwaddress[1], 0xff, 0xff).To4()) {
		goto createVIPv4
	}

	return vipv4, vipv4prefix
}

// generateVirtualIPv6 is to generate virtual IPv6 address.
func generateVirtualIPv6() (net.IP, string) {
	vipv6nwaddress, _, vipv6prefix := getVIPv6NetworkAddressInfo()

	u, err := uuid.NewRandom()
	if err != nil {
		logger.LogErr("Failed to generate UUID for VirtualIPv6", "error", err)
	}

	ub, err := u.MarshalBinary()
	if err != nil {
		logger.LogErr("For VirtualIPv6: Failed encode to binary format", "error", err)
	}

	vipv6 := net.ParseIP(strings.Replace(vipv6nwaddress.String(), "::", fmt.Sprintf(":%x:%x:%x:%x", ub[8:10], ub[10:12], ub[12:14], ub[14:16]), 1))

	return vipv6, vipv6prefix
}
