package queue

import (
	"context"
	"errors"
	"testing"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

func TestQueueServiceCreate(t *testing.T) {
	t.Run("creates valid queue", func(t *testing.T) {
		repo := &fakeQueueServiceRepository{}
		service := NewQueueService(repo)
		queue := &models.Queue{Name: "Support", IsActive: true}

		if err := service.Create(context.Background(), queue); err != nil {
			t.Fatalf("Create returned error: %v", err)
		}
		if !repo.created {
			t.Fatalf("expected repository Create to be called")
		}
	})

	t.Run("rejects empty name", func(t *testing.T) {
		repo := &fakeQueueServiceRepository{}
		service := NewQueueService(repo)

		err := service.Create(context.Background(), &models.Queue{Name: " "})
		if !errors.Is(err, apperrors.ErrValidation) {
			t.Fatalf("expected validation error, got %v", err)
		}
		if repo.created {
			t.Fatalf("expected repository Create not to be called")
		}
	})
}
