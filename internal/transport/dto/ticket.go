package dto

import (
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type CreateTicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TicketResponse struct {
	ID          int64     `json:"id"`
	QueueID     *int64    `json:"queue_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListTicketsResponse struct {
	Tickets []TicketResponse `json:"tickets"`
}

func NewTicketResponse(ticket models.Ticket) TicketResponse {
	return TicketResponse{
		ID:          ticket.ID,
		QueueID:     ticket.QueueID,
		Title:       ticket.Title,
		Description: ticket.Description,
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
	}
}

func NewListTicketsResponse(tickets []models.Ticket) ListTicketsResponse {
	ticketResponses := make([]TicketResponse, len(tickets))
	for i, ticket := range tickets {
		ticketResponses[i] = NewTicketResponse(ticket)
	}
	return ListTicketsResponse{
		Tickets: ticketResponses,
	}
}
