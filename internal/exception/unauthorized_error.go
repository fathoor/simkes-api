package exception

import "net/http"

type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Code() int {
	return http.StatusUnauthorized
}

func (e *UnauthorizedError) Status() string {
	return http.StatusText(http.StatusUnauthorized)
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}
