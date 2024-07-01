package middleware

import (
	"net/http"

	"github.com/cristianortiz/go-TodoAPI-cicd-tf-aws/src/models"
	"github.com/gofiber/fiber/v2"
)

// ValidationMiddleware uses model/validation for request struct fields validations
// intercepting the htttp request before  it reaches the handlers and applyng the mentioned
// validations
func ValidationMiddleware[T any](model T) fiber.Handler {
	//fibers expect that a Middleware returns a fiber.Handler type wich is a function whit this
	//signature func(*Ctx)error, this handler is chained in HTTP request cycle, with access to context
	//trough fiber.Context and in this case, performing the struct model validations specific action,
	return func(c *fiber.Ctx) error {
		// parse request body inputs into struct model
		if err := c.BodyParser(model); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}
		//validate struct model using ValidatiosnStruct
		validationErrors, err := models.ValidateStruct(model)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_errors": validationErrors})
		}

		//if validations passes, the validated struct model  is stored in fiber context
		//for make it available for being used in the next handler or middleware
		c.Locals("validateModel", model)
		//transfering control to the next handler or middleware on chain
		return c.Next()
	}
}
