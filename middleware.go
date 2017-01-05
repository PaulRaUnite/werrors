package service_errors

import "github.com/valyala/fasthttp"

type RequiredFunc func(*fasthttp.RequestCtx) Error

func NewErrorMiddleware(f RequiredFunc) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		err := f(ctx)
		if err.error != Nil {
			ctx.SetBodyString(err.Error())
			ctx.SetStatusCode(err.code)
		}
	}
}
