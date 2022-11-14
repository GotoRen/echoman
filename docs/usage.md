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
