package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DincerY/social-microservice/user-service/internal/user/handler"
	"github.com/DincerY/social-microservice/user-service/internal/user/repository"
	"github.com/DincerY/social-microservice/user-service/internal/user/service"

	"github.com/DincerY/social-microservice/user-service/internal/user"
	"github.com/DincerY/social-microservice/user-service/pkg/config"
	"github.com/DincerY/social-microservice/user-service/pkg/database"
	_ "github.com/DincerY/social-microservice/user-service/pkg/log"
	"github.com/DincerY/social-microservice/user-service/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type ValidationErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func main() {
	appConfig := config.Read()

	defer zap.L().Sync()

	db, err := database.NewPostgresDb(appConfig.DbConnectionString)
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	//testRepository := repository.NewTestRepository(db)
	userService := service.NewCreateUserHandler(userRepository)

	userHandler := handler.NewUserHandler(userService, user.NewValidator())

	app := fiber.New()

	//log header
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)
		ip := c.IP()
		path := c.Request().URI()

		zap.L().Info(fmt.Sprintf("[REQUEST] %s | IP: %s | Duration: %v\n", path, ip, duration))
		return err
	})

	internalGroup := app.Group("/internal")
	internalGroup.Post("/user", userHandler.CreateUser)

	apiGroup := app.Group("/api")

	apiGroup.Get("/healtcheck", func(c *fiber.Ctx) error {
		zap.L().Info("Server is working")
		return c.Status(fiber.StatusOK).SendString("Server is working")
	})

	apiGroup.Get("/user", middleware.JWTMiddleware(appConfig), userHandler.GetUserByUsername)

	adminGroup := app.Group("/admin", middleware.JWTMiddleware(appConfig), middleware.AdminMiddleware("role", "admin"))

	adminGroup.Get("/users", userHandler.GetUsers)

	go app.Listen(":" + appConfig.Port)
	gracefullShutdown(app)
}

func gracefullShutdown(app *fiber.App) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	zap.L().Info("Shutting down server...")
	if err := app.ShutdownWithTimeout(5 * time.Second); err != nil {
		zap.L().Info("Error during server shutdown")
	}
	zap.L().Info("Server gracefully stopped!")
}
