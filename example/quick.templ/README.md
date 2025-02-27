# Quick + Templ

How to build a fullstack application with Go, Quick + Templ

```go
// main.go
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
```


```go
// views/index.templ
package views

templ Index() {
	<!DOCTYPE html>
	<html>
		<head>
			<title>Quick + Templ</title>
		</head>
		<body>
			<h1>Hello World</h1>
			<p>Templ + Quick</p>
		</body>
	</html>
}

```

```bash
go install github.com/a-h/templ/cmd/templ@latest
go mod init hello
go mod tidy

templ fmt views
templ generate
go build
./hello
```

```bash
http://localhost:8080
```