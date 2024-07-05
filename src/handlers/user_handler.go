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

func CreateUserHandler(ur database.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slog.SetDefault(logger)

		// Obtener el modelo validado del contexto
		validatedUser := c.Locals("validatedModel")
		if validatedUser == nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Validated model is missing"})
		}

		// Convertir el modelo validado a *models.User
		user, ok := validatedUser.(*models.User)
		if !ok {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid model type"})
		}
		//log the parsed and already validated body data
		slog.Info("createUser Request", "user:", user)
		// password encryption
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "error encrypting password"})
		}
		user.PasswordHash = string(hashedPassword)

		createdUser, err := ur.CreateUser(user)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "error creating user"})
		}
		return c.Status(http.StatusCreated).JSON(createdUser)

	}

}
