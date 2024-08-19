package util

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	goValidator "github.com/go-playground/validator/v10"
)

var validator *goValidator.Validate

func SetValidator() {
	validator = goValidator.New()
	registerValidatorTag()
	registerCustomValidator()
}

// validateStruct validates and return readable error if its not nil.
func ValidateStruct(s interface{}) (err error) {
	if err = validator.Struct(s); err != nil {
		validationErrors := goValidator.ValidationErrors{}
		if !errors.As(err, &validationErrors) {
			Log().Error(err.Error())
			return
		}

		e := validationErrors[0] // Handle only the first-failed validation.

		switch e.Tag() {
		case "required":
			err = fmt.Errorf("VALIDATION_ERR: %s is required.", e.Field())
		case "oneof":
			err = fmt.Errorf("VALIDATION_ERR: %s must be one of %s.", e.Field(), e.Param())
		case "lte":
			err = fmt.Errorf("VALIDATION_ERR: %s must be lower than or equal %s.", e.Field(), e.Param())
		case "gte":
			err = fmt.Errorf("VALIDATION_ERR: %s must be greater than or equal %s.", e.Field(), e.Param())
		case "min":
			err = fmt.Errorf("VALIDATION_ERR: %s must be at least %s.", e.Field(), e.Param())
		case "max":
			err = fmt.Errorf("VALIDATION_ERR: %s must be less than %s.", e.Field(), e.Param())
		case "len":
			err = fmt.Errorf("VALIDATION_ERR: %s must be %s.", e.Field(), e.Param())
		case "email":
			err = fmt.Errorf("VALIDATION_ERR: %s malformed email.", e.Field())
		case "password":
			err = fmt.Errorf("VALIDATION_ERR: %s is too weak.", e.Field())
		case "uuid":
			err = fmt.Errorf("VALIDATION_ERR: %s malformed uuid.", e.Field())
		}
	}

	return
}

func registerValidatorTag() {
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0] //nolint

		if name == "-" {
			return ""
		}

		return name
	})
}

func registerCustomValidator() {
	if err := validator.RegisterValidationCtx(
		"password",
		func(_ context.Context, fl goValidator.FieldLevel) bool {
			val := fl.Field().String()
			if len(val) < 8 { //nolint
				return false
			}

			var number, upper, special bool
			for _, r := range val {
				switch {
				case unicode.IsNumber(r):
					number = true
				case unicode.IsUpper(r):
					upper = true
				case unicode.IsPunct(r) || unicode.IsSymbol(r):
					special = true
				case unicode.IsLetter(r) || r == ' ':
				}
			}

			return number && upper && special
		}); err != nil {
		Log().Error(err.Error())
		return
	}
}
