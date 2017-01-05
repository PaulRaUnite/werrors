package service_errors

import (
	"strconv"

	"github.com/valyala/fasthttp"
)

var Nil = Error{fasthttp.StatusOK, nil}

type Error struct {
	//unexported fields
	code  int
	error error
}

func (e Error) Error() string {
	return "{code: " + strconv.Itoa(e.code) + ", expl: \"" + e.error.Error() + "\"}"
}

func New(error error, code int) Error {
	return Error{code, error}
}
