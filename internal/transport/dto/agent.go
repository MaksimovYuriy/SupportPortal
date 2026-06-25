package dto

import (
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type AgentResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	IsAvailable bool      `json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"`
}

type AgentListResponse struct {
	Agents []*AgentResponse `json:"agents"`
}

func NewAgentResponse(agent *models.Agent) *AgentResponse {
	return &AgentResponse{
		ID:          agent.ID,
		Name:        agent.Name,
		IsAvailable: agent.IsAvailable,
		CreatedAt:   agent.CreatedAt,
		UpdatedAt:   agent.UpdatedAt,
		UserID:      agent.UserID,
	}
}

func NewAgentListResponse(agents []*models.Agent) *AgentListResponse {
	agentResponses := make([]*AgentResponse, len(agents))
	for i, agent := range agents {
		agentResponses[i] = NewAgentResponse(agent)
	}
	return &AgentListResponse{Agents: agentResponses}
}
