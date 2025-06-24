package user

import "github.com/go-playground/validator"

type UserValidator struct {
	Validator *validator.Validate
}

func NewValidator() *UserValidator {
	return &UserValidator{
		Validator: validator.New(),
	}
}

type ValidationErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func (v UserValidator) Validate(data interface{}) []ValidationErrorResponse {
	validationErrors := []ValidationErrorResponse{}
	errs := v.Validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ValidationErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}
