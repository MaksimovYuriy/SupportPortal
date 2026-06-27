package database

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type TicketRepository interface {
	Create(ctx context.Context, ticket *models.Ticket) error
	List(ctx context.Context) ([]*models.Ticket, error)
	FindByID(ctx context.Context, id int64) (*models.Ticket, error)
	UpdateState(ctx context.Context, ticket *models.Ticket) error
	ListByStatus(ctx context.Context, status string, limit int) ([]*models.Ticket, error)
}
