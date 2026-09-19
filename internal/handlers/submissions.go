package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"HelloWorld/internal/services"

	"github.com/labstack/echo/v4"
)

func SubmitAssignmentHandler(c echo.Context) error {
	courseID := c.FormValue("course_id")
	assignmentIDStr := c.FormValue("assignment_id")
	submissionURL := c.FormValue("submission_url")

	assignmentID, err := strconv.Atoi(assignmentIDStr)
	if err != nil {
		return c.HTML(http.StatusBadRequest, "<p class='text-red-400'>Invalid Assignment ID</p>")
	}

	err = services.SubmitAssignment(courseID, assignmentID, submissionURL)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, fmt.Sprintf("<p class='text-red-400'>Failed to submit: %v</p>", err))
	}

	// Return a success banner fragment via HTMX
	return c.HTML(http.StatusOK, `<div class="p-4 bg-emerald-900/50 border border-emerald-500 text-emerald-300 rounded-lg text-center font-medium">Assignment successfully submitted to Canvas!</div>`)
}
