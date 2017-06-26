package main

import (
	"net/http"

	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
	"github.com/get-ion/ion/core/handlerconv"
)

func main() {
	app := ion.New()
	ionMiddleware := handlerconv.FromStd(nativeTestMiddleware)
	app.Use(ionMiddleware)

	// Method GET: http://localhost:8080/
	app.Get("/", func(ctx context.Context) {
		ctx.HTML("Home")
	})

	// Method GET: http://localhost:8080/ok
	app.Get("/ok", func(ctx context.Context) {
		ctx.HTML("<b>Hello world!</b>")
	})

	// http://localhost:8080
	// http://localhost:8080/ok
	app.Run(ion.Addr(":8080"))
}

func nativeTestMiddleware(w http.ResponseWriter, r *http.Request) {
	println("Request path: " + r.URL.Path)
}

// Look "routing/custom-context" if you want to convert a custom handler with a custom Context
// to a context.Handler.
