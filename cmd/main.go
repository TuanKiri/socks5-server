package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v11"

	"github.com/TuanKiri/socks5"
)

type config struct {
	User               string        `env:"SOCKS5_USER"`
	Password           string        `env:"SOCKS5_PASSWORD"`
	AllowedCommands    []int         `env:"ALLOWED_COMMANDS" envSeparator:","`
	WhiteListIPS       []string      `env:"WHITE_LIST_IPS" envSeparator:","`
	DialTimeout        time.Duration `env:"DIAL_TIMEOUT"`
	ReadTimeout        time.Duration `env:"READ_TIMEOUT"`
	WriteTimeout       time.Duration `env:"WRITE_TIMEOUT"`
	PacketWriteTimeout time.Duration `env:"PACKET_WRITE_TIMEOUT"`
	MaxPacketSize      int           `env:"MAX_PACKET_SIZE"`
	PublicIP           string        `env:"PUBLIC_IP"`
}

func (c *config) toOptions() []socks5.Option {
	opts := []socks5.Option{
		socks5.WithDialTimeout(c.DialTimeout),
		socks5.WithReadTimeout(c.ReadTimeout),
		socks5.WithWriteTimeout(c.WriteTimeout),
		socks5.WithPacketWriteTimeout(c.PacketWriteTimeout),
		socks5.WithPublicIP(net.ParseIP(c.PublicIP)),
		socks5.WithMaxPacketSize(c.MaxPacketSize),
	}

	if c.User != "" && c.Password != "" {
		opts = append(opts,
			socks5.WithPasswordAuthentication(),
			socks5.WithStaticCredentials(map[string]string{
				c.User: c.Password,
			}),
		)
	}

	if len(c.WhiteListIPS) > 0 {
		whiteListIPS := make([]net.IP, len(c.WhiteListIPS))
		for i, ip := range c.WhiteListIPS {
			whiteListIPS[i] = net.ParseIP(ip)
		}
		opts = append(opts,
			socks5.WithWhiteListIPs(whiteListIPS...),
		)
	}

	if len(c.AllowedCommands) > 0 {
		allowedCommands := make([]socks5.Command, len(c.AllowedCommands))
		for i, command := range c.AllowedCommands {
			allowedCommands[i] = socks5.Command(command)
		}
	}

	return opts
}

func main() {
	var cfg config

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := socks5.New(cfg.toOptions()...)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	if err := srv.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
