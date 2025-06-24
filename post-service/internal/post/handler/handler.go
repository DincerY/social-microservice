package handler

import (
	"github.com/DincerY/social-microservice/post-service/internal/post/response"
	"github.com/DincerY/social-microservice/post-service/internal/post/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type PostHandler struct {
	service service.Service
}

func NewPostHandler(service service.Service) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) GetPosts(c *fiber.Ctx) error {
	posts, err := h.service.GetPosts()
	if err != nil {
		zap.L().Info(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[[]service.GetPostsDTO]{
			Success: false,
			Message: "an unexpected error occurred",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.ServiceResponse[[]service.GetPostsDTO]{
		Success: true,
		Data:    &posts,
	})
}

func (h *PostHandler) GetPostByUsername(c *fiber.Ctx) error {
	tokenUser := c.Locals("user").(*jwt.Token)
	claims := tokenUser.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	posts, err := h.service.GetPostsByUsername(username)
	if err != nil {
		zap.L().Info(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Message: "an unexpected server error occurred",
		})
	}

	return c.Status(fiber.StatusOK).JSON(posts)
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	var createPostRequest CreatePostRequest
	if err := c.BodyParser(&createPostRequest); err != nil {
		zap.L().Info(err.Error())

		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Message: "body was not parsed",
		})

	}
	tokenUser := c.Locals("user").(*jwt.Token)
	claims := tokenUser.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	createPostInput := service.CreatePostInput{
		Username: username,
		Content:  createPostRequest.Content,
		MediaURL: createPostRequest.MediaURL,
	}

	if err := h.service.CreatePost(&createPostInput); err != nil {
		zap.L().Info(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Message: "an unexpected server error occurred",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(response.ServiceResponse[service.CreatePostInput]{
		Success: true,
		Data:    &createPostInput,
	})
}
