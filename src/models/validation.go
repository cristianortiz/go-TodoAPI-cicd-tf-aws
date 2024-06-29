package models

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

// ValidateStruc: uses validations v10 library for struct fields validations,
// take a interface as parameter, allowing any struct and validate their fields
// based on the rules defined in every struct tag like, 'required' or 'email' etc
func ValidateStruct(data interface{}) (fiber.Map, error) {
	err := validate.Struct(data)
	if err != nil {
		if vallidationsErrors, ok := err.(validator.ValidationErrors); ok {
			//map for return founded errors in http response
			valErrors := fiber.Map{}
			for _, e := range vallidationsErrors {
				//field and error message added to error map
				valErrors[e.Field()] = e.Error()
				//log the valdiations error by field and and error
				slog.Error(e.Error())

			}
			return valErrors, err
		}
		return nil, err
	}
	return nil, nil
}
