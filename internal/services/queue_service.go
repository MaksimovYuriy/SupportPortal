package services

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type QueueService struct {
	repo database.QueueRepository
}

func NewQueueService(repo database.QueueRepository) *QueueService {
	return &QueueService{
		repo: repo,
	}
}

func (s *QueueService) List(ctx context.Context) ([]models.Queue, error) {
	return s.repo.List(ctx)
}

func (s *QueueService) Create(ctx context.Context, queue *models.Queue) error {
	return s.repo.Create(ctx, queue)
}

func (s *QueueService) FindByID(ctx context.Context, id int64) (*models.Queue, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *QueueService) Update(ctx context.Context, queue *models.Queue) error {
	return s.repo.Update(ctx, queue)
}

func (s *QueueService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
