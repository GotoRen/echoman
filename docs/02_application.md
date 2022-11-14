# Application: An application that communicates over the overlay network

## 🎶 Chorus: echoman test application
- Chorusは、echomanを使用したオーバーレイネットワーク通信をテストするための検証アプリケーションです。
- ポート番号は、それぞれ仮想インターフェースにバインドされた`30910 (UDP)`番を使用する。
- UDP Request packet (`Ping`) をサーバに送信すると、クライアントにUDP Response packet (`Pong`) が返される。

## Design
```
Client                                             Server
+-------------------------------------+            +-------------------------------------+
| +--------------+                    |            |                    +--------------+ |     
| | App (Chorus) |                    |            |                    | App (Chorus) | |       
| +--------------+                    |            |                    +--------------+ |     
| 198.18.x.x:30910                    |            |                   198.18.9.10:30910 |     
|        |                            |            |                            |        |
| +--------------+   +--------------+ |            | +--------------+   +--------------+ |
| |   TUN/TAP    |   |   Real I/F   | | UDP tunnel | |   Real I/F   |   |    TUN/TAP   | |
| |  198.18.x.x  |---|   10.0.3.96  |=|============|=|  10.0.3.95   |---|  198.18.9.10 | |
| +--------------+   +--------------+ |            | +--------------+   +--------------+ |
|                       0.0.0.0:30000 |            | 0.0.0.0:30000                       |
+-------------------------------------+            +-------------------------------------+ 
```

## Warning
```go
/*************************************************************************************
 * README: description for Chorus.app *
**************************************************************************************
 * Checking the TUN -> Application packet flow using source code is complicated.
 *   - For the time being, I will check with wireshark.app.
 * Thefore, if the write to TUN succeeds, we generate and return a response message.
 *   - Write the message generated at this time directly to the Real interface.
 * ### Judgment method ###
 *   - If the destination is "198.18.9.10:30910", judge it as chorus.app and return the message.
 *   - And, return a response to the UDP packet received from the client.
*************************************************************************************/
if net.ParseIP(device.Tun.VIP).To4().Equal(dstIP) && golayers.UDPPort(uint16(device.ChorusPort)) == dstPort {
	logger.LogDebug("Receive chorus message", "chrous", "success")
	res := chorus.GenerateUDPResponsePacket(buf)
	if _, err := device.Peer.ConnUDP.WriteToUDP(res, &device.Peer.PeerEndPoint); err != nil {
		logger.LogErr("[Failed] Send chorus message", "error", err)
	} else {
		logger.LogDebug("Send chorus message", "chrous", "success")
	}
}
```
