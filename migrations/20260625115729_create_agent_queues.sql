-- +goose Up
CREATE TABLE agent_queues (
    agent_id INT NOT NULL REFERENCES agents(id) ON DELETE CASCADE,
    queue_id INT NOT NULL REFERENCES queues(id) ON DELETE CASCADE,
    PRIMARY KEY (agent_id, queue_id),
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS agent_queues;
