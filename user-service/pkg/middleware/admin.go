package middleware

import (
	"github.com/DincerY/social-microservice/user-service/internal/user/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AdminMiddleware(claimKey string, claimVal string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenUser := c.Locals("user").(*jwt.Token)
		claims := tokenUser.Claims.(jwt.MapClaims)
		var role string
		if claims[claimKey] != nil {
			role = claims[claimKey].(string)
		}

		if role != claimVal || role == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
				Success: false,
				Error:   "only admin users can access",
			})
		}

		return c.Next()
	}
}
