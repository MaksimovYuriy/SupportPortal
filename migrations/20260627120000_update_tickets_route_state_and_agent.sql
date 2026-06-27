-- +goose Up
ALTER TABLE tickets
    ALTER COLUMN current_flow_step_id DROP NOT NULL,
    ADD COLUMN assigned_agent_id INT REFERENCES agents(id) ON DELETE SET NULL;

-- +goose Down
ALTER TABLE tickets
    DROP COLUMN assigned_agent_id,
    ALTER COLUMN current_flow_step_id SET NOT NULL;
