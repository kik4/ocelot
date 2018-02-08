package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func handler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, World")
}

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			handler(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}
	fasthttp.ListenAndServe(":8080", requestHandler)
}
