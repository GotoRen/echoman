package layers

import "net"

// IPv6 offset length.
const (
	IPv6offsetPayloadLength = 4                           // IPv6offsetPayloadLength is IPv6 offset payload length.
	IPv6offsetSrc           = 8                           // IPv6offsetSrc is IPv6 offset src length.
	IPv6offsetDst           = IPv6offsetSrc + net.IPv6len // IPv6offsetDst is IPv6 offset dst length.
)
