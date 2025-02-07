package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kingxl111/Pinger/backend/internal/models"
	"github.com/kingxl111/Pinger/backend/internal/storage"
)

const pingerURL = "http://pinger:8081/get-containers-ping"

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
	req, err := http.NewRequestWithContext(ctx, "GET", pingerURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to pinger: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response from pinger: %d", resp.StatusCode)
	}

	var result struct {
		Containers []models.ContainerPing `json:"containers"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Containers, nil
}
func (s *ContainerService) GetContainer(ctx context.Context, containerID int) (models.ContainerPing, error) {
	return s.stg.GetContainer(ctx, containerID)
}
