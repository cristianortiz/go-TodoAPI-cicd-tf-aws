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

	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Importa el driver de archivo
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading .env file, using default values ", err)
	}
	//fiber app init
	app := fiber.New()
	//logger config
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	// MongoDB Connection
	client, err := DBconnect()
	if err != nil {
		return
	}
	defer client.Disconnect(context.TODO())

	if !CheckConnection(client) {
		slog.Error("DB is not running")
		return
	}
	// migrations
	migrationsPath := "file:///app/database/migrations"
	migrateDatabase(client, os.Getenv("DB_NAME"), migrationsPath) // Path to migrations directory
	//routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to a very over engineering API TODO App")
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

// migrateDatabase handles the migration process
func migrateDatabase(db *mongo.Client, dbName, migrationsPath string) {
	// Use "mongodb" as the prefix for the MongoDB driver, e.g., "mongodb://localhost:27017/dbname"
	mongoDriver, err := mongodb.WithInstance(db, &mongodb.Config{
		DatabaseName: dbName,
	})
	if err != nil {
		log.Fatalf("Failed to create MongoDB driver instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, dbName, mongoDriver)
	if err != nil {
		log.Fatalf("Failed to initialize migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	fmt.Println("Migrations applied successfully")
}

func DBconnect() (*mongo.Client, error) {

	//make DB connection
	// MongoDB Connection Setup
	// Load environment variables
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	username := os.Getenv("MONGO_USERNAME")
	slog.Info(username)

	password := os.Getenv("MONGO_PASSWORD")
	database := os.Getenv("DB_NAME")
	opts := os.Getenv("MONGO_OPTIONS")

	// Construct the MongoDB URI
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s%s", username, password, host, port, database, opts)
	slog.Info(mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		slog.ErrorContext(context.Background(), "Failed to connect to MongoDB", slog.Any("error", err))
		panic(err)
	}
	//check if the db is running, err= assign value to a existing value
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		slog.ErrorContext(context.Background(), "DB not avalaible", slog.Any("error", err))
		//return the client object even is empty
		panic(err)
	}
	slog.Info("DB connection is running..")
	//return a valid DB connection
	return client, nil
}

// check the DB with a ping
func CheckConnection(MongoConn *mongo.Client) bool {
	//check if the db is running, in a new err variable
	err := MongoConn.Ping(context.TODO(), nil)

	return err == nil

}
