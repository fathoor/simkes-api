package exception

import "net/http"

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Code() int {
	return http.StatusNotFound
}

func (e *NotFoundError) Status() string {
	return http.StatusText(http.StatusNotFound)
}

func (e *NotFoundError) Error() string {
	return e.Message
}
