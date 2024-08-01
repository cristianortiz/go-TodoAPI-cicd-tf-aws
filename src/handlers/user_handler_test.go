package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/handlers"
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/middleware"
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// userService mock
type MockUserService struct {
	mock.Mock
}

// mocking createUser service method
func (m *MockUserService) CreateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(userID primitive.ObjectID) (*models.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*models.User), args.Error(1)
}
func (m *MockUserService) AllUsers() ([]*models.User, error) {
	args := m.Called()
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *MockUserService) FindUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

// createUserHandler test
func TestCreateUserHandler(t *testing.T) {
	//new app fiber
	app := fiber.New()
	//init middleware
	models.InitValidator()

	//create mock user service
	mockUserService := new(MockUserService)

	user := &models.User{
		Name:         "Test User",
		Email:        "test@example.com",
		PasswordHash: "password123",
		ID:           primitive.NewObjectID(),
	}
	//config mock to return test user data
	mockUserService.On("CreateUser", mock.AnythingOfType("*models.User")).Return(user, nil)

	//register handler with user service mock and middleware
	app.Post("/v1/user", middleware.ValidationMiddleware(&models.User{}), handlers.CreateUserHandler(mockUserService))

	// parse user to JSON
	userJSON, _ := json.Marshal(user)

	//create a test POST request
	req := httptest.NewRequest(http.MethodPost, "/v1/user", bytes.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	//register test request and get response
	resp, _ := app.Test(req, -1)
	//verify response status
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	//verify response content
	var createdUser models.User
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&createdUser)
	assert.Equal(t, user.Name, createdUser.Name)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.NotEmpty(t, createdUser.ID)

}
