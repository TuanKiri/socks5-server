<div align="center">

<h1>SOCKS5 Proxy</h1>

[![license](https://img.shields.io/badge/license-MIT-red.svg)](LICENSE)
[![go version](https://img.shields.io/github/go-mod/go-version/TuanKiri/socks5)](go.mod)
[![docker hub](https://img.shields.io/docker/pulls/tuankiri/socks5)](https://hub.docker.com/r/tuankiri/socks5)
[![docker version](https://img.shields.io/docker/v/tuankiri/socks5)](https://hub.docker.com/r/tuankiri/socks5)
[![docker size](https://img.shields.io/docker/image-size/tuankiri/socks5)](https://hub.docker.com/r/tuankiri/socks5)

An example of a socks5 proxy server based on a package [github.com/TuanKiri/socks5](https://github.com/TuanKiri/socks5).

</div>

## Configuration

Set environment variables as desired.

```sh
# If authentication on the server is required (RFC 1929).
SOCKS5_USER=
SOCKS5_PASSWORD=
# 1 - Connect, 2 - Bind (not support), 3 - UDP Associate.
# example: ALLOWED_COMMANDS=1,3 
# default: permit all commands
ALLOWED_COMMANDS=
# If you need only certain IP addresses to be able to connect to the server.
# example: 192.168.1.243,185.52.141.19
WHITE_LIST_IPS=
# Timeouts for TCP connection. Valid time units are ns, us (or µs), ms, s, m, h.
# example: DIAL_TIMEOUT=5s 
# default: none
DIAL_TIMEOUT=
READ_TIMEOUT=
WRITE_TIMEOUT=
# Timeouts for UDP connection. Valid time units are ns, us (or µs), ms, s, m, h.
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
```

## Docker Compose

```sh
docker-compose up -d socks5
```

## Build

```sh
docker build -t socks5 .
```

## Run

```sh
docker run --env-file .env --network host socks5
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

