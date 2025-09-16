package server

import (
	"context"
	"net/http"
	"time"

	"github.com/liumkssq/eGO/pkg/config"
)

type Server struct {
	srv *http.Server
}

func New(cfg *config.Config) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	return &Server{
		srv: &http.Server{
			Addr:         cfg.BindAddr,
			Handler:      mux,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		_ = s.srv.ListenAndServe()
	}()
	<-ctx.Done()
	return s.srv.Shutdown(context.Background())
}
