package database

import (
	"context"
	"log/slog"
	"os"

	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func CreateUser(db *mongo.Database, user *models.User) (*models.User, error) {
	slog.SetDefault(logger)

	//new user ID init
	user.ID = primitive.NewObjectID()
	result, err := db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		slog.Error(err.Error())

		return nil, err
	}
	slog.Info("new user created", "id", result.InsertedID)
	return user, nil
}
