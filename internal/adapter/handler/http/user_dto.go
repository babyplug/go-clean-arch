package http

import "go-hexagonal-architecture/internal/core/domain"

type registerRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse domain.User

// updateUserRequest represents the request body for updating a user
type updateUserRequest struct {
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Email string `json:"email" binding:"required,email" example:"test@example.com"`
}

type userListResponse struct {
	Meta *meta           `json:"meta"`
	Data []*userResponse `json:"users"`
}
