package errors

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	TraceErr string `json:"-"`
}

func (err CustomError) Error() string {
	return err.Message
}

func (err CustomError) StatusCode() int {
	return err.Code
}

// NewCustomError return custom error
func NewCustomError(mess string, code int) CustomError {
	return CustomError{
		Code:    code,
		Message: mess,
	}
}

var (
	MethodNotAllowedError = func(methodName string) CustomError {
		return NewCustomError(
			fmt.Sprintf("Just accept %v method.", methodName),
			http.StatusMethodNotAllowed,
		)
	}
	RouteNotFoundError      = NewCustomError("Route not found", http.StatusNotFound)
	SomethingWentWrongError = func(err error) CustomError {
		return NewCustomError("Something went wrong", http.StatusInternalServerError)
	}
	RequiredError = func(field string) CustomError {
		return NewCustomError(
			fmt.Sprintf("%v is required", field),
			http.StatusBadRequest,
		)
	}
	InvalidTypeError = func(field string) CustomError {
		return NewCustomError(
			fmt.Sprintf("%v is invalid type", field),
			http.StatusBadRequest,
		)
	}
	NotFoundError = func(field string) CustomError {
		return NewCustomError(
			fmt.Sprintf("%v not found", field),
			http.StatusNotFound,
		)
	}
	BodyNotAllowedError   = NewCustomError("Body not allowed", http.StatusBadRequest)
	PermissionDeniedError = NewCustomError("Permission denied", http.StatusBadRequest)
)
