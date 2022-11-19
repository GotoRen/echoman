# 🐶 echoman: Simple echo server using Rawsocket

## 🌱 Overview
- これはTCP/IPの上でClientがパケットを生成してServerへ送信するだけの簡単なツールです
- Serverはsrc/dst L2,L3アドレス, L4ポート番号を入れ替えてデータをechoします（ただし、TypeCodeやCheckSumなどの一部パラメータは再計算する）
- Rawsocketを使用してL2レベル（`syscall.ETH_P_IP`）でパケットを操作
  - **LayersType: `LayerTypeEthernet`**

## ⚡️ Generating socket descriptors
### Socket functions
<img src="https://user-images.githubusercontent.com/63791288/194802596-fbed4e9f-4877-45a9-817d-14522b8a5c2c.png" alt="ansible" width="280" height="400" />

### Used function
| args | syscall |
| :--- | :---: |
| Domain | `AF_PACKET` | 
| Type | `SOCK_RAW` | 
| Protocols | `ETH_P_IP` | 

```go
### 送信ソケットの生成
func etherSendSock(intfIndex *net.Interface) (int, error) {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_IP)))
	if err != nil {
		return -1, err
	}

	addr := syscall.SockaddrLinklayer{
		Protocol: htons(syscall.ETH_P_ALL),
		Ifindex:  intfIndex.Index,
	}

	if err := syscall.Bind(fd, &addr); err != nil {
		return -1, err
	}

	return fd, nil
}

### 受信ソケットの生成
func etherRecvSock(intfIndex *net.Interface) (int, error) {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(htons(syscall.ETH_P_IP)))
	if err != nil {
		return -1, err
	}

	addr := syscall.SockaddrLinklayer{
		Protocol: htons(syscall.ETH_P_ALL),
		Ifindex:  intfIndex.Index,
	}
	if err := syscall.Bind(fd, &addr); err != nil {
		return -1, err
	}

	// Received in promiscuous mode
	if err := syscall.SetLsfPromisc(intfIndex.Name, true); err != nil {
		return -1, err
	}

	return fd, nil
}

### Ethernetパケット（L2~）の送信
func SendEtherPacket(fd int, b []byte) error {
	if _, err := syscall.Write(fd, b); err != nil {
		return err
	}
	
	return nil
}
```

## 🤝 Requirement 
- **⚠️WARNING**：https://docs.docker.com/desktop/previous-versions/3.x-mac/#new-9

> First version of docker compose (as an alternative to the existing docker-compose). Supports some basic commands but not the complete functionality of docker-compose yet.

| Languages / Frameworks | Version |
| :--- | ---: |
| Golang | 1.19.1 |
| Docker Desktop | 4.12.0 |
| docker | 20.10.17 |
| docker-compose | 1.29.2 |

## 🚀 Usage
```sh
### envをコピー
$ cp server/.env{.sample,}
$ cp client/.env{.sample,}

### .envの中身を環境に合わせて書き換える
※注意: Docker networkではMACアドレスはランダムに生成されます

### docker-composeを起動
$ make up

### 実行
---
### 1枚目のTerminal: Echoman server を起動
$ make exec/server
# make run

### 2枚目のTerminal: Echoman client を起動
$ make exec/client
# make run
---

### Echoman clientはデフォルトでUDPを喋ります
-> .env を編集することで生成するパケットタイプを切り替える
PACKET_TYPE=[パケットタイプ]
```
| PacketType | Env Value |
| :--- | :---: |
| ICMPv4 | `ICMPV4` |
| UDPv4 | `UDPV4` |

## 📖 Default Information

| Device | Information |
| :--- | ---: |
| Echoman Server IPv4 address | `10.0.3.95` |
| Echoman Server Port number | `30005` |
| Echoman Client IPv4 address | `10.0.3.96` |
| Echoman Client Port number | `30006` |

## 📚 References
- [RFC 792](https://www.rfc-editor.org/rfc/rfc792)
- [RFC 768](https://www.rfc-editor.org/rfc/rfc768)
- [Checksum計算](https://o21o21.hatenablog.jp/entry/2019/01/31/120436)
