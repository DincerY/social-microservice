package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/DincerY/social-microservice/auth-service/internal/auth/model"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetAll() ([]model.AuthUser, error) {

	query := `SELECT id, username, password_hash, created_at, updated_at, role FROM auth_users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("AuthRepository.GetAll failed: %w", err)
	}
	defer rows.Close()

	var authUsers []model.AuthUser

	for rows.Next() {
		var authUser model.AuthUser
		if err := rows.Scan(&authUser.ID, &authUser.Username, &authUser.PasswordHash, &authUser.CreatedAt, &authUser.UpdatedAt, &authUser.Role); err != nil {
			return nil, fmt.Errorf("AuthRepository.GetAll failed: %w", err)
		}
		authUsers = append(authUsers, authUser)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("AuthRepository.GetAll failed: %w", err)
	}
	return authUsers, nil

}

func (r *AuthRepository) GetByUsername(username string) (*model.AuthUser, error) {
	var authUser model.AuthUser

	query := `SELECT id, username, password_hash, created_at, updated_at, role FROM auth_users WHERE username = $1`

	row := r.db.QueryRow(query, username)

	if err := row.Scan(&authUser.ID, &authUser.Username, &authUser.PasswordHash, &authUser.CreatedAt, &authUser.UpdatedAt, &authUser.Role); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("AuthRepository.GetByUsername failed: %w", err)
	}
	return &authUser, nil
}

func (r *AuthRepository) Create(authUser *model.AuthUser) error {
	query := `INSERT INTO auth_users (id, username, password_hash, created_at, updated_at,role)
		          VALUES ($1, $2, $3, $4, $5,$6)`

	_, err := r.db.Exec(query, authUser.ID, authUser.Username, authUser.PasswordHash, authUser.CreatedAt, authUser.UpdatedAt, authUser.Role)
	if err != nil {
		return fmt.Errorf("AuthRepository.Create failed: %w", err)
	}
	return nil
}

func (r *AuthRepository) Update(user *model.AuthUser) error {
	return errors.New("Update")
}
func (r *AuthRepository) Delete(id string) error {
	return errors.New("Delete")
}
func (r *AuthRepository) SoftDelete(userID string) error {
	return errors.New("SoftDelete")
}

func (r *AuthRepository) ExistsByUsername(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM auth_users WHERE username = $1) AS exists"
	err := r.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("AuthRepository.ExistsByUsername failed: %w", err)
	}
	return exists, nil
}
