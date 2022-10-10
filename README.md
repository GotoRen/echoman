# echoman

## Overview
- Rawsocketを使用してL2レベルでTCP/IPスイートを使用してパケットを操作
  - **LayersType: `LayerTypeEthernet`**
- Clientでパケットを生成してServerへ送信
- Serverはsrc/dst L2,L3アドレスを入れ替えてデータをechoします（ただし、TypeCodeやCheckSumなどの一部パラメータは再計算する）

## Generating socket descriptors
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

## Usage
```
### envをコピー
$ cp server/.env{.sample,}
$ cp client/.env{.sample,}

### docker-composeを起動
$ make up

### 各コンテナに入る
$ make exec/server
$ make exec/client

### 実行（各コンテナ内）
# make run
# make run
```
