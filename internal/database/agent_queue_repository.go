package database

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type AgentQueueRepository interface {
	Create(ctx context.Context, agentQueue *models.AgentQueue) error
	Delete(ctx context.Context, agentQueue *models.AgentQueue) error
	FindByAgentID(ctx context.Context, id int64) ([]*models.AgentQueue, error)
	FindByQueueID(ctx context.Context, id int64) ([]*models.AgentQueue, error)
	Exists(ctx context.Context, agentID int64, queueID int64) (bool, error)
}
