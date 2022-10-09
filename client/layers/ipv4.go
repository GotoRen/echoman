package layers

/***********************************
[*] L3(IPv4) : 20 byte
ex:)
Version: 69
Diffrentiated Services: 0
Total Length: 72
Identification: 35 73
Offset: 0 0
TTL: 255
protocol: 17
checksum: 151 92
srcIP: 0 0 0 0
dstIP: 255 255 255 255
***********************************/

type IPv4Header struct {
	VersionAndHeaderLenght []byte // 1 byte
	ServiceType            []byte // 1 byte
	TotalPacketLength      []byte // 2 byte
	PacketIdentification   []byte // 2 byte
	FlagOffset             []byte //  2 byte
	TTL                    []byte // 1 byte
	Protocol               []byte // 1 byte
	HeaderCheckSum         []byte // 2 byte
	SourceIPAddr           []byte // 4 byte
	DstIPAddr              []byte // 4 byte
}
