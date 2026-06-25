package handlers

import (
	"net/http"

	"github.com/MaksimovYuriy/SupportPortal/internal/models"
	"github.com/MaksimovYuriy/SupportPortal/internal/services"
	"github.com/MaksimovYuriy/SupportPortal/internal/transport/dto"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.ListUsers(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewUserListResponse(users))
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request dto.UserRequest
	if err := decodeJSONBody(r, &request); err != nil {
		handleError(w, err)
		return
	}
	user := models.User{
		Email:        request.Email,
		PasswordHash: "###",
		Role:         request.Role,
		IsActive:     request.IsActive,
	}
	if err := h.userService.CreateUser(r.Context(), &user); err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, dto.NewUserResponse(&user))
}

func (h *UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r)
	if err != nil {
		handleError(w, err)
		return
	}
	user, err := h.userService.FindUserByID(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.NewUserResponse(user))
}
