package database

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type QueueRepository interface {
	Create(ctx context.Context, queue *models.Queue) error
	List(ctx context.Context) ([]*models.Queue, error)
	FindByID(ctx context.Context, id int64) (*models.Queue, error)
}
