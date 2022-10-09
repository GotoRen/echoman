package logger

// import (
// 	"fmt"

// 	"github.com/google/gopacket"
// 	"github.com/google/gopacket/layers"
// 	golayers "github.com/google/gopacket/layers"
// 	"github.com/krolaw/dhcp4"
// )

// // This function is important!!
// /****************
// [*] L2 : 14 byte
// dstMac 255 255 255 255 255 255
// srcMac 208 55 69 44 226 53
// protcolType 8 0

// [*] L3 : 20 byte
// Version 69
// Diffrentiated Services 0
// Total Length 72
// Identification 35 73
// Offset 0 0
// TTL 255
// protocol 17
// checksum 151 92
// srcIP 0 0 0 0
// dstIP 255 255 255 255

// [*] L4(UDP) : 8 byte
// srcPort 0 68
// dstPort 0 67
// Length 1 52
// checksum 22 139
// ****************/
// func DebugUDPTypePacketLayer(b []byte) {
// 	fmt.Println("---------------------")
// 	fmt.Println("[*] DstMAC:", b[:6])
// 	fmt.Println("[*] SrcMAC:", b[6:12])
// 	fmt.Println("[*] ProtcolType:", b[12:14])
// 	fmt.Println("---------------------")
// 	fmt.Println("[*] Version:", b[14])
// 	fmt.Println("[*] Differentiated Services:", b[15])
// 	fmt.Println("[*] Total Length:", b[16])
// 	fmt.Println("[*] Identification:", b[17:19])
// 	fmt.Println("[*] Offset:", b[20:21])
// 	fmt.Println("[*] TTL:", b[22])
// 	fmt.Println("[*] Protocol:", b[23])
// 	fmt.Println("[*] Checksum:", b[24:26])
// 	fmt.Println("[*] SrcIP:", b[26:30])
// 	fmt.Println("[*] DstIP:", b[30:34])
// 	fmt.Println("---------------------")
// 	fmt.Println("[*] SrcPort:", b[34:36])
// 	fmt.Println("[*] DstPort:", b[36:38])
// 	fmt.Println("[*] Length:", b[38:40])
// 	fmt.Println("[*] Checksum:", b[40:42])
// 	fmt.Println("---------------------")
// 	fmt.Println("[*] Payload:", b[42:])
// 	fmt.Println("---------------------")
// }

// func DebugDHCPResponse(dhcpresp *dhcp4.Packet) {
// 	fmt.Println("-------------------------------------------------")
// 	fmt.Println("[*] DHCP Response - Type:", dhcpresp.Type)
// 	fmt.Println("[*] DHCP Response - TransactionID:", dhcpresp.TransactionID)
// 	fmt.Println("[*] DHCP Response - Broadcast:", dhcpresp.Broadcast)
// 	fmt.Println("[*] DHCP Response - HardwareAddr:", dhcpresp.HardwareAddr)
// 	fmt.Println("[*] DHCP Response - ClientAddr:", dhcpresp.ClientAddr)
// 	fmt.Println("[*] DHCP Response - YourAddr:", dhcpresp.YourAddr)
// 	fmt.Println("[*] DHCP Response - ServerAddr:", dhcpresp.ServerAddr)
// 	fmt.Println("[*] DHCP Response - RelayAddr:", dhcpresp.RelayAddr)
// 	fmt.Println("[*] DHCP Response - BootServerName:", dhcpresp.BootServerName)
// 	fmt.Println("[*] DHCP Response - BootFilename:", dhcpresp.BootFilename)
// 	fmt.Println("[*] DHCP Response - Options:", dhcpresp.Options)
// 	fmt.Println("-------------------------------------------------")
// }

