package dto

import (
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type AgentQueueRequest struct {
	AgentID int64 `json:"agent_id"`
	QueueID int64 `json:"queue_id"`
}

func NewAgentQueueRequest(aq *models.AgentQueue) *AgentQueueRequest {
	return &AgentQueueRequest{
		AgentID: aq.AgentID,
		QueueID: aq.QueueID,
	}
}

type AgentQueueResponse struct {
	AgentID   int64     `json:"agent_id"`
	QueueID   int64     `json:"queue_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AgentQueueListResponse struct {
	Queues []*AgentQueueResponse `json:"queues"`
}

func NewAgentQueueResponse(aq *models.AgentQueue) *AgentQueueResponse {
	return &AgentQueueResponse{
		AgentID:   aq.AgentID,
		QueueID:   aq.QueueID,
		CreatedAt: aq.CreatedAt,
	}
}

func NewAgentQueueListResponse(queues []*models.AgentQueue) *AgentQueueListResponse {
	responses := make([]*AgentQueueResponse, len(queues))
	for i, q := range queues {
		responses[i] = NewAgentQueueResponse(q)
	}
	return &AgentQueueListResponse{Queues: responses}
}
