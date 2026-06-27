package models

import (
	"strings"
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
)

const (
	TicketStatusNew               = "new"
	TicketStatusInQueue           = "in_queue"
	TicketStatusInProgress        = "in_progress"
	TicketStatusWaitingTransition = "waiting_transition"
	TicketStatusClosed            = "closed"
)

type Ticket struct {
	ID                int64     `json:"id"`
	FlowID            int64     `json:"flow_id"`
	CurrentFlowStepID *int64    `json:"current_flow_step_id"`
	AssignedAgentID   *int64    `json:"assigned_agent_id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (t *Ticket) Validate() error {
	if t.FlowID <= 0 || strings.TrimSpace(t.Title) == "" {
		return apperrors.ErrValidation
	}
	return nil
}
