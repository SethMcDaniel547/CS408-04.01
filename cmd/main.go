package main

import (
	"HelloWorld/views" // Adjust this to match your go.mod module name
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Simple middleware/helper wrapper to render Templ components over Echo Context
func RenderTempl(c echo.Context, status int, cmp templ.Component) error {
	c.Response().Writer.WriteHeader(status)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(context.Background(), c.Response().Writer)
}

func main() {
	e := echo.New()

	// Serve your compiled Tailwind file statically
	e.Static("/public", "public")

	// Page route
	e.GET("/", func(c echo.Context) error {
		return RenderTempl(c, http.StatusOK, views.Home())
	})

	// HTMX API fragment route
	e.GET("/api/hello", func(c echo.Context) error {
		return RenderTempl(c, http.StatusOK, views.HelloFragment())
	})

	e.Logger.Fatal(e.Start(":8080"))
}
