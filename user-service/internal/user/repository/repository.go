package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/DincerY/social-microservice/user-service/internal/user/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, email, username, bio, profile_image, created_at, updated_at FROM users")
	if err != nil {
		//bu şekilde yaparak dönülen hatayı logladığımızda hangi fonksiyondan döndüğünü görebileceğiz
		//bu şekilde bir çalışma yapmak hatayı anlamalandırmak veya yorumlamak değildir
		return nil, fmt.Errorf("UserRepository.GetAll failed: %w", err)
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Bio, &user.ProfileImage, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("UserRepository.GetAll failed: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("UserRepository.GetAll failed: %w", err)
	}
	return users, nil
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User

	query := `SELECT id, email, username, bio, profile_image, created_at, updated_at FROM users WHERE username = $1`

	row := r.db.QueryRow(query, username)

	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Bio, &user.ProfileImage, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("UserRepository.GetByUsername failed: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User

	query := `SELECT id, email, username, bio, profile_image, created_at, updated_at FROM users WHERE email = $1`

	row := r.db.QueryRow(query, email)

	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Bio, &user.ProfileImage, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("UserRepository.GetByEmail failed: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) Create(user *model.User) error {
	query := `INSERT INTO users (id, username, email, bio, profile_image, created_at, updated_at)
		          VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, user.ID, user.Username, user.Email, user.Bio, user.ProfileImage, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("UserRepository.Create failed: %w", err)
	}
	return nil
}

func (r *UserRepository) Update(user *model.User) error {
	return errors.New("Update")
}

func (r *UserRepository) Delete(id string) error {
	return errors.New("Delete")
}

func (r *UserRepository) SoftDelete(userID string) error {
	return errors.New("Delete")
}

func (r *UserRepository) ExistsByUsername(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1) AS exists"
	err := r.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("UserRepository.ExistsByUsername failed: %w", err)
	}
	return exists, nil
}
