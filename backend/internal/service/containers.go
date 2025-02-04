package service

import (
	"context"

	"github.com/kingxl111/Pinger/backend/internal/models"
	"github.com/kingxl111/Pinger/backend/internal/storage"
)

var _ ContainerManagerService = (*ContainerService)(nil)

type ContainerService struct {
	stg *storage.Storage
}

func NewContainerService(stg *storage.Storage) *ContainerService {
	return &ContainerService{stg}
}

func (s *ContainerService) NewContainerPing(ctx context.Context, container models.ContainerPing) error {
	return s.stg.NewContainer(ctx, container)
}
func (s *ContainerService) GetContainers(ctx context.Context) ([]models.ContainerPing, error) {
	return s.stg.GetContainers(ctx)
}

func (s *ContainerService) GetContainer(ctx context.Context, containerID int) (models.ContainerPing, error) {
	return s.stg.GetContainer(ctx, containerID)
}
