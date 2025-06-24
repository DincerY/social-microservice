package repository

import (
	"database/sql"
	"errors"

	"github.com/DincerY/social-microservice/user-service/internal/user/model"
)

type TestRepository struct {
	db *sql.DB
}

func NewTestRepository(db *sql.DB) *TestRepository {
	return &TestRepository{db: db}
}

func (r *TestRepository) GetAll() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, email, username, bio, profile_image, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Bio, &user.ProfileImage, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *TestRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User

	query := `SELECT id, email, username, bio, profile_image, created_at, updated_at FROM users WHERE username = $1`

	row := r.db.QueryRow(query, username)

	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Bio, &user.ProfileImage, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (r *TestRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User

	query := `SELECT id, email, username, bio, profile_image, created_at, updated_at FROM users WHERE email = $1`

	row := r.db.QueryRow(query, email)

	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Bio, &user.ProfileImage, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (r *TestRepository) Create(user *model.User) error {
	query := `INSERT INTO users (id, username, email, bio, profile_image, created_at, updated_at)
		          VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, user.ID, user.Username, user.Email, user.Bio, user.ProfileImage, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *TestRepository) Update(user *model.User) error {
	return errors.New("Update")
}

func (r *TestRepository) Delete(id string) error {
	return errors.New("Delete")
}

func (r *TestRepository) ExistsByUsername(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1) AS exists"
	err := r.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
