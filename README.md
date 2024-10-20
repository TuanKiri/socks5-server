<div align="center">

<h1>SOCKS5 Proxy</h1>

[![license](https://img.shields.io/badge/license-MIT-red.svg)](LICENSE)
[![go version](https://img.shields.io/github/go-mod/go-version/TuanKiri/socks5)](go.mod)
[![docker hub](https://img.shields.io/docker/pulls/tuankiri/socks5)](https://hub.docker.com/r/tuankiri/socks5)
[![docker version](https://img.shields.io/docker/v/tuankiri/socks5)](https://hub.docker.com/r/tuankiri/socks5)
[![docker size](https://img.shields.io/docker/image-size/tuankiri/socks5)](https://hub.docker.com/r/tuankiri/socks5)

</div>

## Overview

SOCKS5 proxy server based on a package [socks5](https://github.com/TuanKiri/socks5).


## Features

- SOCKS5 CONNECT command
- SOCKS5 UDP ASSOCIATE command
- IPv4, IPv6, FQDN
- Username/Password Authentication
- ACL (Access Control List)

## Configuration

Set environment variables:

```sh
# If authentication on the server is required (RFC 1929).
# default: no authentication required
SOCKS5_USER=
SOCKS5_PASSWORD=
# 1 - Connect, 2 - Bind (not support), 3 - UDP Associate.
# example: ALLOWED_COMMANDS=1,3 
# default: permit all commands
ALLOWED_COMMANDS=
# If you need only certain IP addresses to be able to connect to the server.
# example: 192.168.1.243,185.52.141.19
# default: none
WHITE_LIST_IPS=
# Timeouts for TCP connection. 
# Valid time units are ns, us (or µs), ms, s, m, h.
# example: DIAL_TIMEOUT=5s 
# default: none
DIAL_TIMEOUT=
READ_TIMEOUT=
WRITE_TIMEOUT=
# Timeouts for UDP connection. 
# Valid time units are ns, us (or µs), ms, s, m, h.
# example: PACKET_WRITE_TIMEOUT=5s
# default: none
PACKET_WRITE_TIMEOUT=
# Maximum size in bytes for the datagram to be read from the socket.
# example: MAX_PACKET_SIZE=65507
# default: 1500
MAX_PACKET_SIZE=
# IP address that is visible on the external Internet,
# accessible to users outside the local network and will be sent to clients in
# response to a connection request.
# example: 192.168.1.243
# default: 127.0.0.1
PUBLIC_IP=
# Enable logging
# default: FALSE
LOGGING=
# Cleanup period for UDP packets that do not receive a response. 
# Valid time units are ns, us (or µs), ms, s, m, h.
# example: NAT_CLEANUP_PERIOD=5s
# default: none
NAT_CLEANUP_PERIOD=
# Lifetime of each udp packet. 
# Valid time units are ns, us (or µs), ms, s, m, h.
# example: TTLPacket=5s
# default: none
TTLPacket=
```

## Pull from GitHub Container Registry

`linux/amd64`

```sh
docker pull ghcr.io/tuankiri/socks5:latest
```

## Docker Compose

```sh
docker-compose up -d socks5
```

## Test

```sh
curl -x socks5://127.0.0.1:1080 http://example.com
```

## Contributing
Feel free to open tickets or send pull requests with improvements. Thanks in advance for your help!

Please follow the [contribution guidelines](.github/CONTRIBUTING.md).


## License

All source code is licensed under the MIT License.

