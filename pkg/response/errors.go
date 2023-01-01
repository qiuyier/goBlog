package response

import "strconv"

func NewError(code int, text string) *CodeError {
	return &CodeError{code, text, nil}
}

type CodeError struct {
	Code    int
	Message string
	Data    interface{}
}

func (e *CodeError) Error() string {
	return strconv.Itoa(e.Code) + ": " + e.Message
}
