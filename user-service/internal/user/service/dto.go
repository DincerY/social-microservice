package service

import "time"

type CreateUserInput struct {
	Username     string
	Email        string
	Bio          string
	ProfileImage string
}

type CreateUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type GetUsersDTO struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Bio          string    `json:"bio"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetUserByUsernameDTO struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Bio          string `json:"bio"`
	ProfileImage string `json:"profile_image"`
}

type GetUserProfileDTO struct {
	Bio          string `json:"bio"`
	ProfileImage string `json:"profile_image"`
}
