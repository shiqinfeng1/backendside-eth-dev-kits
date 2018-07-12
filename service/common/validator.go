package common

import "github.com/go-playground/validator"

type SimpleValidator struct {
	Validator *validator.Validate
}

func (cv *SimpleValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
