package security

import (
	"fmt"
	"time"

	"github.com/DincerY/social-microservice/auth-service/internal/auth/model"
	"github.com/DincerY/social-microservice/auth-service/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(authUser *model.AuthUser) (string, error) {
	claims := jwt.MapClaims{
		"username": authUser.Username,
		"role":     authUser.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.JwtSecret))
	if err != nil {
		return "", fmt.Errorf("token did not create : %v", err)
	}
	return tokenString, nil
}
