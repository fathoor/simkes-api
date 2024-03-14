package exception

import "net/http"

type ForbiddenError struct {
	Message string
}

func (e *ForbiddenError) Code() int {
	return http.StatusForbidden
}

func (e *ForbiddenError) Status() string {
	return http.StatusText(http.StatusForbidden)
}

func (e *ForbiddenError) Error() string {
	return e.Message
}
