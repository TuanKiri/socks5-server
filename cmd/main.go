package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/TuanKiri/socks5"

	"github.com/TuanKiri/simple-socks5-proxy/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	opts := config.NewFromEnv()

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
