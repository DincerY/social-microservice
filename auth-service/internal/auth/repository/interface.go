package repository

import (
	"github.com/DincerY/social-microservice/auth-service/internal/auth/model"
)

type Respository interface {
	GetAll() ([]model.AuthUser, error)
	GetByUsername(username string) (*model.AuthUser, error)
	Create(user *model.AuthUser) error
	Update(user *model.AuthUser) error
	Delete(id string) error
	SoftDelete(userID string) error
	ExistsByUsername(username string) (bool, error)
}
