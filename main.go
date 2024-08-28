package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/database"
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/handlers"
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/middleware"
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/services"
	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Importa el driver de archivo
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading .env file, using default values ", "error", err.Error())
	}
	//fiber app init
	app := fiber.New()
	//logger config
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	// MongoDB Connection
	db := database.DBconnect()
	defer database.Disconnect()
	//initialize models struct fields validator,  applied via middleware
	models.InitValidator()
	//userRepository instance and injected with db MongoDB Database
	userRepo := database.NewMongoUserRepository(db)
	//create userService instance to be injected with userRepo struct
	userService := services.NewUserService(userRepo)

	//---------------routes-----------------------------------
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to a very over engineering API TODO App")
	})
	//grouping user endpoints
	user := app.Group("/v1")

	//create user EP with struct fields validations, and injected with userService
	user.Post("/user", middleware.ValidationMiddleware(&models.User{}), handlers.CreateUserHandler(userService))

	//----------------------------------------------------------------------
	port := os.Getenv("SERVER_PORT")

	go func() {
		slog.Info("Server running ", "port", port)

		if err := app.Listen(":" + port); err != nil && err != http.ErrServerClosed {
			slog.Error("error", "Error starting server: ", err.Error())

		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		slog.Error("error", "Server shutdown failed: ", err.Error())
	}
	slog.Info("Server shutdown gracefully")

}
