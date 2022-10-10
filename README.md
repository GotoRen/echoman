# echoman

## 🌱 Overview
- これはTCP/IPの上でClientがパケットを生成してServerへ送信するだけの簡単なツールです
- Serverはsrc/dst L2,L3アドレス, L4ポート番号を入れ替えてデータをechoします（ただし、TypeCodeやCheckSumなどの一部パラメータは再計算する）
- Rawsocketを使用してL2レベル（`syscall.ETH_P_IP`）でパケットを操作
  - **LayersType: `LayerTypeEthernet`**

## ⚡️ Generating socket descriptors
### Socket functions
<img src="https://user-images.githubusercontent.com/63791288/194802596-fbed4e9f-4877-45a9-817d-14522b8a5c2c.png" alt="ansible" width="280" height="400" />

### Used function: client
| args | syscall |
| :--- | :---: |
| Domain | `AF_PACKET` | 
| Type | `SOCK_RAW` | 
| Protocols | `ETH_P_IP` | 

```go
### Client側: 受信ソケットの生成
func EtherSendSock(intfIndex *net.Interface) (int, error) {
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
```

### Used function: server

## 🚀 Usage
```sh
### envをコピー
$ cp server/.env{.sample,}
$ cp client/.env{.sample,}

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
-> .env を編集することで生成するパケットを（ICMPv4, UDP）切り替えることができます
```

## 📖 Information

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
