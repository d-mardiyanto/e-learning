package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct validates a struct based on its tags.
func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}