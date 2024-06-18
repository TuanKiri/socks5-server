# simple-socks5-proxy

A simple example of a socks5 proxy server based on a package [JC5LZiy3HVfV5ux/socks5](https://github.com/JC5LZiy3HVfV5ux/socks5).

## Install

GitHub:
```sh
git clone https://github.com/JC5LZiy3HVfV5ux/simple-socks5-proxy.git
```

Docker Hub
```sh
docker pull jc5lziy3hvfv5ux/socks5-alpine3.19:latest
```

## Configuration

Set environment variables as desired.

```sh
# If authentication on the server is required (RFC 1929).
SOCKS5_USER=
SOCKS5_PASSWORD=

# If you need only certain IP addresses to be able to connect to the server.
# example: 192.168.1.243,185.52.141.19
WHITE_LIST_IPS=
```

## Run

Golang:
```sh
go run ./cmd/main.go
```

Docker run:
```sh
docker run -d -p 1080:1080/tcp jc5lziy3hvfv5ux/socks5-alpine3.19:latest 
```

Docker compose:
```sh
docker-compose up -d socks5
```

## License

All source code is licensed under the MIT License.

