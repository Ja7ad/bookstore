package errors

// RestErr is a collection of errors
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
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
