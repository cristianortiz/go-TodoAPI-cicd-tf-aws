package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance    *mongo.Client
	clientInstanceErr error
	mongoOnce         sync.Once
	mongoDatabase     *mongo.Database
)

// DBconnect returns single mongoBD connection instance
func DBconnect() *mongo.Database {
	mongoOnce.Do(func() {

		// Load environment variables
		host := os.Getenv("MONGO_HOST")
		port := os.Getenv("MONGO_PORT")
		username := os.Getenv("MONGO_USERNAME")
		dbName := os.Getenv("DB_NAME")
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
		clientInstance = client
		mongoDatabase = client.Database(dbName)
		slog.Info("DB connection is running..")
	})
	if clientInstanceErr != nil {
		slog.ErrorContext(context.Background(), "error mongoDB instanceClient", slog.Any("error", clientInstanceErr))

	}
	//return a valid DB connection
	return mongoDatabase
}

func Disconnect() {
	if clientInstance != nil {
		if err := clientInstance.Disconnect(context.TODO()); err != nil {
			slog.ErrorContext(context.Background(), "failed to Disconnect MongoDB", slog.Any("error", err))
			panic(err)
		}
		slog.Info("disconnected from MongoDB")

	}
}
