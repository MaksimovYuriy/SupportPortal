package services

import (
	"context"

	"github.com/MaksimovYuriy/SupportPortal/internal/database"
	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type UserService struct {
	userRepo     database.UserRepository
	agentService *AgentService
}

func NewUserService(userRepo database.UserRepository, agentService *AgentService) *UserService {
	return &UserService{userRepo: userRepo, agentService: agentService}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	err = s.agentService.CreateAgentForUser(ctx, user)
	if err != nil {
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
