package models

import (
	"strings"
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
)

type Flow struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (f *Flow) Validate() error {
	if strings.TrimSpace(f.Name) == "" {
		return apperrors.ErrValidation
	}
	return nil
}
