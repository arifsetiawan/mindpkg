package apierror

import "errors"

// APIError ...
// e, ok := err.(*apierror.APIError)
type APIError struct {
	HTTPStatus int    `json:"-"`
	Code       int    `json:"code,omitempty"`
	Message    string `json:"error"`
	Err        error  `json:"-"`
}

func (e *APIError) Error() string {
	return e.Message
}

// NewError ...
func NewError(status int, message string) *APIError {
	return &APIError{
		HTTPStatus: status,
		Code:       status,
		Message:    message,
		Err:        errors.New(message),
	}
}

// NewErrorWithInternal ...
func NewErrorWithInternal(status int, message string, internal error) *APIError {
	return &APIError{
		HTTPStatus: status,
		Code:       status,
		Message:    message,
		Err:        internal,
	}
}
