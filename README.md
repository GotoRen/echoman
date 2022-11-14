# 📣 echoman: simple echo server

## Overview
- This is a simple tool where a client generates packets over TCP/IP and sends them to a server.
- The echo server echoes data by swapping src/dst L3 address and L4 port number. (However, some parameters such as TypeCode and CheckSum are recalculated.)

## Requirement
- **⚠️WARNING**：https://docs.docker.com/desktop/previous-versions/3.x-mac/#new-9

> First version of docker compose (as an alternative to the existing docker-compose). Supports some basic commands but not the complete functionality of docker-compose yet.

| Languages / Frameworks | Version |
| :--- | ---: |
| Golang | 1.19.1 |
| Docker Desktop | 4.12.0 |
| docker | 20.10.17 |
| docker-compose | 1.29.2 |

## Documents
- [Usage](./docs/usage.md)
- [Information](./docs/information.md)

## References
- [RFC 792: Internet Control Message Protocol](https://www.rfc-editor.org/rfc/rfc792)
- [RFC 768: User Datagram Protocol](https://www.rfc-editor.org/rfc/rfc768)
- [Universal TUN/TAP device driver](https://docs.kernel.org/networking/tuntap.html)
- [Checksum calculation](https://o21o21.hatenablog.jp/entry/2019/01/31/120436)
