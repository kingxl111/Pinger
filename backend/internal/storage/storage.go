package storage

import (
	"context"

	"github.com/kingxl111/Pinger/backend/internal/models"
)

type ContainerManager interface {
	NewContainer(ctx context.Context, container models.ContainerPing) error
	GetContainers(ctx context.Context) ([]models.ContainerPing, error)
	GetContainer(ctx context.Context, containerID int) (models.ContainerPing, error)
}

type Storage struct {
	ContainerManager
}

func NewStorage(db *DB) *Storage {
	return &Storage{
		ContainerManager: NewContainersPG(db),
	}
}