// func DebugARPPacket(arp *golayers.ARP) {
// 	fmt.Println("[DEBUG] ARP HardwareType:", arp.AddrType)
// 	fmt.Println("[DEBUG] ARP ProtocolType:", arp.Protocol)
// 	fmt.Println("[DEBUG] ARP OpCode:", arp.Operation) // 1: req / 2: res
// 	fmt.Println("[DEBUG] ARP SHA:", arp.SourceHwAddress)
// 	fmt.Println("[DEBUG] ARP SPA:", arp.SourceProtAddress)
// 	fmt.Println("[DEBUG] ARP DHA:", arp.DstHwAddress)
// 	fmt.Println("[DEBUG] ARP DPA:", arp.DstProtAddress)
// }

// /****************
// [*] ICMP : 64 byte
// Type : 8
// Code : 0
// Checksum : 136 236
// Identifier : 33 186
// SequenceNumber : 0 0
// TimeStamp : 97 174 80 213 0 12 175 198
// Data : 48
// ****************/

// func DebugIPv4Packet(b []byte) {
// 	fmt.Println()
// 	fmt.Println("---------------------------------------------")
// 	fmt.Println("IPv4 Layer")
// 	fmt.Println("---------------------------------------------")
// 	fmt.Println("[*] Version:", b[0])
// 	fmt.Println("[*] Differentiated Services:", b[1])
// 	fmt.Println("[*] Total Length:", b[2])
// 	fmt.Println("[*] Identification:", b[3:5])
// 	fmt.Println("[*] Offset:", b[6:8])
// 	fmt.Println("[*] TTL:", b[9])
// 	fmt.Println("[*] Protocol:", b[10])
// 	fmt.Println("[*] Checksum:", b[11:12])
// 	fmt.Println("[*] SrcIP:", b[12:16])
// 	fmt.Println("[*] DstIP:", b[16:20])
// 	fmt.Println()
// }

// func DebugICMPv4Packet(b []byte) {
// 	fmt.Println("---------------------------------------------")
// 	fmt.Println("ICMP Layer")
// 	fmt.Println("---------------------------------------------")
// 	fmt.Println("[*] Type:", b[20])
// 	fmt.Println("[*] Code:", b[21])
// 	fmt.Println("[*] Checksum:", b[22:24])
// 	fmt.Println("[*] Identifier:", b[24:26])
// 	fmt.Println("[*] SequenceNumber:", b[26:28])
// 	fmt.Println("[*] TimeStamp:", b[28:36])
// 	fmt.Println("[*] Data:", b[36:84])
// }

// func ICMPv4Info(buf []byte) {
// 	packet := gopacket.NewPacket(buf, golayers.LayerTypeICMPv4, gopacket.Default)
// 	icmpLayer := packet.Layer(golayers.LayerTypeICMPv4)
// 	if icmpLayer != nil {
// 		icmp, _ := icmpLayer.(*golayers.ICMPv4)
// 		fmt.Println("[*] ICMP - TypeCode:", icmp.TypeCode)
// 		fmt.Println("[*] ICMP - Chechsum:", icmp.Checksum)
// 		fmt.Println("[*] ICMP - Identifier:", icmp.Id)
// 		fmt.Println("[*] ICMP - SequenceNumber:", icmp.Seq)
// 	}
// }

// func DebugTCPPacket(buf []byte) {
// 	packet := gopacket.NewPacket(buf, golayers.LayerTypeTCP, gopacket.Default)
// 	tcpLayer := packet.Layer(golayers.LayerTypeTCP)
// 	if tcpLayer != nil {
// 		tcp, _ := tcpLayer.(*golayers.TCP)
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("TCP Layer")
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("[*] TCP - SrcPort:", tcp.SrcPort.String())
// 		fmt.Println("[*] TCP - DstPort:", tcp.DstPort.String())
// 		fmt.Println("[*] TCP - Sequence:", tcp.Seq)
// 		fmt.Println("[*] TCP - Ack number:", tcp.Ack)
// 		fmt.Println("[*] TCP - DataOffset:", tcp.DataOffset)
// 		fmt.Println("[*] TCP - FIN:", tcp.FIN)
// 		fmt.Println("[*] TCP - SYN:", tcp.SYN)
// 		fmt.Println("[*] TCP - RST:", tcp.RST)
// 		fmt.Println("[*] TCP - PSH:", tcp.PSH)
// 		fmt.Println("[*] TCP - ACK:", tcp.ACK)
// 		fmt.Println("[*] TCP - URG:", tcp.URG)
// 		fmt.Println("[*] TCP - ECE:", tcp.ECE)
// 		fmt.Println("[*] TCP - CWR:", tcp.CWR)
// 		fmt.Println("[*] TCP - NS:", tcp.NS)
// 		fmt.Println("[*] TCP - Window Size:", tcp.Window)
// 		fmt.Println("[*] TCP - Checksum:", tcp.Checksum)
// 		fmt.Println("[*] TCP - Urgent:", tcp.Urgent)
// 		fmt.Println("[*] TCP - Options:", tcp.Options)
// 	}
// }

