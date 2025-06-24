package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/DincerY/social-microservice/post-service/internal/post/model"
	"github.com/DincerY/social-microservice/post-service/internal/post/repository"
	"github.com/google/uuid"
)

type PostService struct {
	repository repository.Repository
}

func NewPostService(repository repository.Repository) *PostService {
	return &PostService{repository: repository}
}

func (s *PostService) GetPosts() ([]GetPostsDTO, error) {
	posts, err := s.repository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("PostService.GetPosts failed: %w", err)
	}
	var result []GetPostsDTO
	for _, post := range posts {
		temp := GetPostsDTO{
			ID:        post.ID,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
		}
		result = append(result, temp)
	}
	return result, nil
}

func (s *PostService) GetPostsByUsername(username string) ([]GetPostByUsernameDTO, error) {
	posts, err := s.repository.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("PostService.GetPostsByUsername failed: %w", err)
	}
	result := make([]GetPostByUsernameDTO, 0, len(posts))
	for _, post := range posts {
		result = append(result, GetPostByUsernameDTO{
			ID:        post.ID,
			Content:   post.Content,
			MediaURL:  post.MediaURL,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}
	return result, nil
}

func (s *PostService) CreatePost(createPostInput *CreatePostInput) error {
	post := model.Post{
		ID:        uuid.New().String(),
		Username:  createPostInput.Username,
		Content:   createPostInput.Content,
		MediaURL:  createPostInput.MediaURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	}
	err := s.repository.Create(post)
	if err != nil {
		return fmt.Errorf("PostService.CreatePost failed: %w", err)
	}
	return nil
}

func (s *PostService) DeletePost(id string) error {
	return errors.New("DeletePost")

}
func (s *PostService) UpdatePost(id string) error {
	return errors.New("UpdatePost")

}
