package models

import (
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
)

type AgentQueue struct {
	AgentID   int64     `json:"agent_id"`
	QueueID   int64     `json:"queue_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (aq *AgentQueue) Validate() error {
	if aq.AgentID <= 0 || aq.QueueID <= 0 {
		return apperrors.ErrValidation
	}
	return nil
}
