package service

import (
	"fmt"
	"time"

	"github.com/DincerY/social-microservice/user-service/internal/user/model"
	"github.com/DincerY/social-microservice/user-service/internal/user/repository"
	"github.com/google/uuid"
)

type UserService struct {
	repository repository.Respository
}

func NewCreateUserHandler(repository repository.Respository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetUsers() ([]GetUsersDTO, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		//bu şekilde yaparak dönülen hatayı logladığımızda hangi fonksiyondan döndüğünü görebileceğiz
		//bu şekilde bir çalışma yapmak hatayı anlamalandırmak veya yorumlamak değildir
		return nil, fmt.Errorf("UserService.GetUsers failed: %w", err)
	}

	//Eğer ki use case için bir iş mantığı hatası ise
	if len(users) == 0 {
		return nil, ErrUserNotFound
	}

	getUsersDTO := make([]GetUsersDTO, 0, len(users))
	for _, user := range users {
		getUsersDTO = append(getUsersDTO, GetUsersDTO{
			ID:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			Bio:          user.Bio,
			ProfileImage: user.ProfileImage,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		})
	}
	return getUsersDTO, nil
}

// Bu fonksiyon, "bu kullanıcı kesinlikle olmalı" yaklaşımıyla yazıldı
func (s *UserService) GetUserByUsername(username string) (*GetUserByUsernameDTO, error) {
	user, err := s.repository.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("UserService.GetUserByUsername failed: %w", err)
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return &GetUserByUsernameDTO{
		Username:     user.Username,
		Email:        user.Email,
		Bio:          user.Bio,
		ProfileImage: user.ProfileImage,
	}, nil
}

func (s *UserService) GetUserProfile(userID string) (*GetUserProfileDTO, error) {
	return nil, nil

}

func (s *UserService) CreateUser(createUserInput *CreateUserInput) (*CreateUserDTO, error) {
	exist, err := s.repository.ExistsByUsername(createUserInput.Username)
	if err != nil {
		return nil, fmt.Errorf("UserService.CreateUser failed: %w", err)
	}
	if exist {
		return nil, ErrUsernameTaken
	}
	user := model.User{
		ID:           uuid.New().String(),
		Username:     createUserInput.Username,
		Email:        createUserInput.Email,
		Bio:          createUserInput.Bio,
		ProfileImage: createUserInput.ProfileImage,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = s.repository.Create(&user)
	if err != nil {
		return nil, fmt.Errorf("UserService.CreateUser failed: %w", err)
	}
	return &CreateUserDTO{
		Username: user.Username,
		Email:    user.Email,
		Bio:      user.Bio,
	}, nil
}

func (s *UserService) DeleteUser(id string) error {
	return nil
}

func (s *UserService) UpdateUser(id string) error {
	return nil
}

func (s *UserService) UpdateUserTitle(id, newTitle string) error {
	return nil
}

func (s *UserService) ExistsByUsername(username string) (bool, error) {
	user, err := s.repository.GetByUsername(username)
	if err != nil {
		return false, fmt.Errorf("UserService.ExistsByUsername failed: %w", err)
	}
	if user == nil {
		return false, nil
	}
	return true, nil
}

func (s *UserService) ChangePassword(userID string, oldPassword, newPassword string) error {
	return nil
}

func (s *UserService) DeactivateUser(userID string) error {
	return nil
}

func (s *UserService) IsUsernameAvailable(username string) (bool, error) {
	return true, nil
}
