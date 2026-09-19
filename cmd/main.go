package main

import (
	"context"
	"log"
	"os"

	"HelloWorld/internal/handlers"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// RenderTempl is your helper wrapper to render Templ components over Echo Context cleanly
func RenderTempl(c echo.Context, status int, cmp templ.Component) error {
	c.Response().Writer.WriteHeader(status)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(context.Background(), c.Response().Writer)
}

func main() {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on system environment variables")
	}

	e := echo.New()

	// Serve your compiled Tailwind file statically
	e.Static("/public", "public")

	// Page routes (wiring up to your handlers package)
	e.GET("/", handlers.HomeHandler)

	e.GET("/courses/:id/assignments", handlers.AssignmentsHandler)

	e.POST("/submit-assignment", handlers.SubmitAssignmentHandler)
	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on http://localhost:%s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
