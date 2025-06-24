package repository

import "github.com/DincerY/social-microservice/user-service/internal/user/model"

type Respository interface {
	GetAll() ([]model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id string) error
	SoftDelete(userID string) error
	ExistsByUsername(username string) (bool, error)
}
