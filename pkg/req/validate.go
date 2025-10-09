package req

import (
	"github.com/go-playground/validator/v10"
)

func Validate[T any](paylaod T) error {
	validate := validator.New()
	err := validate.Struct(paylaod)
	return err
}
