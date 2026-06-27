package models

import (
	"strings"
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
)

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password"`
	Role         string    `json:"role"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

const (
	UserRoleAdmin = "admin"
	UserRoleAgent = "agent"
)

func (u *User) Validate() error {
	if strings.TrimSpace(u.Email) == "" || strings.TrimSpace(u.PasswordHash) == "" {
		return apperrors.ErrValidation
	}
	switch u.Role {
	case UserRoleAdmin, UserRoleAgent:
		return nil
	default:
		return apperrors.ErrValidation
	}
}
