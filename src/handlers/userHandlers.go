package handlers

import (
	"log/slog"
	"net/http"
	"os"

	database "github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/database"
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func CreateUserHandler(c *fiber.Ctx) error {
	slog.SetDefault(logger)
	db := database.DBconnect()

	//get user struct model validates trough  ValidationMiddleware from fiber context
	user := c.Locals("validatedModel").(*models.User)

	//log the parsed and already validated body data
	slog.Info("createUser Request", "user:", user)
	// password encryption
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "error encrypting password"})
	}
	user.PasswordHash = string(hashedPassword)

	createdUser, err := database.CreateUser(db, user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "error creating user"})
	}
	return c.Status(http.StatusCreated).JSON(createdUser)
}
