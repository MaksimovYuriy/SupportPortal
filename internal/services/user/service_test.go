package user

import (
	"context"
	"errors"
	"testing"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

func TestUserServiceCreateUser(t *testing.T) {
	t.Run("creates valid user with agent", func(t *testing.T) {
		repo := &fakeUserServiceRepository{}
		service := NewUserService(repo)
		user := &models.User{
			Email:        "agent@example.com",
			PasswordHash: "hash",
			Role:         models.UserRoleAgent,
			IsActive:     true,
		}

		if err := service.CreateUser(context.Background(), user); err != nil {
			t.Fatalf("CreateUser returned error: %v", err)
		}
		if !repo.createWithAgentCalled {
			t.Fatalf("expected CreateWithAgent to be called")
		}
	})

	t.Run("rejects invalid role", func(t *testing.T) {
		repo := &fakeUserServiceRepository{}
		service := NewUserService(repo)
		user := &models.User{
			Email:        "agent@example.com",
			PasswordHash: "hash",
			Role:         "unknown",
		}

		err := service.CreateUser(context.Background(), user)
		if !errors.Is(err, apperrors.ErrValidation) {
			t.Fatalf("expected validation error, got %v", err)
		}
		if repo.createWithAgentCalled {
			t.Fatalf("expected CreateWithAgent not to be called")
		}
	})
}
