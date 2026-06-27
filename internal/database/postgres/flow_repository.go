package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type FlowRepository struct {
	db *sql.DB
}

var _ database.FlowRepository = (*FlowRepository)(nil)

func NewFlowRepository(db *sql.DB) *FlowRepository {
	return &FlowRepository{db: db}
}

func (r *FlowRepository) Create(ctx context.Context, flow *models.Flow) error {
	query := `
		INSERT INTO flows (name, description, is_active)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`
	row := r.db.QueryRowContext(ctx, query, flow.Name, flow.Description, flow.IsActive)
	if err := row.Scan(
		&flow.ID,
		&flow.CreatedAt,
		&flow.UpdatedAt,
	); err != nil {
		return fmt.Errorf("Failed to scan flow row: %w", err)
	}
	return nil
}

func (r *FlowRepository) CreateWithSteps(ctx context.Context, flow *models.Flow, steps []*models.FlowStep) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("Failed to begin flow transaction: %w", err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	flowQuery := `
		INSERT INTO flows (name, description, is_active)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`
	flowRow := tx.QueryRowContext(ctx, flowQuery, flow.Name, flow.Description, flow.IsActive)
	if err := flowRow.Scan(
		&flow.ID,
		&flow.CreatedAt,
		&flow.UpdatedAt,
	); err != nil {
		return fmt.Errorf("Failed to scan flow row: %w", err)
	}

	stepQuery := `
		INSERT INTO flow_steps (flow_id, queue_id, position, name)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`
	for _, step := range steps {
		step.FlowID = flow.ID
		stepRow := tx.QueryRowContext(ctx, stepQuery, step.FlowID, step.QueueID, step.Position, step.Name)
		if err := stepRow.Scan(
			&step.ID,
			&step.CreatedAt,
			&step.UpdatedAt,
		); err != nil {
			return fmt.Errorf("Failed to scan flow step row: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Failed to commit flow transaction: %w", err)
	}
	return nil
}

func (r *FlowRepository) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM flows
		WHERE id = $1
	`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete flow: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Failed to get deleted flow count: %w", err)
	}
	if rowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}

func (r *FlowRepository) List(ctx context.Context) ([]*models.Flow, error) {
	query := `
		SELECT id, name, COALESCE(description, ''), is_active, created_at, updated_at
		FROM flows
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute flow query: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	flows := make([]*models.Flow, 0)
	for rows.Next() {
		var flow models.Flow
		if err := rows.Scan(
			&flow.ID,
			&flow.Name,
			&flow.Description,
			&flow.IsActive,
			&flow.CreatedAt,
			&flow.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("Failed to scan flow row: %w", err)
		}
		flows = append(flows, &flow)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over flow rows: %w", err)
	}
	return flows, nil
}

func (r *FlowRepository) FindByID(ctx context.Context, id int64) (*models.Flow, error) {
	query := `
		SELECT id, name, COALESCE(description, ''), is_active, created_at, updated_at
		FROM flows
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)
	var flow models.Flow
	if err := row.Scan(
		&flow.ID,
		&flow.Name,
		&flow.Description,
		&flow.IsActive,
		&flow.CreatedAt,
		&flow.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound
		}
		return nil, fmt.Errorf("Failed to scan flow row: %w", err)
	}
	return &flow, nil
}
