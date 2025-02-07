package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kingxl111/Pinger/backend/internal/models"
	"github.com/sirupsen/logrus"

	"github.com/go-chi/render"
)

func (h *Handler) CreateContainerPing(ctx context.Context, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.CreateContainerPing"
		var req models.CreateContainerPingRequest

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error(op, "failed to decode request", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		defer func() {
			if err := r.Body.Close(); err != nil {
				log.Error(err.Error())
			}
		}()

		fmt.Println(req.ContPing)
		err = h.services.ContainerManagerService.NewContainerPing(ctx, req.ContPing)
		if err != nil {
			log.Error(op, "failed to create container ping", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Info(op)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		jsonResponse := models.CreateContainerPingResponse{
			Success: true,
		}
		if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
			log.Error(err.Error())
		}
	}
}

func (h *Handler) GetContainersPing(ctx context.Context, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		containers, err := h.services.ContainerManagerService.GetContainers(ctx)
		if err != nil {
			log.Error("failed to fetch containers: ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		response := models.GetContainersPingResponse{Containers: containers}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
