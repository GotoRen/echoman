FROM golang:1.19

ENV GO111MODULE on

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y iproute2 iputils-ping dnsutils traceroute iperf3 lsof tcpdump wireshark vim
