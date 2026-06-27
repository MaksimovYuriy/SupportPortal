package dto

import (
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type FlowStepRequest struct {
	QueueID  int    `json:"queue_id"`
	Position int    `json:"position"`
	Name     string `json:"name"`
}

type FlowStepResponse struct {
	ID        int       `json:"id"`
	FlowID    int       `json:"flow_id"`
	QueueID   int       `json:"queue_id"`
	Position  int       `json:"position"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewFlowStepResponse(step *models.FlowStep) *FlowStepResponse {
	return &FlowStepResponse{
		ID:        step.ID,
		FlowID:    step.FlowID,
		QueueID:   step.QueueID,
		Position:  step.Position,
		Name:      step.Name,
		CreatedAt: step.CreatedAt,
		UpdatedAt: step.UpdatedAt,
	}
}

func NewFlowStepListResponse(steps []*models.FlowStep) []*FlowStepResponse {
	if steps == nil {
		return nil
	}
	stepResponses := make([]*FlowStepResponse, len(steps))
	for i, step := range steps {
		stepResponses[i] = NewFlowStepResponse(step)
	}
	return stepResponses
}
