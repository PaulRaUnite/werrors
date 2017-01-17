package service_errors

import (
	"log"

	"github.com/valyala/fasthttp"
)

type RequiredFunc func(*fasthttp.RequestCtx) error

func NewErrorMiddleware(f RequiredFunc) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		err := f(ctx)
		if err != nil {
			log.Println(err.Error())
			httpErr, ok := (err).(Error)
			if ok {
				ctx.SetBodyString(httpErr.Marshal())
				ctx.SetStatusCode(httpErr.code)
				return
			} else {
				ctx.SetBodyString("{error: \"internal server error\"}")
				ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			}
		}
	}
}
