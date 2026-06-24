package database

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type QueueRepository interface {
	List(ctx context.Context) ([]models.Queue, error)
	Create(ctx context.Context, queue *models.Queue) error
	FindByID(ctx context.Context, id int64) (*models.Queue, error)
	FindByName(ctx context.Context, name string) (*models.Queue, error)
	Update(ctx context.Context, queue *models.Queue) error
	Delete(ctx context.Context, id int64) error
}
