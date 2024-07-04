package database

import (
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user repository define methods to access user data operations
type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	AllUsers() ([]*models.User, error)
	GetUserByID(userID primitive.ObjectID) (*models.User, error)
}
