# simple-socks5-proxy

A simple example of a socks5 proxy server based on a package [TuanKiri/socks5](https://github.com/TuanKiri/socks5).

## Install

```sh
git clone https://github.com/TuanKiri/simple-socks5-proxy.git
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

## Build

Set your linux user's `uid` and `gid`.

```sh
docker build --build-arg USER_UID=${USER_UID:-1000} --build-arg USER_GID=${USER_GID:-1000} -t socks5 .
```

## Run

```sh
docker run -d -v ./runtime/logs/socks5:/var/log/socks5 -p 1080:1080/tcp socks5 
```

## Test

```sh
curl -x socks5://127.0.0.1:1080 http://example.com
```

## License

All source code is licensed under the MIT License.

