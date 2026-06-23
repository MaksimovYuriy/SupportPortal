package repositories

import (
	"context"
	"database/sql"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type PostgresQueueRepository struct {
	db *sql.DB
}

var _ database.QueueRepository = (*PostgresQueueRepository)(nil)

func NewPostgresQueueRepository(db *sql.DB) *PostgresQueueRepository {
	return &PostgresQueueRepository{
		db: db,
	}
}

func (r *PostgresQueueRepository) List(ctx context.Context) ([]models.Queue, error) {
	query := `
		SELECT id, name, is_active, created_at, updated_at
		FROM queues
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	queues := make([]models.Queue, 0)
	for rows.Next() {
		var queue models.Queue
		if err := rows.Scan(
			&queue.ID,
			&queue.Name,
			&queue.IsActive,
			&queue.CreatedAt,
			&queue.UpdatedAt,
		); err != nil {
			return nil, err
		}

		queues = append(queues, queue)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return queues, nil
}

func (r *PostgresQueueRepository) Create(ctx context.Context, queue *models.Queue) error {
	query := `
		INSERT INTO queues (name, is_active)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`

	row := r.db.QueryRowContext(ctx, query, queue.Name, queue.IsActive)
	if err := row.Scan(&queue.ID, &queue.CreatedAt, &queue.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (r *PostgresQueueRepository) FindByID(ctx context.Context, id int64) (*models.Queue, error) {
	query := `
		SELECT id, name, is_active, created_at, updated_at
		FROM queues
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	queue := &models.Queue{}
	if err := row.Scan(&queue.ID, &queue.Name, &queue.IsActive, &queue.CreatedAt, &queue.UpdatedAt); err != nil {
		return nil, err
	}
	return queue, nil
}

func (r *PostgresQueueRepository) Update(ctx context.Context, queue *models.Queue) error {
	query := `
		UPDATE queues
		SET name = $2, is_active = $3, updated_at = NOW()
		WHERE id = $1
		RETURNING updated_at
	`

	row := r.db.QueryRowContext(ctx, query, queue.ID, queue.Name, queue.IsActive)
	if err := row.Scan(&queue.UpdatedAt); err != nil {
		return err
	}
	return nil
}

func (r *PostgresQueueRepository) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM queues
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
