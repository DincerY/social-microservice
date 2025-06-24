package service

import (
	"errors"
)

var (
	ErrPasswordTooShort       = errors.New("password too short")
	ErrUserNotFound           = errors.New("user not found")
	ErrInvalidPassword        = errors.New("invalid password")
	ErrUsernameTaken          = errors.New("username already exists")
	ErrInvalidCredentials     = errors.New("username or password wrong")
	ErrUserServiceUnavailable = errors.New("user service is unavailable")
	ErrUserServiceError       = errors.New("user service has error(s)")
)

type Service interface {
	Login(loginInput *LoginInput) (*TokenResult, error)
	Register(registerInput *RegisterInput) error
}
