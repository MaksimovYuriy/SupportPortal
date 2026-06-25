package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type PostgresTicketRepository struct {
	db *sql.DB
}

var _ database.TicketRepository = (*PostgresTicketRepository)(nil)

func NewPostgresTicketRepository(db *sql.DB) *PostgresTicketRepository {
	return &PostgresTicketRepository{
		db: db,
	}
}

func (r *PostgresTicketRepository) List(ctx context.Context) ([]models.Ticket, error) {
	query := `
		SELECT id, queue_id, title, description, created_at, updated_at
		FROM tickets
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute ticket query: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	tickets := make([]models.Ticket, 0)
	for rows.Next() {
		var ticket models.Ticket
		if err := rows.Scan(
			&ticket.ID,
			&ticket.QueueID,
			&ticket.Title,
			&ticket.Description,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan ticket row: %w", err)
		}

		tickets = append(tickets, ticket)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ticket row iteration error: %w", err)
	}
	return tickets, nil
}

func (r *PostgresTicketRepository) Create(ctx context.Context, ticket *models.Ticket) error {
	query := `
		INSERT INTO tickets (title, description)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`

	row := r.db.QueryRowContext(ctx, query, ticket.Title, ticket.Description)
	if err := row.Scan(
		&ticket.ID,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
	); err != nil {
		return fmt.Errorf("failed to scan row: %w", err)
	}
	return nil
}

func (r *PostgresTicketRepository) FindByID(ctx context.Context, id int64) (*models.Ticket, error) {
	query := `
		SELECT id, queue_id, title, description, created_at, updated_at
		FROM tickets
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	ticket := &models.Ticket{}
	if err := row.Scan(
		&ticket.ID,
		&ticket.QueueID,
		&ticket.Title,
		&ticket.Description,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to find ticket by id: %w", err)
	}
	return ticket, nil
}

func (r *PostgresTicketRepository) Update(ctx context.Context, ticket *models.Ticket) error {
	query := `
		UPDATE tickets
		SET	queue_id = $2, title = $3, description = $4, updated_at = NOW()
		WHERE id = $1
		RETURNING updated_at
	`
	row := r.db.QueryRowContext(ctx, query, ticket.ID, ticket.QueueID, ticket.Title, ticket.Description)
	if err := row.Scan(&ticket.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return apperrors.ErrNotFound
		}
		return fmt.Errorf("failed to scan row: %w", err)
	}
	return nil
}

func (r *PostgresTicketRepository) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM tickets
		WHERE id = $1
	`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}
