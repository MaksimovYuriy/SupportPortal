package models

import (
	"strings"
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
)

type FlowStep struct {
	ID        int       `json:"id"`
	FlowID    int       `json:"flow_id"`
	QueueID   int       `json:"queue_id"`
	Position  int       `json:"position"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *FlowStep) Validate() error {
	if s.QueueID <= 0 || s.Position <= 0 || strings.TrimSpace(s.Name) == "" {
		return apperrors.ErrValidation
	}
	return nil
}
