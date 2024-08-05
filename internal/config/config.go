package config

import (
	"net"
	"os"
	"strings"
	"time"

	"github.com/TuanKiri/socks5"
)

func NewFromEnv() []socks5.Option {
	opts := []socks5.Option{
		socks5.WithReadTimeout(60 * time.Second),
		socks5.WithWriteTimeout(60 * time.Second),
		socks5.WithDialTimeout(60 * time.Second),
		socks5.WithAllowCommands(
			socks5.Connect,
		),
	}

	if ok, credentials := getStaticCredentials(); ok {
		opts = append(opts,
			socks5.WithPasswordAuthentication(),
			socks5.WithStaticCredentials(credentials),
		)
	}

	if ok, whiteListIPs := getWhiteListIPs(); ok {
		opts = append(opts,
			socks5.WithWhiteListIPs(whiteListIPs...),
		)
	}

	return opts
}

func getStaticCredentials() (bool, map[string]string) {
	user := os.Getenv("SOCKS5_USER")
	password := os.Getenv("SOCKS5_PASSWORD")

	return user != "" && password != "", map[string]string{user: password}
}

func getWhiteListIPs() (bool, []net.IP) {
	var whiteListIPs []net.IP

	if whiteList := os.Getenv("WHITE_LIST_IPS"); whiteList != "" {
		ips := strings.Split(whiteList, ",")

		for _, ip := range ips {
			whiteListIPs = append(whiteListIPs, net.ParseIP(ip))
		}
	}

	return len(whiteListIPs) > 0, whiteListIPs
}
