package middleware

import (
	"fmt"

	"github.com/DincerY/social-microservice/user-service/internal/user/response"
	"github.com/DincerY/social-microservice/user-service/pkg/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(appConfig *config.AppConfig) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(appConfig.JwtSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
				Success: false,
				Error:   fmt.Sprintf("invalid or expired token %v", err.Error()),
			})

		},
	})
}
