package config

import "os"

type Config struct {
	BindAddr string
}

func Load() *Config {
	addr := os.Getenv("BIND_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	return &Config{BindAddr: addr}
}
