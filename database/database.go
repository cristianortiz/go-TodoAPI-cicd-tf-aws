package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConn, _ = DBconnect()

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
	opts := os.Getenv("MONGODB_OPTIONS")

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
func CheckConnection() bool {
	//check if the db is running, in a new err variable
	err := MongoConn.Ping(context.TODO(), nil)

	return err == nil

}
