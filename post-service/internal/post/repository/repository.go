package repository

import (
	"database/sql"
	"fmt"

	"github.com/DincerY/social-microservice/post-service/internal/post/model"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}

}
func (r *PostRepository) Create(post model.Post) error {
	query := `INSERT INTO posts (id, username, content, media_url, created_at, updated_at, is_deleted)
		          VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, post.ID, post.Username, post.Content, post.MediaURL, post.CreatedAt, post.UpdatedAt, post.IsDeleted)
	if err != nil {
		return fmt.Errorf("PostRepository.Create failed: %w", err)
	}
	return nil
}

func (r *PostRepository) GetAll() ([]model.Post, error) {
	query := `SELECT * FROM posts`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("PostRepository.GetAll failed: %w", err)
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.ID, &post.Username, &post.Content, &post.MediaURL, &post.CreatedAt, &post.UpdatedAt, &post.IsDeleted)
		if err != nil {
			return nil, fmt.Errorf("PostRepository.GetAll failed: %w", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("PostRepository.GetAll failed: %w", err)
	}
	return posts, nil
}

func (r *PostRepository) GetByUsername(username string) ([]model.Post, error) {
	query := `SELECT * FROM posts WHERE username = $1`
	rows, err := r.db.Query(query, username)
	if err != nil {
		return nil, fmt.Errorf("PostRepository.GetByUsername failed: %w", err)
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.ID, &post.Username, &post.Content, &post.MediaURL, &post.CreatedAt, &post.UpdatedAt, &post.IsDeleted)
		if err != nil {
			return nil, fmt.Errorf("PostRepository.GetByUsername failed: %w", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("PostRepository.GetByUsername failed: %w", err)
	}
	return posts, nil
}
