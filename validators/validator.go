package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateRequest(toBeValidated interface{}) error {
	v := validator.New()
	v.RegisterValidation("enum_validator", enumValidator)
	validationErr := v.Struct(toBeValidated)

	if validationErr != nil {
		return validationErr
	}
	return nil
}

func enumValidator(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	allowed := strings.Split(fl.Param(), " ")
	allowedMapper := make(map[string]bool, len(allowed))

	for _, param := range allowed {
		allowedMapper[param] = true
	}

	return val == "" || allowedMapper[val]
}
