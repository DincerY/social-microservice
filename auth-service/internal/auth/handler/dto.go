package handler

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"` // Şifre hash'i JSON çıktısında görünmemeli
}

type LoginResponse struct {
	Token       string `json:"token"`
	ExpiresDate int64  `json:"expires_date"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type RegisterRequest struct {
	Username     string `json:"username" validate:"required,min=3,max=20"`
	Email        string `json:"email"`
	Password     string `json:"password" validate:"required,min=3,max=20"`
	Bio          string `json:"bio"`
	ProfileImage string `json:"profile_image,omitempty"`
}
type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
