package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/liumkssq/eGO/internal/server"
	"github.com/liumkssq/eGO/pkg/config"
)

func main() {
	cfg := config.Load() // 简单读取 env/config
	s := server.New(cfg)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := s.Start(ctx); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
