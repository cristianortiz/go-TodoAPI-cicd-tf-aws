package services_test

import (
	"testing"

	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock de UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}
func (m *MockUserRepository) CreateUserMock(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(userID primitive.ObjectID) (*models.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*models.User), args.Error(1)
}
func (m *MockUserRepository) AllUsers() ([]*models.User, error) {
	args := m.Called()
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *MockUserRepository) FindUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, nil // Retorna un nil que es tipo *models.User
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func TestServiceCreateUser(t *testing.T) {
	//create new test user
	user := &models.User{
		Name:         "Test User",
		Email:        "test@example.com",
		PasswordHash: "password123",
		ID:           primitive.NewObjectID(),
	}
	//config userRepository mock
	mockRepo := new(MockUserRepository)
	//simulate that a user with that email is not exists
	mockRepo.On("FindUserByEmail", user.Email).Return(nil, nil)
	//simulate that was succesfully created
	mockRepo.On("CreateUser", mock.AnythingOfType("*models.User")).Return(user, nil)
	//create instance os userService with repository mock injected
	userService := services.NewUserService(mockRepo)
	//try to create new user
	createdUser, err := userService.CreateUser(user)
	//check that any error wa returned
	assert.NoError(t, err)
	//check if the new created user email, and ID are the same of user test
	assert.Equal(t, user.Email, createdUser.Email)
	assert.NotEmpty(t, createdUser.ID)

	//check that the mock was called as expected
	mockRepo.AssertExpectations(t)

}
