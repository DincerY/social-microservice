package service

import "time"

type SomethingInput struct {
}

type SomethingResult struct {
}

type GetPostsInput struct {
}

type GetPostsDTO struct {
	ID        string
	Content   string
	CreatedAt time.Time
}

type GetPostByUsernameDTO struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	MediaURL  *string   `json:"media_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreatePostInput struct {
	Username string  `json:"username"`
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	MediaURL *string `json:"media_url"`
}
