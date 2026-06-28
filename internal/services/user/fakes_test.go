package user

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/apperrors"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type fakeUserServiceRepository struct {
	createWithAgentCalled bool
	users                 []*models.User
}

func (r *fakeUserServiceRepository) Create(_ context.Context, user *models.User) error {
	r.users = append(r.users, user)
	return nil
}

func (r *fakeUserServiceRepository) CreateWithAgent(_ context.Context, user *models.User) error {
	r.createWithAgentCalled = true
	r.users = append(r.users, user)
	return nil
}

func (r *fakeUserServiceRepository) List(_ context.Context) ([]*models.User, error) {
	return r.users, nil
}

func (r *fakeUserServiceRepository) FindByID(_ context.Context, id int64) (*models.User, error) {
	for _, user := range r.users {
		if int64(user.ID) == id {
			return user, nil
		}
	}
	return nil, apperrors.ErrNotFound
}
