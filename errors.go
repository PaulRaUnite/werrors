package service_errors

import (
	"errors"
)

type Error struct {
	//unexported fields
	code  int
	error error
}

func (e Error) Error() string {
	return e.error.Error()
}

func (e Error) Marshal() string {
	return "{error: \"" + e.error.Error() + "\"}"
}

func New(error string, code int) Error {
	return Error{code, errors.New(error)}
}
