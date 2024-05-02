package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// init() is called before main, ideal to load env vars before anything else
func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file, using default values ", err)
	}
}

func main() {
	//fiber app init
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to very over engineering API TODO App")
	})
	port := os.Getenv("SERVER_PORT")

	go func() {
		slog.Info("Server running ", "port", port)

		if err := app.Listen(":" + port); err != nil && err != http.ErrServerClosed {
			slog.Error("Error starting server: ", err)

		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		slog.Error("Server shutdown failed: ", err)
	}
	slog.Info("Server shutdown gracefully")

}
