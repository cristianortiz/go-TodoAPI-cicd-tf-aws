package repository

import (
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user repository interface define methods for user CRUD operations
// this CRUD ops are implemented in pkg database/user_repository.go
type UserRepositoryIf interface {
	CreateUser(user *models.User) (*models.User, error)
	AllUsers() ([]*models.User, error)
	GetUserByID(userID primitive.ObjectID) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
}
