package handlers

import (
	"HelloWorld/internal/models"
	"HelloWorld/internal/services"
	"HelloWorld/internal/views"

	"github.com/labstack/echo/v4"
)

// AssignmentsHandler handles GET /courses/:id/assignments requests via HTMX
func AssignmentsHandler(c echo.Context) error {
	courseID := c.Param("id")

	assignments, err := services.FetchAssignments(courseID)
	if err != nil {
		assignments = []models.Assignment{}
	}

	// Render just the assignments fragment/view to swap into the DOM
	component := views.AssignmentView(courseID, assignments)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