// func DebugDNSPacket(buf []byte) {
// 	packet := gopacket.NewPacket(buf, golayers.LayerTypeUDP, gopacket.Default)
// 	udpLayer := packet.Layer(golayers.LayerTypeDNS)
// 	if udpLayer != nil {
// 		udp, _ := udpLayer.(*golayers.DNS)
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("UDP Layer")
// 		fmt.Println("---------------------------------------------")
// 		fmt.Print("[*] UDP - Answers:", udp.Answers)
// 	}
// }

// //===================================================================================//
// // Version: [0]
// // Differentiated Services: [1]
// // Total Length: [2]
// // Identification: [3:5]
// // Offset: [6:8]
// // TTL: [9]
// // Protocol: [10]
// // Checksum: [11:12]
// // SrcIP: [12:16]
// // DstIP: [16:20]

// // SrcPort: [20:22]
// // DstPort: [22:24]
// // Sequence: [24:28]
// // Ack number: [28:32]
// // Flags: [32:34]
// // Window Size: [34:36]
// // Checksum: [36:38]
// // Urgent: [38:40]
// // Data: [41:]

// func TCPPacektDebug(b []byte) {
// 	fmt.Println("---------------------------------------------")
// 	// fmt.Printf("[*] IP - Version: %x\n", b[0])
// 	// fmt.Printf("[*] IP - Differentiated Services: %x\n", b[1])
// 	// fmt.Printf("[*] IP - Total Length: %x\n", b[2:3])
// 	// fmt.Printf("[*] IP - Identification: %x\n", b[4:5])
// 	// fmt.Printf("[*] IP - Offset: %x\n", b[6:8])
// 	// fmt.Printf("[*] IP - Time To Live: %x\n", b[9])
// 	// fmt.Printf("[*] IP - Protocol: %x\n", b[10])
// 	// fmt.Printf("[*] IP - Checksum: %x\n", b[11:12])
// 	// fmt.Printf("[*] IP - SrcIP: %x\n", b[12:16])
// 	// fmt.Printf("[*] IP - DstIP: %x\n", b[16:20])

// 	// fmt.Printf("[*] TCP - SrcPort: %x\n", b[20:22])
// 	// fmt.Printf("[*] TCP - DstPort: %x\n", b[22:24])
// 	// fmt.Printf("[*] TCP - Sequence: %x\n", b[24:28])
// 	// fmt.Printf("[*] TCP - Ack number: %x\n", b[28:32])
// 	// fmt.Printf("[*] TCP - Flags: %x\n", b[32:34])
// 	// fmt.Printf("[*] TCP - Window Size: %x\n", b[34:36])
// 	// fmt.Printf("[*] TCP - Checksum: %x\n", b[36:38])
// 	// fmt.Printf("[*] TCP - Urgent: %x\n", b[38:40])
// 	fmt.Println("[*] IP - Version:", b[0])
// 	fmt.Println("[*] IP - Differentiated Services:", b[1])
// 	fmt.Println("[*] IP - Total Length:", b[2:4])
// 	fmt.Println("[*] IP - Identification:", b[4:6])
// 	fmt.Println("[*] IP - Offset:", b[6:8])
// 	fmt.Println("[*] IP - Time To Live:", b[8])
// 	fmt.Println("[*] IP - Protocol:", b[9])
// 	fmt.Println("[*] IP - Checksum:", b[10:12])
// 	fmt.Println("[*] IP - SrcIP:", b[12:16])
// 	fmt.Println("[*] IP - DstIP:", b[16:20])

