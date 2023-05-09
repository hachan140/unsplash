package httperror

import "net/http"

type Error struct {
	Status  int
	message string
}

func (e *Error) Error() string {
	return e.message
}

func NewError(status int, message string) error {
	return &Error{
		Status:  status,
		message: message,
	}
}

func BadRequest(message string) error {
	return NewError(http.StatusBadRequest, message)
}
