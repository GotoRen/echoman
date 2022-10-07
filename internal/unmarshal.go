package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

func RewriteEtherAddr(s string) []byte {
	replaced := strings.Replace(s, ":", "", -1)
	p, err := hex.DecodeString(replaced)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func main() {
	buf := make([]byte, 100) // length

	// // Ether: 6byte
	// // d0:37:45:50:0b:06 -> 208 55 69 80 11 6
	// dstMac := [8]byte{
	// 	0xd0, 0x37, 0x45, 0x50,
	// 	0x0b, 0x06, 0x00, 0x00,
	// }
	dstMac := RewriteEtherAddr("d0:37:45:50:0b:06")

	// // Ether: 6byte
	// // d0:37:45:2c:e2:35 -> 208 55 69 44 226 53
	// srcMac := [8]byte{
	// 	0xd0, 0x37, 0x45, 0x2c,
	// 	0xe2, 0x35, 0x00, 0x00,
	// }
	srcMac := RewriteEtherAddr("d0:37:45:2c:e2:35")

	binary.BigEndian.PutUint64(buf[0:8], binary.BigEndian.Uint64(dstMac[0:8]))  // 宛先Mac
	binary.BigEndian.PutUint64(buf[6:14], binary.BigEndian.Uint64(srcMac[0:8])) // 送信元Mac

	fmt.Println("buf:", buf[:12])
}
