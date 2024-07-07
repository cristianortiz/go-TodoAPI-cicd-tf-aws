package database

import (
	"context"
	"log/slog"
	"os"

	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

// this struct implements UserRepositoryIf interface using mongoDB
type MongoUserRepository struct {
	db *mongo.Database
}

// NewMongoUserRepository func creates a new instance of the above struct
// to implement the methods defined in repository/UserRepositoryIf
func NewMongoUserRepository(db *mongo.Database) repository.UserRepositoryIf {
	return &MongoUserRepository{db: db}
}

func (r *MongoUserRepository) CreateUser(user *models.User) (*models.User, error) {
	slog.SetDefault(logger)

	//new user ID init
	user.ID = primitive.NewObjectID()
	result, err := r.db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	slog.Info("new user created", "id", result.InsertedID)
	return user, nil
}
func (r *MongoUserRepository) AllUsers() ([]*models.User, error) {
	var users []*models.User
	return users, nil

}
func (r *MongoUserRepository) GetUserByID(userID primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.db.Collection("users").FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	filter := bson.M{"email": email}
	err := r.db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	return &user, nil
}
