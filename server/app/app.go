package app

// func CreateApplicaton(device *internal.Device) {
// 	connUDP, err := CreateUDPConnection(30005)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	device.Mock.ConnUDP = connUDP
// }

// // makeUDPConnection makes UDP Connection.
// func makeUDPConnection(dstPort int) (connUDP *net.UDPConn, err error) {
// 	udpAddr, err := net.ResolveUDPAddr("udp", "198.18.9.10:"+strconv.Itoa(dstPort))
// 	if err != nil {
// 		return nil, err
// 	}

// 	connUDP, err = net.ListenUDP("udp", udpAddr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return connUDP, nil
// }

// // CreateUDPConnection creates a UDP connection with the peer node.
// func CreateUDPConnection(peerUDPport int) (*net.UDPConn, error) {
// 	connUDP, err := makeUDPConnection(peerUDPport)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return connUDP, nil
// }
