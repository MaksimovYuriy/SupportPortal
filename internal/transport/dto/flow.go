package dto

import (
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type FlowRequest struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	IsActive    bool               `json:"is_active"`
	Steps       []*FlowStepRequest `json:"steps"`
}

type FlowResponse struct {
	ID          int                 `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	IsActive    bool                `json:"is_active"`
	Steps       []*FlowStepResponse `json:"steps,omitempty"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

type FlowListResponse struct {
	Flows []*FlowResponse `json:"flows"`
}

func NewFlowResponse(flow *models.Flow) *FlowResponse {
	return NewFlowResponseWithSteps(flow, nil)
}

func NewFlowResponseWithSteps(flow *models.Flow, steps []*models.FlowStep) *FlowResponse {
	return &FlowResponse{
		ID:          flow.ID,
		Name:        flow.Name,
		Description: flow.Description,
		IsActive:    flow.IsActive,
		Steps:       NewFlowStepListResponse(steps),
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
