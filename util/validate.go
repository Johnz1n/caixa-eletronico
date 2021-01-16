package util

import (
	"errors"

	"gopkg.in/go-playground/validator.v9"
)

// Validate Validate struct from request body.
func Validate(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(validMessage(err.Field(), err.Tag(), err.Param()))
		}
		return nil
	}
	return nil
}

func validMessage(field string, tag string, param string) string {
	switch tag {
	case "required":
		return "Field " + field + " is required"
	default:
		return "Field " + field + " failed validation" + tag + " " + param
	}
}
