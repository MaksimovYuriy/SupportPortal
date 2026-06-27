package services

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type UserService struct {
	userRepo database.UserRepository
}

func NewUserService(userRepo database.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if err := s.userRepo.CreateWithAgent(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]*models.User, error) {
	users, err := s.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) FindUserByID(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
