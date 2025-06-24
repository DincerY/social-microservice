package service

type CreateUserPayload struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Bio          string `json:"bio"`
	ProfileImage string `json:"profile_image"`
}

type LoginInput struct {
	Username string
	Password string
}

type TokenResult struct {
	AccessToken string
	ExpiresIn   int64
}

type RegisterInput struct {
	Username     string
	Email        string
	Password     string
	Bio          string
	ProfileImage string
}

type RegisterResult struct {
	Username string
}
