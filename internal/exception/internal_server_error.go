package exception

import "net/http"

type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Code() int {
	return http.StatusInternalServerError
}

func (e *InternalServerError) Status() string {
	return http.StatusText(http.StatusInternalServerError)
}

func (e *InternalServerError) Error() string {
	return e.Message
}
