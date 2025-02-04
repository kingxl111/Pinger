package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/kingxl111/Pinger/backend/internal/config"
)

type Server struct {
	HTTPServer *http.Server
}

func (s *Server) Run(handler http.Handler, cfg *config.Config) error {

	timeout, err := time.ParseDuration(cfg.HTTPServer.Timeout)
	if err != nil {
		log.Fatalf("invalid timeout configuration: %s", err.Error())
	}

	s.HTTPServer = &http.Server{
		Handler:     handler,
		Addr:        cfg.HTTPServer.Address,
		ReadTimeout: timeout,
		IdleTimeout: timeout * 10,
	}

	return s.HTTPServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.HTTPServer.Shutdown(ctx)
}
