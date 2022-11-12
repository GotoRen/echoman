package app

// func NewICMPv4ReplayPacket(req []byte) []byte {
// 	// ここで送信元と宛先の情報を入れ替える
// 	dstMacAddr := net.HardwareAddr(req[layers.SrcMACAddrOffset : layers.SrcMACAddrOffset+layers.SrcMacLength])
// 	srcMacAddr := net.HardwareAddr(req[layers.DstMACAddrOffset : layers.DstMACAddrOffset+layers.DstMacLength])
// 	srcIPv4Addr := req[layers.DstIPv4AddrOffset : layers.DstIPv4AddrOffset+layers.DstIPv4Length]
// 	dstIPv4Addr := req[layers.SrcIPv4AddrOffset : layers.SrcIPv4AddrOffset+layers.SrcIPv4Length]

// 	res_data := req[layers.ICMPv4Dataoffset : layers.ICMPv4Dataoffset+layers.ICMPv4DataLength+layers.ICMPv4TimeStampLength]

// 	ether := golayers.Ethernet{
// 		DstMAC:       dstMacAddr,
// 		SrcMAC:       srcMacAddr,
// 		EthernetType: golayers.EthernetTypeIPv4,
// 	}

// 	ip := golayers.IPv4{
// 		Version:    4,
// 		TOS:        0,
// 		Length:     0,
// 		Id:         0,
// 		FragOffset: 0,
// 		TTL:        255,
// 		Protocol:   golayers.IPProtocolICMPv4,
// 		Checksum:   0,
// 		SrcIP:      srcIPv4Addr,
// 		DstIP:      dstIPv4Addr,
// 	}

// 	icmpv4 := golayers.ICMPv4{
// 		TypeCode: golayers.CreateICMPv4TypeCode(0, 0),
// 		Checksum: 0,
// 		Id:       10,
// 		Seq:      1,
// 	}

// 	options := gopacket.SerializeOptions{
// 		ComputeChecksums: true,
// 		FixLengths:       true,
// 	}

// 	buffer := gopacket.NewSerializeBuffer()

// 	if err := gopacket.SerializeLayers(buffer, options,
// 		&ether,
// 		&ip,
// 		&icmpv4,
// 		gopacket.Payload(res_data),
// 	); err != nil {
// 		logger.LogErr("Serialize error", "error", err)
// 		return nil
// 	}

// 	outgoingPacket := buffer.Bytes()

// 	return outgoingPacket
// }
