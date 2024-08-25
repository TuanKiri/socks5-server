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

func main() {
	var cfg config

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	opts := []socks5.Option{
		socks5.WithDialTimeout(cfg.DialTimeout),
		socks5.WithReadTimeout(cfg.ReadTimeout),
		socks5.WithWriteTimeout(cfg.WriteTimeout),
		socks5.WithPacketWriteTimeout(cfg.PacketWriteTimeout),
		socks5.WithPublicIP(net.ParseIP(cfg.PublicIP)),
		socks5.WithMaxPacketSize(cfg.MaxPacketSize),
	}

	if cfg.User != "" && cfg.Password != "" {
		opts = append(opts,
			socks5.WithPasswordAuthentication(),
			socks5.WithStaticCredentials(map[string]string{
				cfg.User: cfg.Password,
			}),
		)
	}

	if len(cfg.WhiteListIPS) > 0 {
		whiteListIPS := make([]net.IP, len(cfg.WhiteListIPS))
		for i, ip := range cfg.WhiteListIPS {
			whiteListIPS[i] = net.ParseIP(ip)
		}
		opts = append(opts,
			socks5.WithWhiteListIPs(whiteListIPS...),
		)
	}

	if len(cfg.AllowedCommands) > 0 {
		allowedCommands := make([]socks5.Command, len(cfg.AllowedCommands))
		for i, command := range cfg.AllowedCommands {
			allowedCommands[i] = socks5.Command(command)
		}
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := socks5.New(opts...)

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
