package util

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type errorResponse struct {
	FailedField string
	Tag         string
	Type        string
}

func ValidateRequest(input any) []*errorResponse {
	var errors []*errorResponse
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element errorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Type = err.Kind().String()
			errors = append(errors, &element)
		}
	}
	return errors
}
