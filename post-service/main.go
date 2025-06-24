package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DincerY/social-microservice/post-service/internal/post/handler"
	"github.com/DincerY/social-microservice/post-service/internal/post/repository"
	"github.com/DincerY/social-microservice/post-service/internal/post/service"
	"github.com/DincerY/social-microservice/post-service/pkg/config"
	"github.com/DincerY/social-microservice/post-service/pkg/database"
	_ "github.com/DincerY/social-microservice/post-service/pkg/log"
	"github.com/DincerY/social-microservice/post-service/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	appConfig := config.Read()

	db, err := database.NewPostgresDb(appConfig.DbConnectionString)
	if err != nil {
		panic(err)
	}
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	app := fiber.New()

	apiGroup := app.Group("api", middleware.JWTMiddleware(appConfig))
	apiGroup.Post("/post", postHandler.CreatePost)
	apiGroup.Get("/post", postHandler.GetPostByUsername)

	adminGroup := app.Group("/admin", middleware.JWTMiddleware(appConfig), middleware.AdminMiddleware("role", "admin"))

	adminGroup.Get("/posts", postHandler.GetPosts)

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
