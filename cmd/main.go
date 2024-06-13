package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/JC5LZiy3HVfV5ux/socks5"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := socks5.New()

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
