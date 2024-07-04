package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// createUserHandler test
func TestCreateUserHandler(t *testing.T) {
	//new app fiber
	app := fiber.New()
	//grouping user endpoints

	userEP := app.Group("/v1")

	//create user EP with struc
	//init route to test handler
	userEP.Post("/user", CreateUserHandler)
	//create new test user
	user := models.User{
		Name:         "Test User",
		Email:        "test@example.com",
		PasswordHash: "password123",
	}
	// parse uset to JSON
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
