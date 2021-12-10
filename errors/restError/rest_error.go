// Package restError provides a simple REST error implementation.
package restError

import "errors"

// RestErr is a collection of errors
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewErrror(msg string) error {
	return errors.New(msg)
}

// NewBadRequestError - returns a new bad request error
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  400,
		Error:   "bad_request",
	}
}

// NewNotFoundError - returns a new not found error
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  404,
		Error:   "not_found",
	}
}

// NewInternalServerError - returns a new internal server error
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  500,
		Error:   "internal_server_error",
	}
}
