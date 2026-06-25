-- +goose Up
CREATE TABLE flow_steps (
    id SERIAL PRIMARY KEY,
    flow_id INT NOT NULL REFERENCES flows(id) ON DELETE CASCADE,
    queue_id INT NOT NULL REFERENCES queues(id) ON DELETE RESTRICT,
    position INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (flow_id, position),
    CHECK (position > 0)
);

-- +goose Down
DROP TABLE IF EXISTS flow_steps;
