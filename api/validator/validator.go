package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	customValidator *validator.Validate
	usernameRegex   = regexp.MustCompile("^[A-Za-z0-9][A-Za-z0-9_-]*$")
)

func InitValidator() {
	customValidator = validator.New()
	customValidator.RegisterValidation("username", usernameValidator)
}

func Struct(s interface{}) error {
	return customValidator.Struct(s)
}

func usernameValidator(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
}
