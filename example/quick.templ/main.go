package main

import (
	"hello/views"

	"github.com/a-h/templ"
	"github.com/jeffotoni/quick"
)

func Render(c *quick.Ctx, t templ.Component) error {
	return t.Render(c.Request.Context(), c.Response)
}

func main() {
	app := quick.New()

	app.Get("/", func(c *quick.Ctx) error {
		return Render(c, views.Index())
	})

	app.Listen("0.0.0.0:8080")
}
