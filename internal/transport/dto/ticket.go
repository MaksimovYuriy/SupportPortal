package dto

import (
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type TicketRequest struct {
	FlowID      int64  `json:"flow_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TicketResponse struct {
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

type TicketListResponse struct {
	Tickets []*TicketResponse `json:"tickets"`
}

func NewTicketResponse(ticket *models.Ticket) *TicketResponse {
	return &TicketResponse{
		ID:                ticket.ID,
		FlowID:            ticket.FlowID,
		CurrentFlowStepID: ticket.CurrentFlowStepID,
		AssignedAgentID:   ticket.AssignedAgentID,
		Title:             ticket.Title,
		Description:       ticket.Description,
		Status:            ticket.Status,
		CreatedAt:         ticket.CreatedAt,
		UpdatedAt:         ticket.UpdatedAt,
	}
}

func NewTicketListResponse(tickets []*models.Ticket) *TicketListResponse {
	ticketResponses := make([]*TicketResponse, len(tickets))
	for i, ticket := range tickets {
		ticketResponses[i] = NewTicketResponse(ticket)
	}
	return &TicketListResponse{Tickets: ticketResponses}
}
