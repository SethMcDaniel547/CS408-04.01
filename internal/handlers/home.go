package handlers

import (
	"HelloWorld/internal/models"
	"HelloWorld/internal/services"
	"HelloWorld/internal/views"

	"github.com/labstack/echo/v4"
)

// HomeHandler handles requests to the root path "/"
func HomeHandler(c echo.Context) error {
	// Fetch courses using our backend service layer (which reads the .env token)
	courses, err := services.FetchCourses()
	if err != nil {
		// If Canvas fails to load, you can pass an empty slice
		// so the view can gracefully display a "failed to connect" message
		courses = []models.Course{}
	}

	// Render the Templ view, passing the courses slice directly
	component := views.Home(courses)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
