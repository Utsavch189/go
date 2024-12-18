package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}

func FormatValidationError(err error) map[string]string {
	errors := map[string]string{}
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()
		errors[field] = "Failed validation: " + tag
	}
	return errors
}

// Regular expression compiled once
var alphanumRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

// custom validator
func init() {
	// Register alphanumeric validation function
	validate.RegisterValidation("alphanum", func(fl validator.FieldLevel) bool {
		return alphanumRegex.MatchString(fl.Field().String())
	})
}

/*
required	Field is required.
min=X	Minimum length for a string or minimum value for a number.
max=X	Maximum length for a string or maximum value for a number.
gt=X	Greater than a specific value.
lt=X	Less than a specific value.
email	Validates email format.
url	Validates URL format.
len=X	Exact length of a string.
*/
