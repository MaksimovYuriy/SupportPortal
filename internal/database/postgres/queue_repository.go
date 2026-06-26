package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type QueueRepository struct {
	db *sql.DB
}

var _ database.QueueRepository = (*QueueRepository)(nil)

func NewQueueRepository(db *sql.DB) *QueueRepository {
	return &QueueRepository{db: db}
}

func (r *QueueRepository) Create(ctx context.Context, queue *models.Queue) error {
	query := `
		INSERT INTO queues (name, is_active)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`
	row := r.db.QueryRowContext(ctx, query, queue.Name, queue.IsActive)
	if err := row.Scan(
		&queue.ID,
		&queue.CreatedAt,
		&queue.UpdatedAt,
	); err != nil {
		return fmt.Errorf("Failed to scan queue row: %w", err)
	}
	return nil
}

func (r *QueueRepository) List(ctx context.Context) ([]*models.Queue, error) {
	query := `
		SELECT id, name, is_active, created_at, updated_at
		FROM queues
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed queue query: %w", err)
	}
	queues := make([]*models.Queue, 0)
	for rows.Next() {
		var queue models.Queue
		if err := rows.Scan(
			&queue.ID,
			&queue.Name,
			&queue.IsActive,
			&queue.CreatedAt,
			&queue.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("Failed to scan queue row: %w", err)
		}
		queues = append(queues, &queue)
	}
	return queues, nil
}

func (r *QueueRepository) FindByID(ctx context.Context, id int64) (*models.Queue, error) {
	query := `
		SELECT id, name, is_active, created_at, updated_at
		FROM queues
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)
	var queue models.Queue
	if err := row.Scan(
		&queue.ID,
		&queue.Name,
		&queue.IsActive,
		&queue.CreatedAt,
		&queue.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("Failed to scan queue row: %w", err)
	}
	return &queue, nil
}
