package services

import (
	"errors"

	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// userService is the implementation of UserServiceIf interface,
// and uses the userRepositoryIf interface to acces and operates with the DB
type UserService struct {
	repo repository.UserRepositoryIf
}

// NewUserService creates a new instance of UserService
// must implements all methods defined in UserServiceIf
// and also  acces useRepositoryIf
func NewUserService(repo repository.UserRepositoryIf) *UserService {
	return &UserService{repo: repo}
}

// implementacion of UserServiceIf methods, using userRepositoryIf
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {

	existingUser, err := s.repo.FindUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with the same email already exists")
	}
	// password encryption
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(hashedPassword)
	//use userRepositoryIf to acces CreateUser implementation in database/user_repository
	//and perform the create user operation
	return s.repo.CreateUser(user)
}
func (s *UserService) AllUsers() ([]*models.User, error) {
	return s.repo.AllUsers()
}
func (s *UserService) GetUserByID(userID primitive.ObjectID) (*models.User, error) {
	return s.repo.GetUserByID(userID)
}
