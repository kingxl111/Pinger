package service

import (
	"context"

	"github.com/kingxl111/Pinger/backend/internal/models"
	"github.com/kingxl111/Pinger/backend/internal/storage"
)

type Service struct {
	ContainerManagerService
}

type ContainerManagerService interface {
	NewContainerPing(ctx context.Context, container models.ContainerPing) error
	GetContainers(ctx context.Context) ([]models.ContainerPing, error)
	GetContainer(ctx context.Context, containerID int) (models.ContainerPing, error)
}

func NewService(stg *storage.Storage) *Service {
	return &Service{
		ContainerManagerService: NewContainerService(stg),
	}
}
