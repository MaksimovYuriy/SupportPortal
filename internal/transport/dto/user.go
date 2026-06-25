package dto

import (
	"time"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
)

type UserResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserListResponse struct {
	Users []*UserResponse `json:"users"`
}

func NewUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func NewUserListResponse(users []*models.User) *UserListResponse {
	userResponses := make([]*UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = NewUserResponse(user)
	}
	return &UserListResponse{Users: userResponses}
}

type UserRequest struct {
	Email    string `json:"email"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}

func NewUserRequest(email, role string, isActive bool) *UserRequest {
	return &UserRequest{
		Email:    email,
		Role:     role,
		IsActive: isActive,
	}
}
