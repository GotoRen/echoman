package layers

/****************
[*] ICMPv4 : 64 byte
ex.)
Type : 8
Code : 0
Checksum : 136 236
Identifier : 33 186
SequenceNumber : 0 0
TimeStamp : 97 174 80 213 0 12 175 198
Data : 48
****************/

type ICMPv4 struct {
	Type           []byte // 1 byte
	Code           []byte // 1 byte
	CheckSum       []byte // 2 byte
	Identification []byte // 2 byte
	SequenceNumber []byte // 2 byte
	TimeStamp      []byte // 8 byte
	Data           []byte // 48 byte
}
