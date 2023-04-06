package helpers

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code    int         `json:"code"`
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// APIResponse is for generating template responses
func APIResponse(message string, code int, err bool, data interface{}) Response {

	return Response{
		Code:    code,
		Error:   err,
		Message: message,
		Data:    data,
	}
}

// FormatValidationError func which holds errors during user input validation
func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
