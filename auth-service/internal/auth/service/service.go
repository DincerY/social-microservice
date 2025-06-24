package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DincerY/social-microservice/auth-service/internal/auth/model"
	"github.com/DincerY/social-microservice/auth-service/internal/auth/repository"
	"github.com/DincerY/social-microservice/auth-service/internal/security"
	"github.com/google/uuid"
)

type AuthService struct {
	repository repository.Respository
}

func NewAuthService(repository repository.Respository) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Register(registerInput *RegisterInput) error {
	registeredUser, err := s.repository.ExistsByUsername(registerInput.Username)

	if err != nil {
		return fmt.Errorf("AuthService.Register failed: %w", err)
	}
	if registeredUser {
		return ErrUsernameTaken
	}

	if len(registerInput.Password) < 6 {
		return ErrPasswordTooShort // iş kuralı hatası
	}

	hash, err := security.HashPassword(registerInput.Password)
	if err != nil {
		return fmt.Errorf("AuthService.Register failed: %w", err)
	}

	var user CreateUserPayload
	user.ID = uuid.New().String()
	user.Bio = registerInput.Bio
	user.Email = registerInput.Email
	user.ProfileImage = registerInput.ProfileImage
	user.Username = registerInput.Username

	payloadBytes, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("AuthService.Register failed: %w", err)
	}

	userServiceURL := "http://localhost:3000/internal/user"
	resp, err := http.Post(userServiceURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return ErrUserServiceUnavailable
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return ErrUserServiceError
	}
	authUser := model.AuthUser{
		ID:           user.ID,
		Username:     registerInput.Username,
		Role:         "user",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		PasswordHash: hash,
	}
	err = s.repository.Create(&authUser)
	if err != nil {
		return fmt.Errorf("AuthService.Register failed: %w", err)
	}
	return nil
}

func (s *AuthService) Login(loginInput *LoginInput) (*TokenResult, error) {

	authUser, err := s.repository.GetByUsername(loginInput.Username)
	if err != nil {
		return nil, fmt.Errorf("AuthService.Login failed: %w", err)
	}

	if authUser == nil {
		return nil, ErrInvalidCredentials
	}

	if !security.CheckPasswordHash(loginInput.Password, authUser.PasswordHash) {
		return nil, ErrInvalidCredentials
	}
	token, err := security.CreateJWT(authUser)
	if err != nil {
		return nil, fmt.Errorf("AuthService.Login failed: %w", err)
	}
	return &TokenResult{
		AccessToken: token,
		ExpiresIn:   1,
	}, nil
}
