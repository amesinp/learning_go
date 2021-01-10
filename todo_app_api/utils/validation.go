package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// SetupValidator configures the request validator
func SetupValidator() {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}

// ValidateDTO validates a data transfer object
func ValidateDTO(data interface{}) string {
	errors := validate.Struct(data)
	if errors != nil {
		err := errors.(validator.ValidationErrors)[0]

		switch err.Tag() {
		case "required":
			return fmt.Sprintf(`"%s" is required`, err.Field())
		default:
			return fmt.Sprintf("Invalid request body")
		}
	}

	return ""
}
