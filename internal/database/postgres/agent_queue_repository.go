package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type AgentQueueRepository struct {
	db *sql.DB
}

var _ database.AgentQueueRepository = (*AgentQueueRepository)(nil)

func NewAgentQueueRepository(db *sql.DB) *AgentQueueRepository {
	return &AgentQueueRepository{db: db}
}

func (r *AgentQueueRepository) Create(ctx context.Context, agentQueue *models.AgentQueue) error {
	query := `
		INSERT INTO agent_queues (agent_id, queue_id)
		VALUES ($1, $2)
		RETURNING created_at
	`
	row := r.db.QueryRowContext(ctx, query, agentQueue.AgentID, agentQueue.QueueID)
	if err := row.Scan(&agentQueue.CreatedAt); err != nil {
		return fmt.Errorf("Failed to create agent queue: %w", err)
	}
	return nil
}

func (r *AgentQueueRepository) Delete(ctx context.Context, agentQueue *models.AgentQueue) error {
	query := `
		DELETE FROM agent_queues
		WHERE agent_id = $1 AND queue_id = $2
	`
	_, err := r.db.ExecContext(ctx, query, agentQueue.AgentID, agentQueue.QueueID)
	if err != nil {
		return fmt.Errorf("Failed to delete agent queue: %w", err)
	}
	return nil
}

func (r *AgentQueueRepository) FindByAgentID(ctx context.Context, id int64) ([]*models.AgentQueue, error) {
	query := `
		SELECT agent_id, queue_id, created_at
		FROM agent_queues
		WHERE agent_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute query: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()
	var queues []*models.AgentQueue
	for rows.Next() {
		var aq models.AgentQueue
		if err := rows.Scan(
			&aq.AgentID,
			&aq.QueueID,
			&aq.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("Failed to scan agent queue row: %w", err)
		}
		queues = append(queues, &aq)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %w", err)
	}
	return queues, nil
}

func (r *AgentQueueRepository) FindByQueueID(ctx context.Context, id int64) ([]*models.AgentQueue, error) {
	query := `
		SELECT agent_id, queue_id, created_at
		FROM agent_queues
		WHERE queue_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute query: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()
	var agents []*models.AgentQueue
	for rows.Next() {
		var aq models.AgentQueue
		if err := rows.Scan(
			&aq.AgentID,
			&aq.QueueID,
			&aq.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("Failed to scan agent queue row: %w", err)
		}
		agents = append(agents, &aq)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %w", err)
	}
	return agents, nil
}

func (r *AgentQueueRepository) Exists(ctx context.Context, agentID int64, queueID int64) (bool, error) {
	query := `
		SELECT 1
		FROM agent_queues
		WHERE agent_id = $1 AND queue_id = $2
	`
	row := r.db.QueryRowContext(ctx, query, agentID, queueID)
	var exists int
	if err := row.Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("Failed to scan agent queue existence: %w", err)
	}
	return exists == 1, nil
}
