package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type FlowStepRepository struct {
	db *sql.DB
}

var _ database.FlowStepRepository = (*FlowStepRepository)(nil)

func NewFlowStepRepository(db *sql.DB) *FlowStepRepository {
	return &FlowStepRepository{db: db}
}

func (r *FlowStepRepository) ListByFlowID(ctx context.Context, flowID int64) ([]*models.FlowStep, error) {
	query := `
		SELECT id, flow_id, queue_id, position, name, created_at, updated_at
		FROM flow_steps
		WHERE flow_id = $1
		ORDER BY position ASC
	`
	rows, err := r.db.QueryContext(ctx, query, flowID)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute flow step query: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	steps := make([]*models.FlowStep, 0)
	for rows.Next() {
		var step models.FlowStep
		if err := rows.Scan(
			&step.ID,
			&step.FlowID,
			&step.QueueID,
			&step.Position,
			&step.Name,
			&step.CreatedAt,
			&step.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("Failed to scan flow step row: %w", err)
		}
		steps = append(steps, &step)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over flow step rows: %w", err)
	}
	return steps, nil
}

func (r *FlowStepRepository) FindByID(ctx context.Context, id int64) (*models.FlowStep, error) {
	query := `
		SELECT id, flow_id, queue_id, position, name, created_at, updated_at
		FROM flow_steps
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)
	var step models.FlowStep
	if err := row.Scan(
		&step.ID,
		&step.FlowID,
		&step.QueueID,
		&step.Position,
		&step.Name,
		&step.CreatedAt,
		&step.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound
		}
		return nil, fmt.Errorf("Failed to scan flow step row: %w", err)
	}
	return &step, nil
}
