package services

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type TicketService struct {
	repo database.TicketRepository
}

func NewTicketService(repo database.TicketRepository) *TicketService {
	return &TicketService{
		repo: repo,
	}
}

func (s *TicketService) List(ctx context.Context) ([]models.Ticket, error) {
	return s.repo.List(ctx)
}

func (s *TicketService) Create(ctx context.Context, ticket *models.Ticket) error {
	return s.repo.Create(ctx, ticket)
}

func (s *TicketService) FindByID(ctx context.Context, id int64) (*models.Ticket, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *TicketService) Update(ctx context.Context, ticket *models.Ticket) error {
	return s.repo.Update(ctx, ticket)
}

func (s *TicketService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
