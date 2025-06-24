package handler

import "time"

type CreatePostRequest struct {
	Content  string  `json:"content" db:"content"`
	MediaURL *string `json:"media_url,omitempty" db:"media_url"`
}

type CreatePostResponse struct {
}

type GetPostsByUsernameRequest struct {
	Content  string  `json:"content" db:"content"`
	MediaURL *string `json:"media_url,omitempty" db:"media_url"`
}

type GetPostsByUsernameResponse struct {
}

type GetPostsRequest struct {
}
type GetPostsResponse struct {
	ID        string    `json:"id" db:"id"`
	Username  string    `json:"user_id" db:"username"`
	Content   string    `json:"content" db:"content"`
	MediaURL  *string   `json:"media_url,omitempty" db:"media_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
