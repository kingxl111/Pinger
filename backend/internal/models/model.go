package models

import "time"

type ContainerPing struct {
	ID                  int       `json:"id"`
	IP                  string    `json:"ip"`
	PingTime            time.Time `json:"ping_time"`
	LastSuccessPingTime time.Time `json:"last_success_ping"`
}

type CreateContainerPingRequest struct {
	ContPing ContainerPing `json:"cont_ping"`
}

type CreateContainerPingResponse struct {
	Success bool `json:"success"`
}

type GetContainersPingRequest struct{}

type GetContainersPingResponse struct {
	Containers []ContainerPing `json:"containers"`
}

type GetContainerPingRequest struct {
	ContainerID int
}

type GetContainerPingResponse struct {
	ContPing ContainerPing `json:"container"`
}
