package exception

import "net/http"

type BadRequestError struct {
	Message string
}

func (e *BadRequestError) Code() int {
	return http.StatusBadRequest
}

func (e *BadRequestError) Status() string {
	return http.StatusText(http.StatusBadRequest)
}

func (e *BadRequestError) Error() string {
	return e.Message
}
