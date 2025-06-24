package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/DincerY/social-microservice/auth-service/pkg/log"
	_ "github.com/lib/pq"

	"github.com/DincerY/social-microservice/auth-service/internal/auth"
	"github.com/DincerY/social-microservice/auth-service/internal/auth/handler"
	"github.com/DincerY/social-microservice/auth-service/internal/auth/repository"
	"github.com/DincerY/social-microservice/auth-service/internal/auth/service"
	"github.com/DincerY/social-microservice/auth-service/pkg/config"
	"github.com/DincerY/social-microservice/auth-service/pkg/database"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	appConfig := config.Read()

	defer zap.L().Sync()

	db, err := database.NewPostgresDb(appConfig.DbConnectionString)
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	defer db.Close()

	authRepository := repository.NewAuthRepository(db)

	authService := service.NewAuthService(authRepository)

	authHandler := handler.NewAuthHandler(authService, auth.NewValidator())

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

	app.Get("/healtcheck", func(c *fiber.Ctx) error {
		zap.L().Info("Server is working")
		return c.Status(fiber.StatusOK).SendString("Server is working")
	})

	authGroup := app.Group("/auth")
	authGroup.Post("/login", authHandler.Login)
	authGroup.Post("/register", authHandler.Register)

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