// 	fmt.Println("[*] TCP - SrcPort:", b[20:22])
// 	fmt.Println("[*] TCP - DstPort:", b[22:24])
// 	fmt.Println("[*] TCP - Sequence:", b[24:28])
// 	fmt.Println("[*] TCP - Ack number:", b[28:32])
// 	fmt.Println("[*] TCP - Flags:", b[32:34])
// 	fmt.Println("[*] TCP - Window Size:", b[34:36])
// 	fmt.Println("[*] TCP - Checksum:", b[36:38])
// 	fmt.Println("[*] TCP - Urgent:", b[38:40])
// 	// Options
// 	// Data
// 	fmt.Println("---------------------------------------------")
// }

// // ===================================================================================//
// func PrintIPPacketInfo(data []byte) {
// 	packet := gopacket.NewPacket(data, golayers.LayerTypeIPv4, gopacket.Default)

// 	// IP Layer
// 	ipLayer := packet.Layer(layers.LayerTypeIPv4)
// 	if ipLayer != nil {
// 		ip, _ := ipLayer.(*layers.IPv4)
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("IP Layer")
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("[*] IP - Version:", ip.Version)
// 		fmt.Println("[*] IP - IHL:", ip.IHL)
// 		fmt.Println("[*] IP - Type of Service:", ip.TOS)
// 		fmt.Println("[*] IP - Total Length:", ip.Length)
// 		fmt.Println("[*] IP - Identification:", ip.Id)
// 		fmt.Println("[*] IP - Flags:", ip.Flags)
// 		fmt.Println("[*] IP - Fragment Offset:", ip.FragOffset)
// 		fmt.Println("[*] IP - Time To Live:", ip.TTL)
// 		fmt.Println("[*] IP - Protocol:", ip.Protocol)
// 		fmt.Println("[*] IP - Header Checksum:", ip.Checksum)
// 		fmt.Println("[*] IP - SrcIP:", ip.SrcIP)
// 		fmt.Println("[*] IP - DstIP:", ip.DstIP)
// 		fmt.Println("[*] IP - Options:", ip.Options)
// 	}

// 	// UDP Layer
// 	udpLayer := packet.Layer(golayers.LayerTypeUDP)
// 	if udpLayer != nil {
// 		udp, _ := udpLayer.(*golayers.UDP)
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("UDP Layer")
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("[*] UDP - SrcPort:", udp.SrcPort.String())
// 		fmt.Println("[*] UDP - DstPort:", udp.DstPort.String())
// 		fmt.Println("[*] UDP - Length:", udp.Length)
// 		fmt.Println("[*] UDP - Checksum:", udp.Checksum)
// 	}

// 	// TCP Layer
// 	tcpLayer := packet.Layer(golayers.LayerTypeTCP)
// 	if tcpLayer != nil {
// 		tcp, _ := tcpLayer.(*golayers.TCP)
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("TCP Layer")
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("[*] TCP - SrcPort:", tcp.SrcPort.String())
// 		fmt.Println("[*] TCP - DstPort:", tcp.DstPort.String())
// 		fmt.Println("[*] TCP - Sequence:", tcp.Seq)
// 		fmt.Println("[*] TCP - Ack number:", tcp.Ack)
// 		fmt.Println("[*] TCP - DataOffset:", tcp.DataOffset)
// 		fmt.Println("[*] TCP - FIN:", tcp.FIN)
// 		fmt.Println("[*] TCP - SYN:", tcp.SYN)
// 		fmt.Println("[*] TCP - RST:", tcp.RST)
// 		fmt.Println("[*] TCP - PSH:", tcp.PSH)
// 		fmt.Println("[*] TCP - ACK:", tcp.ACK)
// 		fmt.Println("[*] TCP - URG:", tcp.URG)
// 		fmt.Println("[*] TCP - ECE:", tcp.ECE)
// 		fmt.Println("[*] TCP - CWR:", tcp.CWR)
// 		fmt.Println("[*] TCP - NS:", tcp.NS)
// 		fmt.Println("[*] TCP - Window Size:", tcp.Window)
// 		fmt.Println("[*] TCP - Checksum:", tcp.Checksum)
// 		fmt.Println("[*] TCP - Urgent:", tcp.Urgent)
// 		fmt.Println("[*] TCP - Options:", tcp.Options)
// 	}
// }

