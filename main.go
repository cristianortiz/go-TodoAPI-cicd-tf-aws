package main

import (
	"context"
	"log"
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
		log.Println("Error loading .env file, using default values.")
	}
}

func main() {
	//fiber app init
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to very over engineering API TODO App")
	})

	go func() {
		if err := app.Listen(":3000"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)

		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server shutdown gracefully")

}
