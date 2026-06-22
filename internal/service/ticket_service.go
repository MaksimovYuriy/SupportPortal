package service

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/domain"
	"github.com/MaksimovYuriy/SupportPortal/internal/repository"
)

type TicketService struct {
	repo repository.TicketRepository
}

func NewTicketService(repo repository.TicketRepository) *TicketService {
	return &TicketService{
		repo: repo,
	}
}

func (s *TicketService) List(ctx context.Context) ([]domain.Ticket, error) {
	return s.repo.List(ctx)
}

func (s *TicketService) Create(ctx context.Context, ticket *domain.Ticket) error {
	return s.repo.Create(ctx, ticket)
}

func (s *TicketService) FindByID(ctx context.Context, id int64) (*domain.Ticket, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *TicketService) Update(ctx context.Context, ticket *domain.Ticket) error {
	return s.repo.Update(ctx, ticket)
}

func (s *TicketService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
