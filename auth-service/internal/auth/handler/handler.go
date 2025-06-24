package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/DincerY/social-microservice/auth-service/internal/auth"
	"github.com/DincerY/social-microservice/auth-service/internal/auth/response"
	"github.com/DincerY/social-microservice/auth-service/internal/auth/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthHandler struct {
	service   service.Service
	validator *auth.AuthValidator
}

func NewAuthHandler(service service.Service, validator *auth.AuthValidator) *AuthHandler {
	return &AuthHandler{
		service:   service,
		validator: validator,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var loginRequest LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		zap.L().Info(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "body was not parsed",
		})
	}
	loginInput := service.LoginInput{
		Username: loginRequest.Username,
		Password: loginRequest.Password,
	}
	tokenResult, err := h.service.Login(&loginInput)

	if err != nil {
		zap.L().Info(err.Error())
		if errors.Is(err, service.ErrInvalidCredentials) {
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
				Success: false,
				Error:   "invalid username or password",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "an unexpected error occurred",
		})
	}
	return c.Status(fiber.StatusOK).JSON(response.APIResponse[LoginResponse]{
		Success: false,
		Data: &LoginResponse{
			Token:       tokenResult.AccessToken,
			ExpiresDate: tokenResult.ExpiresIn,
		},
		Message: "login successful",
	})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var registerRequest RegisterRequest
	if err := c.BodyParser(&registerRequest); err != nil {
		zap.L().Info(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "body was not parsed",
		})
	}

	if errs := h.validator.Validate(registerRequest); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
		errMessage := strings.Join(errMsgs, " and ")
		zap.L().Info(errMessage)
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   errMessage,
		})
	}

	registerInput := service.RegisterInput{
		Username:     registerRequest.Username,
		Email:        registerRequest.Email,
		Password:     registerRequest.Password,
		Bio:          registerRequest.Bio,
		ProfileImage: registerRequest.ProfileImage,
	}

	if err := h.service.Register(&registerInput); err != nil {
		zap.L().Info(err.Error())
		if errors.Is(err, service.ErrUsernameTaken) {
			return c.Status(fiber.StatusConflict).JSON(response.APIResponse[any]{
				Success: false,
				Error:   "username already taken",
			})
		}

		if errors.Is(err, service.ErrPasswordTooShort) {
			return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
				Success: false,
				Error:   "password must be at least 6 characters long",
			})
		}

		if errors.Is(err, service.ErrUserServiceUnavailable) {
			return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
				Success: false,
				Error:   "registration was not successful beacuse user service is currently unavailable",
			})
		}

		if errors.Is(err, service.ErrUserServiceError) {
			return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
				Success: false,
				Error:   "registration was not successful beacuse user service has error(s)",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "an unexpected error occurred",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.APIResponse[RegisterResponse]{
		Success: true,
		Data: &RegisterResponse{
			Username: registerRequest.Username,
			Email:    registerRequest.Email,
		},
		Message: "user registered",
	})
}
