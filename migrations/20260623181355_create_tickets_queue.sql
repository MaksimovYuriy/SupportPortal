-- +goose Up
CREATE TABLE queues (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE tickets
ADD COLUMN queue_id INT REFERENCES queues(id) ON DELETE RESTRICT;

CREATE INDEX idx_tickets_queue_id ON tickets(queue_id);

-- +goose Down
DROP INDEX IF EXISTS idx_tickets_queue_id;

ALTER TABLE tickets
DROP COLUMN IF EXISTS queue_id;

DROP TABLE IF EXISTS queues;
