package database

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type AgentRepository interface {
	List(ctx context.Context) ([]*models.Agent, error)
	FindByID(ctx context.Context, id int64) (*models.Agent, error)
	CreateForUser(ctx context.Context, user *models.User) error
}
