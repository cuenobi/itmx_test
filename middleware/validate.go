package middleware

import (
	"fmt"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func Validate(s interface{}) error {
	return validate.Struct(s)
}

func ErrorResponse(err error) map[string]interface{} {
	validationErrors := err.(validator.ValidationErrors)
	errorMessages := make(map[string]interface{})

	for _, fieldError := range validationErrors {
		fieldName := fieldError.Field()
		errorMessage := fmt.Sprintf("%s is %s", fieldName, fieldError.Tag())
		errorMessages[fieldName] = errorMessage
	}

	return errorMessages
}
