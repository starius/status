package status

import (
	"fmt"
)

type statusError struct {
	text string
	code int
}

func (e statusError) Error() string {
	return e.text
}

func Format(message string, arguments ...interface{}) error {
	code := 0
	if len(arguments) > 0 {
		if e, ok := arguments[len(arguments)-1].(statusError); ok {
			code = e.code
		}
	}
	return statusError{
		text: fmt.Sprintf(message, arguments...),
		code: code,
	}
}

func AttachCode(code int, err error) error {
	switch e := err.(type) {
	case statusError:
		e.code = code
		return e
	default:
		return statusError{
			text: err.Error(),
			code: code,
		}
	}
}

func WithCode(code int, message string, arguments ...interface{}) error {
	return AttachCode(code, Format(message, arguments...))
}

func Code(err error) int {
	switch e := err.(type) {
	case statusError:
		return e.code
	default:
		return 0
	}
}
