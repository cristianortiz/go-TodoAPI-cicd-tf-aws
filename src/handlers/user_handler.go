package handlers

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/services"
	"github.com/gofiber/fiber/v2"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

// CreateUserHandler receive the http request, in this case after validateMiddleware
// interception to body fields validations, then acces service layer methods
func CreateUserHandler(us services.UserServiceIf) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slog.SetDefault(logger)

		// extract validated model struct from context after validateMiddleware check
		validatedUser := c.Locals("validatedModel")
		if validatedUser == nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Validated model is missing"})
		}

		// Convert validated model to  *models.User for this specific userHandler
		user, ok := validatedUser.(*models.User)
		if !ok {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid model type"})
		}
		//log the parsed and already validated body data
		slog.Info("createUser Request", "user:", user)

		//additional bussines logic and creteUser Op is delegated to userService
		createdUser, err := us.CreateUser(user)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"createUser error": err.Error()})
		}
		return c.Status(http.StatusCreated).JSON(createdUser)

	}

}
