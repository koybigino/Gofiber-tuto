package validation

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

var validate = validator.New()

func ValidateStruct(s interface{}) []*ErrorResponse {
	var errors []*ErrorResponse

	errs := validate.Struct(s)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var e ErrorResponse
			e.FailedField = err.StructField()
			e.Tag = err.Tag()
			e.Value = err.Param()
			errors = append(errors, &e)
		}
	}

	return errors
}
