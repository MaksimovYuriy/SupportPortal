package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type TicketRepository struct {
	db *sql.DB
}

var _ database.TicketRepository = (*TicketRepository)(nil)

func NewTicketRepository(db *sql.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) Create(ctx context.Context, ticket *models.Ticket) error {
	query := `
		INSERT INTO tickets (flow_id, current_flow_step_id, assigned_agent_id, title, description, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`
	row := r.db.QueryRowContext(
		ctx,
		query,
		ticket.FlowID,
		nullableInt64Value(ticket.CurrentFlowStepID),
		nullableInt64Value(ticket.AssignedAgentID),
		ticket.Title,
		ticket.Description,
		ticket.Status,
	)
	if err := row.Scan(
		&ticket.ID,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
	); err != nil {
		return fmt.Errorf("Failed to scan ticket row: %w", err)
	}
	return nil
}

func (r *TicketRepository) List(ctx context.Context) ([]*models.Ticket, error) {
	query := `
		SELECT id, flow_id, current_flow_step_id, assigned_agent_id, title, COALESCE(description, ''), status, created_at, updated_at
		FROM tickets
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute ticket query: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	tickets := make([]*models.Ticket, 0)
	for rows.Next() {
		ticket, err := scanTicket(rows)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan ticket row: %w", err)
		}
		tickets = append(tickets, ticket)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over ticket rows: %w", err)
	}
	return tickets, nil
}

func (r *TicketRepository) FindByID(ctx context.Context, id int64) (*models.Ticket, error) {
	query := `
		SELECT id, flow_id, current_flow_step_id, assigned_agent_id, title, COALESCE(description, ''), status, created_at, updated_at
		FROM tickets
		WHERE id = $1
	`
	row := r.db.QueryRowContext(ctx, query, id)
	ticket, err := scanTicket(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound
		}
		return nil, fmt.Errorf("Failed to scan ticket row: %w", err)
	}
	return ticket, nil
}

func (r *TicketRepository) UpdateState(ctx context.Context, ticket *models.Ticket) error {
	query := `
		UPDATE tickets
		SET current_flow_step_id = $2,
			assigned_agent_id = $3,
			status = $4,
			updated_at = NOW()
		WHERE id = $1
		RETURNING updated_at
	`
	row := r.db.QueryRowContext(
		ctx,
		query,
		ticket.ID,
		nullableInt64Value(ticket.CurrentFlowStepID),
		nullableInt64Value(ticket.AssignedAgentID),
		ticket.Status,
	)
	if err := row.Scan(&ticket.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return apperrors.ErrNotFound
		}
		return fmt.Errorf("Failed to update ticket state: %w", err)
	}
	return nil
}

func (r *TicketRepository) ListByStatus(ctx context.Context, status string, limit int) ([]*models.Ticket, error) {
	query := `
		SELECT id, flow_id, current_flow_step_id, assigned_agent_id, title, COALESCE(description, ''), status, created_at, updated_at
		FROM tickets
		WHERE status = $1
		ORDER BY created_at ASC
		LIMIT $2
	`
	rows, err := r.db.QueryContext(ctx, query, status, limit)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute ticket status query: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	tickets := make([]*models.Ticket, 0)
	for rows.Next() {
		ticket, err := scanTicket(rows)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan ticket row: %w", err)
		}
		tickets = append(tickets, ticket)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over ticket rows: %w", err)
	}
	return tickets, nil
}

type ticketScanner interface {
	Scan(dest ...any) error
}

func scanTicket(scanner ticketScanner) (*models.Ticket, error) {
	var ticket models.Ticket
	var currentFlowStepID sql.NullInt64
	var assignedAgentID sql.NullInt64

	if err := scanner.Scan(
		&ticket.ID,
		&ticket.FlowID,
		&currentFlowStepID,
		&assignedAgentID,
		&ticket.Title,
		&ticket.Description,
		&ticket.Status,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
	); err != nil {
		return nil, err
	}

	ticket.CurrentFlowStepID = nullInt64Ptr(currentFlowStepID)
	ticket.AssignedAgentID = nullInt64Ptr(assignedAgentID)
	return &ticket, nil
}

func nullInt64Ptr(value sql.NullInt64) *int64 {
	if !value.Valid {
		return nil
	}
	return &value.Int64
}

func nullableInt64Value(value *int64) any {
	if value == nil {
		return nil
	}
	return *value
}
