package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Importa el driver de archivo
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
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
	// MongoDB Connection
	client := db.DBconnect()
	if !db.CheckConnection() {
		slog.Error("DB is not running")
		return
	}
	// migrations
	mongoURI := os.Getenv("MONGODB_URI")     // Ensure it includes the dbname if not already part of URI
	migrationsPath := "file://db/migrations" // Path to migrations directory
	//routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to a very over engineering API TODO App")
	})
	migrateDatabase(client, mongoURI, migrationsPath)
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

// migrateDatabase handles the migration process
func migrateDatabase(db *mongo.Client, mongoURI, migrationsPath string) {
	// Use "mongodb" as the prefix for the MongoDB driver, e.g., "mongodb://localhost:27017/dbname"
	mongoDriver, err := mongodb.WithInstance(db, &mongodb.Config{})
	if err != nil {
		log.Fatalf("Failed to create MongoDB driver instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, os.Getenv("DB_NAME"), mongoDriver)
	if err != nil {
		log.Fatalf("Failed to initialize migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	fmt.Println("Migrations applied successfully")
}
