package repository

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/domain"
)

type TicketRepository interface {
	List(ctx context.Context) ([]domain.Ticket, error)
	Create(ctx context.Context, ticket *domain.Ticket) error
	FindByID(ctx context.Context, id int64) (*domain.Ticket, error)
	Update(ctx context.Context, ticket *domain.Ticket) error
	Delete(ctx context.Context, id int64) error
}
