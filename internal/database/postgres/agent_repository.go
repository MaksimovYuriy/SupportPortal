package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type AgentRepository struct {
	db *sql.DB
}

var _ database.AgentRepository = (*AgentRepository)(nil)

func NewAgentRepository(db *sql.DB) *AgentRepository {
	return &AgentRepository{db: db}
}

func (r *AgentRepository) List(ctx context.Context) ([]*models.Agent, error) {
	query := `
		SELECT id, name, is_available, user_id, created_at, updated_at
		FROM agents
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute query: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()
	agents := make([]*models.Agent, 0)
	for rows.Next() {
		var agent models.Agent
		if err := rows.Scan(
			&agent.ID,
			&agent.Name,
			&agent.IsAvailable,
			&agent.UserID,
			&agent.CreatedAt,
			&agent.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("Failed to scan agent: %w", err)
		}
		agents = append(agents, &agent)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %w", err)
	}
	return agents, nil
}

func (r *AgentRepository) FindByID(ctx context.Context, id int64) (*models.Agent, error) {
	query := `
		SELECT id, name, is_available, user_id, created_at, updated_at
		FROM agents
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)
	var agent models.Agent
	if err := row.Scan(
		&agent.ID,
		&agent.Name,
		&agent.IsAvailable,
		&agent.UserID,
		&agent.CreatedAt,
		&agent.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound
		}
		return nil, fmt.Errorf("Failed to scan agent row: %w", err)
	}
	return &agent, nil
}

func (r *AgentRepository) CreateForUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO agents (name, is_available, user_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	row := r.db.QueryRowContext(ctx, query, user.Email, user.IsActive, user.ID)
	var id int
	if err := row.Scan(&id); err != nil {
		return fmt.Errorf("Failed creating agent for user: %w", err)
	}
	return nil
}
