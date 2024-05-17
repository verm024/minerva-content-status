package validators

import (
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(toBeValidated interface{}) error {
	v := validator.New()
	validationErr := v.Struct(toBeValidated)

	if validationErr != nil {
		return validationErr
	}
	return nil
}
