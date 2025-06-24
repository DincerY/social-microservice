package service

import (
	"errors"
)

//Zorunlu alan kontrolü handler'da olabilir ama "bu kullanıcı zaten bir post paylaşmış mı?" gibi bir kuralsal kontrol burada yapılmalı.

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrEmailTaken      = errors.New("email already exists")
	ErrUsernameTaken   = errors.New("username already exists")
)

type Service interface {
	//model.User yerine servisin kendi dto ları olabilir

	GetUsers() ([]GetUsersDTO, error)
	GetUserByUsername(username string) (*GetUserByUsernameDTO, error)
	GetUserProfile(userID string) (*GetUserProfileDTO, error)

	CreateUser(createUserInput *CreateUserInput) (*CreateUserDTO, error)
	DeleteUser(id string) error
	UpdateUser(id string) error

	UpdateUserTitle(id, newTitle string) error
	ExistsByUsername(username string) (bool, error)
	ChangePassword(userID string, oldPassword, newPassword string) error
	DeactivateUser(userID string) error
	IsUsernameAvailable(username string) (bool, error)
}
