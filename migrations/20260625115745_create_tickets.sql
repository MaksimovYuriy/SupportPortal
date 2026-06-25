-- +goose Up
CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    flow_id INT NOT NULL REFERENCES flows(id) ON DELETE RESTRICT,
    current_flow_step_id INT NOT NULL REFERENCES flow_steps(id) ON DELETE RESTRICT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(32) NOT NULL DEFAULT 'new',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_ticket_status CHECK (status IN ('new', 'in_queue', 'in_progress', 'waiting_transition', 'closed'))
);

-- +goose Down
DROP TABLE IF EXISTS tickets;
