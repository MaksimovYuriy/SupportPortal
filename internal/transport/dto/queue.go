package dto

import (
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type QueueResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type QueueListResponse struct {
	Queues []*QueueResponse `json:"queues"`
}

func NewQueueResponse(queue *models.Queue) *QueueResponse {
	return &QueueResponse{
		ID:        queue.ID,
		Name:      queue.Name,
		IsActive:  queue.IsActive,
		CreatedAt: queue.CreatedAt,
		UpdatedAt: queue.UpdatedAt,
	}
}

func NewQueueListResponse(queues []*models.Queue) *QueueListResponse {
	queueResponses := make([]*QueueResponse, len(queues))
	for i, queue := range queues {
		queueResponses[i] = NewQueueResponse(queue)
	}
	return &QueueListResponse{Queues: queueResponses}
}

type QueueRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func NewQueueRequest(queue *models.Queue) *QueueRequest {
	return &QueueRequest{
		Name:     queue.Name,
		IsActive: queue.IsActive,
	}
}
