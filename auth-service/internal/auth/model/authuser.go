package model

import "time"

type AuthUser struct {
	ID           string    `json:"id"`
	PasswordHash string    `json:"password_hash"`
	Username     string    `json:"username"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
