package handler

import (
	"time"

	"github.com/DincerY/social-microservice/user-service/internal/user/model"
)

type CreateUserRequest struct {
	ID           string `json:"id"`
	Username     string `json:"username" validate:"required,min=5,max=20"`
	Email        string `json:"email" validate:"required,email"`
	Bio          string `json:"bio" validate:"required,min=5,max=100"`
	ProfileImage string `json:"profile_image"`
}

type GetUserByUsernameRequest struct {
	Username string `json:"username"`
}

type GetUserByUsernameResponse struct {
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Bio          string    `json:"bio"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetUsersRequest struct {
}

type GetUsersResponse struct {
	Users []model.User
}
