package handlers

import (
	"context"
	"net/http"

	"github.com/kingxl111/Pinger/backend/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) NewRouter(ctx *context.Context, log *logrus.Logger, env string) http.Handler {
	log.Info(
		"STARTING PingerAPP",
		log.WithFields(logrus.Fields{
			"env":     env,
			"version": "1.0",
		}),
	)
	log.Debug("debug messages are enabled")

	router := chi.NewRouter()
	router.Use(NewLogger(log))
	router.Route("/", func(r chi.Router) {
		r.Post("/new-container-ping", h.CreateContainerPing(*ctx, log))
		r.Get("/get-containers-ping", h.GetContainersPing(*ctx, log))
	})

	return router
}
