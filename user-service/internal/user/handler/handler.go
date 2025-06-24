package handler

import (
	"fmt"
	"strings"

	"github.com/DincerY/social-microservice/user-service/internal/user"
	"github.com/DincerY/social-microservice/user-service/internal/user/response"
	"github.com/DincerY/social-microservice/user-service/internal/user/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type UserHandler struct {
	service   service.Service
	validator *user.UserValidator
}

func NewUserHandler(service service.Service, validator *user.UserValidator) *UserHandler {
	return &UserHandler{
		service:   service,
		validator: validator,
	}
}

// Bu fonksiyon başka servisler tarafından kullanıcak harici bir client tarafından değil
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var createUserRequest CreateUserRequest
	if err := c.BodyParser(&createUserRequest); err != nil {
		zap.L().Info(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ServiceResponse[any]{
			Success: false,
			Error: &response.ServiceError{
				Code:    "INTERNAL_ERROR",
				Message: "body was not parsed",
			},
		})
	}

	if errs := h.validator.Validate(createUserRequest); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)
		failedField := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
			failedField = append(failedField, err.FailedField)
		}
		errMessage := strings.Join(errMsgs, " and ")
		zap.L().Info(errMessage)
		return c.Status(fiber.StatusBadRequest).JSON(response.ServiceResponse[any]{
			Success: false,
			Error: &response.ServiceError{
				Code:    "VALIDATION_ERROR",
				Message: errMessage,
				Field:   strings.Join(failedField, " , "),
			},
		})
	}

	createUserInput := service.CreateUserInput{
		Username:     createUserRequest.Username,
		Email:        createUserRequest.Email,
		Bio:          createUserRequest.Bio,
		ProfileImage: createUserRequest.ProfileImage,
	}
	createUserDTO, err := h.service.CreateUser(&createUserInput)
	if err != nil {
		zap.L().Info(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ServiceResponse[any]{
			Success: false,
			Error: &response.ServiceError{
				Code:    "INTERNAL_ERROR",
				Message: "user did not created",
			},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.ServiceResponse[service.CreateUserDTO]{
		Success: true,
		Data:    createUserDTO,
	})
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	getUsersDto, err := h.service.GetUsers()
	if err != nil {
		zap.L().Info(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "an unexpected error occurred",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.APIResponse[[]service.GetUsersDTO]{
		Success: true,
		Data:    &getUsersDto,
	})
}

func (h *UserHandler) GetUserByUsername(c *fiber.Ctx) error {
	tokenUser := c.Locals("user").(*jwt.Token)
	claims := tokenUser.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

	getUserByUsernameDto, err := h.service.GetUserByUsername(userName)
	if err != nil {
		zap.L().Info(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ServiceResponse[any]{
			Success: false,
			Error: &response.ServiceError{
				Code:    "INTERNAL_ERROR",
				Message: "an unexpected error occurred",
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.ServiceResponse[service.GetUserByUsernameDTO]{
		Success: true,
		Data:    getUserByUsernameDto,
	})
}
