package dto

import "github.com/MaksimovYuriy/SupportPortal/internal/models"

type CreateQueueRequest struct {
	Name string `json:"name"`
}

type UpdateQueueRequest struct {
	Name string `json:"name"`
}

type QueueResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ListQueuesResponse struct {
	Queues []QueueResponse `json:"queues"`
}

func NewQueueResponse(queue models.Queue) QueueResponse {
	return QueueResponse{
		ID:   queue.ID,
		Name: queue.Name,
	}
}

func NewListQueuesResponse(queues []models.Queue) ListQueuesResponse {
	queueResponses := make([]QueueResponse, len(queues))
	for i, queue := range queues {
		queueResponses[i] = NewQueueResponse(queue)
	}
	return ListQueuesResponse{
		Queues: queueResponses,
	}
}