// func PrintEtherFrame(data []byte) {
// 	packet := gopacket.NewPacket(data, golayers.LayerTypeEthernet, gopacket.Default)

// 	// Ether Layer
// 	etherLayer := packet.Layer(layers.LayerTypeEthernet)
// 	if etherLayer != nil {
// 		ether, _ := etherLayer.(*layers.Ethernet)
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("Ether Layer")
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("[*] Ether - DstMAC:", ether.DstMAC)
// 		fmt.Println("[*] Ether - SrcMAC:", ether.SrcMAC)
// 		fmt.Println("[*] Ether - Protocol Type:", ether.EthernetType)
// 	}

// 	// IP Layer
// 	ipLayer := packet.Layer(layers.LayerTypeIPv4)
// 	if ipLayer != nil {
// 		ip, _ := ipLayer.(*layers.IPv4)
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("IP Layer")
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("[*] IP - Version:", ip.Version)
// 		fmt.Println("[*] IP - IHL:", ip.IHL)
// 		fmt.Println("[*] IP - Type of Service:", ip.TOS)
// 		fmt.Println("[*] IP - Total Length:", ip.Length)
// 		fmt.Println("[*] IP - Identification:", ip.Id)
// 		fmt.Println("[*] IP - Flags:", ip.Flags)
// 		fmt.Println("[*] IP - Fragment Offset:", ip.FragOffset)
// 		fmt.Println("[*] IP - Time To Live:", ip.TTL)
// 		fmt.Println("[*] IP - Protocol:", ip.Protocol)
// 		fmt.Println("[*] IP - Header Checksum:", ip.Checksum)
// 		fmt.Println("[*] IP - SrcIP:", ip.SrcIP)
// 		fmt.Println("[*] IP - DstIP:", ip.DstIP)
// 		fmt.Println("[*] IP - Options:", ip.Options)
// 	}

// 	// UDP Layer
// 	udpLayer := packet.Layer(golayers.LayerTypeUDP)
// 	if udpLayer != nil {
// 		udp, _ := udpLayer.(*golayers.UDP)
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("UDP Layer")
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("[*] UDP - SrcPort:", udp.SrcPort.String())
// 		fmt.Println("[*] UDP - DstPort:", udp.DstPort.String())
// 		fmt.Println("[*] UDP - Length:", udp.Length)
// 		fmt.Println("[*] UDP - Checksum:", udp.Checksum)
// 	}

