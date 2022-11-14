# Usage: How to start echoman

## Requirement
- **⚠️WARNING**：https://docs.docker.com/desktop/previous-versions/3.x-mac/#new-9

> First version of docker compose (as an alternative to the existing docker-compose). Supports some basic commands but not the complete functionality of docker-compose yet.

| Languages / Frameworks | Version |
| :--- | ---: |
| Golang | 1.19.1 |
| Docker Desktop | 4.12.0 |
| docker | 20.10.17 |
| docker-compose | 1.29.2 |

## Usage
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
```

## UDPパケットの生成
- デフォルトでUDPパケットを生成しない場合、コメントアウトして下さい。
> ./echoman/client/exec/run.go

```go
for {
    <-t.C
    device.NewChorusUDPPacket() // If you want to generate UDP packets, please uncomment here.
}
```
