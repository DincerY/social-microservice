package repository

import "github.com/DincerY/social-microservice/post-service/internal/post/model"

type Repository interface {
	Create(post model.Post) error
	GetAll() ([]model.Post, error)
	GetByUsername(username string) ([]model.Post, error)
}