// 	// TCP Layer
// 	tcpLayer := packet.Layer(golayers.LayerTypeTCP)
// 	if tcpLayer != nil {
// 		tcp, _ := tcpLayer.(*golayers.TCP)
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("TCP Layer")
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("[*] TCP - SrcPort:", tcp.SrcPort.String())
// 		fmt.Println("[*] TCP - DstPort:", tcp.DstPort.String())
// 		fmt.Println("[*] TCP - Sequence:", tcp.Seq)
// 		fmt.Println("[*] TCP - Ack number:", tcp.Ack)
// 		fmt.Println("[*] TCP - DataOffset:", tcp.DataOffset)
// 		fmt.Println("[*] TCP - FIN:", tcp.FIN)
// 		fmt.Println("[*] TCP - SYN:", tcp.SYN)
// 		fmt.Println("[*] TCP - RST:", tcp.RST)
// 		fmt.Println("[*] TCP - PSH:", tcp.PSH)
// 		fmt.Println("[*] TCP - ACK:", tcp.ACK)
// 		fmt.Println("[*] TCP - URG:", tcp.URG)
// 		fmt.Println("[*] TCP - ECE:", tcp.ECE)
// 		fmt.Println("[*] TCP - CWR:", tcp.CWR)
// 		fmt.Println("[*] TCP - NS:", tcp.NS)
// 		fmt.Println("[*] TCP - Window Size:", tcp.Window)
// 		fmt.Println("[*] TCP - Checksum:", tcp.Checksum)
// 		fmt.Println("[*] TCP - Urgent:", tcp.Urgent)
// 		fmt.Println("[*] TCP - Options:", tcp.Options)
// 	}
// }

// func PrintIPv6PacketInfo(data []byte) {
// 	packet := gopacket.NewPacket(data, golayers.LayerTypeIPv6, gopacket.Default)

// 	ipv6Layer := packet.Layer(layers.LayerTypeIPv6)
// 	if ipv6Layer != nil {
// 		ipv6, _ := ipv6Layer.(*layers.IPv6)
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("IPv6 Layer")
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("[*] IPv6 - Version:", ipv6.Version)
// 		fmt.Println("[*] IPv6 - TrafficClass:", ipv6.TrafficClass)
// 		fmt.Println("[*] IPv6 - FlowLabel:", ipv6.FlowLabel)
// 		fmt.Println("[*] IPv6 - Length:", ipv6.Length)
// 		fmt.Println("[*] IPv6 - NextHeader:", ipv6.NextHeader)
// 		fmt.Println("[*] IPv6 - HopLimit:", ipv6.HopLimit)
// 		fmt.Println("[*] IPv6 - SrcIP:", ipv6.SrcIP)
// 		fmt.Println("[*] IPv6 - DstIP:", ipv6.DstIP)
// 	}
// 	icmpv6Layer := packet.Layer(layers.LayerTypeICMPv6)
// 	if icmpv6Layer != nil {
// 		icmpv6, _ := icmpv6Layer.(*layers.ICMPv6)
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("ICMPv6 Layer")
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("[*] ICMPv6 - TypeCode:", icmpv6.TypeCode)
// 	}

// }

// func PrintIPv4PacketInfo(data []byte) {
// 	packet := gopacket.NewPacket(data, golayers.LayerTypeIPv4, gopacket.Default)

// 	ipv4Layer := packet.Layer(layers.LayerTypeIPv4)
// 	if ipv4Layer != nil {
// 		ipv4, _ := ipv4Layer.(*layers.IPv4)
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("IPv4 Layer")
// 		fmt.Println("---------------------------------------------")
// 		fmt.Println("[*] IPv4 - Version:", ipv4.Version)
// 		fmt.Println("[*] IPv4 - IHL:", ipv4.IHL)
// 		fmt.Println("[*] IPv4 - Type of Service:", ipv4.TOS)
// 		fmt.Println("[*] IPv4 - Total Length:", ipv4.Length)
// 		fmt.Println("[*] IPv4 - Identification:", ipv4.Id)
// 		fmt.Println("[*] IPv4 - Flags:", ipv4.Flags)
// 		fmt.Println("[*] IPv4 - Fragment Offset:", ipv4.FragOffset)
// 		fmt.Println("[*] IPv4 - Time To Live:", ipv4.TTL)
// 		fmt.Println("[*] IPv4 - Protocol:", ipv4.Protocol)
// 		fmt.Println("[*] IPv4 - Header Checksum:", ipv4.Checksum)
// 		fmt.Println("[*] IPv4 - SrcIP:", ipv4.SrcIP)
// 		fmt.Println("[*] IPv4 - DstIP:", ipv4.DstIP)
// 		fmt.Println("[*] IPv4 - Options:", ipv4.Options)
// 	}
// }
