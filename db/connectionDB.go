package db

import (
	"context"
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConn = DBconnect()

func DBconnect() *mongo.Client {
	//make DB connection
	// MongoDB Connection Setup
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		slog.Error("Failed to connect to MongoDB: %v", err)
		//return the client object even is empty
		return client
	}
	//check if the db is running, err= assign value to a existing value
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		slog.Error(err.Error())
		//return the client object even is empty
		return client
	}
	slog.Info("DB connection is running..")
	//return a valid DB connection
	return client
}

// check the DB with a ping
func CheckConnection() bool {
	//check if the db is running, in a new err variable
	err := MongoConn.Ping(context.TODO(), nil)

	return err == nil

}
