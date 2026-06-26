package models

import "time"

type AgentQueue struct {
	AgentID   int64     `json:"agent_id"`
	QueueID   int64     `json:"queue_id"`
	CreatedAt time.Time `json:"created_at"`
}
