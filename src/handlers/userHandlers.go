package handlers

import (
	"log/slog"
	"net/http"
	"os"

	database "github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/database"
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func CreateUserHandler(c *fiber.Ctx) error {
	slog.SetDefault(logger)
	validate := validator.New(validator.WithRequiredStructEnabled())
	db := database.DBconnect()

	var user models.User

	//parse the user data from request body into user struct
	if err := c.BodyParser(&user); err != nil {

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := validate.Struct(user)
	if err != nil {
		ve := fiber.Map{}

		for _, e := range err.(validator.ValidationErrors) {

			ve[e.Field()] = e.Error()
			slog.Error(e.Error())

		}

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status_code": http.StatusBadRequest, "validation_err": ve})

	}

	//log the parsed body data
	slog.Info("createUser Request", "body:", user)
	// password encryption
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error encrypting password"})
	}
	user.PasswordHash = string(hashedPassword)

	createdUser, err := database.CreateUser(db, &user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating user"})
	}
	return c.Status(http.StatusCreated).JSON(createdUser)
}
