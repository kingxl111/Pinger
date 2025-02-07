package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-ping/ping"
)

const backendURL = "http://backend:8080"

type Container struct {
	ID int    `json:"id"`
	IP string `json:"ip"`
}

type PingResult struct {
	ContainerID         int       `json:"container_id"`
	PingTime            time.Time `json:"ping_time"`
	LastSuccessPingTime time.Time `json:"last_success_ping"`
}

func getContainers() ([]Container, error) {
	resp, err := http.Get(backendURL + "/get-containers-ping")
	if err != nil {
		return nil, fmt.Errorf("failed to get containers: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Containers []Container `json:"containers"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result.Containers, nil
}

func pingContainer(container Container) (PingResult, error) {
	pinger, err := ping.NewPinger(container.IP)
	if err != nil {
		return PingResult{}, fmt.Errorf("failed to create pinger: %w", err)
	}
	pinger.Count = 3
	pinger.Timeout = 5 * time.Second

	err = pinger.Run()
	if err != nil {
		return PingResult{}, fmt.Errorf("failed to ping: %w", err)
	}

	stats := pinger.Statistics()
	successTime := time.Now()
	if stats.PacketsRecv == 0 {
		successTime = time.Time{}
	}

	return PingResult{
		ContainerID:         container.ID,
		PingTime:            time.Now(),
		LastSuccessPingTime: successTime,
	}, nil
}

func sendPingResult(result PingResult) error {
	data, err := json.Marshal(map[string]interface{}{
		"cont_ping": result,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	resp, err := http.Post(backendURL+"/new-container-ping", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to send ping result: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func main() {
	for {
		containers, err := getContainers()
		if err != nil {
			log.Printf("Error fetching containers: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		for _, container := range containers {
			result, err := pingContainer(container)
			if err != nil {
				log.Printf("Ping failed for %s: %v", container.IP, err)
				continue
			}

			err = sendPingResult(result)
			if err != nil {
				log.Printf("Failed to send result for %s: %v", container.IP, err)
			}
		}

		time.Sleep(30 * time.Second)
	}
}
