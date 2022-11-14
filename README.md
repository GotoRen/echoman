# 📣 echoman: simple echo server

## Overview
- This is a simple tool where a client generates packets over TCP/IP and sends them to a server.
- The echo server echoes data by swapping src/dst L3 address and L4 port number. (However, some parameters such as TypeCode and CheckSum are recalculated.)

## Documents
- [Information: Network design](./docs/01_information.md)
- [Application: An application that communicates over the overlay network](./docs/02_application.md)
- [Usage: How to start echoman](./docs/03_usage.md)
- [Appendix: About sockets](./docs/04_appendix.md)

## References
- [RFC 792 - Internet Control Message Protocol](https://www.rfc-editor.org/rfc/rfc792)
- [RFC 768 - User Datagram Protocol](https://www.rfc-editor.org/rfc/rfc768)
- [Universal TUN/TAP device driver](https://docs.kernel.org/networking/tuntap.html)
- [Checksum calculation](https://o21o21.hatenablog.jp/entry/2019/01/31/120436)
