package database

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type TicketRepository interface {
	List(ctx context.Context) ([]models.Ticket, error)
	Create(ctx context.Context, ticket *models.Ticket) error
	FindByID(ctx context.Context, id int64) (*models.Ticket, error)
	Update(ctx context.Context, ticket *models.Ticket) error
	Delete(ctx context.Context, id int64) error
}
