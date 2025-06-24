package service

type Service interface {
	GetPosts() ([]GetPostsDTO, error)
	GetPostsByUsername(username string) ([]GetPostByUsernameDTO, error)
	CreatePost(createPostInput *CreatePostInput) error
	DeletePost(id string) error
	UpdatePost(id string) error
}
