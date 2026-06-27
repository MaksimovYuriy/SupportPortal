package dto

import (
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type FlowRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type FlowResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FlowListResponse struct {
	Flows []*FlowResponse `json:"flows"`
}

func NewFlowResponse(flow *models.Flow) *FlowResponse {
	return &FlowResponse{
		ID:          flow.ID,
		Name:        flow.Name,
		Description: flow.Description,
		IsActive:    flow.IsActive,
		CreatedAt:   flow.CreatedAt,
		UpdatedAt:   flow.UpdatedAt,
	}
}

func NewFlowListResponse(flows []*models.Flow) *FlowListResponse {
	flowResponses := make([]*FlowResponse, len(flows))
	for i, flow := range flows {
		flowResponses[i] = NewFlowResponse(flow)
	}
	return &FlowListResponse{Flows: flowResponses}
}
