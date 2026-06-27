package models

import (
	"strings"
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
)

type Queue struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queue) Validate() error {
	if strings.TrimSpace(q.Name) == "" {
		return apperrors.ErrValidation
	}
	return nil
}
