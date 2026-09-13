package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"HelloWorld/views" // Adjust to match your go.mod module name

	"github.com/labstack/echo/v4"
)

// 1. UNIT TEST: Verifying the Templ Component directly
func TestHelloFragmentComponent(t *testing.T) {
	// Initialize a buffer to catch the rendered HTML string
	w := &strings.Builder{}

	// Render the individual component fragment
	component := views.HelloFragment()
	err := component.Render(context.Background(), w)
	if err != nil {
		t.Fatalf("failed to render component: %v", err)
	}

	resultHTML := w.String()

	// Assertions: Ensure the fragment contains your structural text and Tailwind animations
	if !strings.Contains(resultHTML, "Hello from your Echo server!") {
		t.Errorf("expected HTML to contain success message, got: %s", resultHTML)
	}
	if !strings.Contains(resultHTML, "animate-bounce") {
		t.Errorf("expected HTML to contain Tailwind animation utility class 'animate-bounce'")
	}
}

// 2. INTEGRATION TEST: Testing the Echo Routing and HTMX Endpoint Response
func TestHelloAPIHandler(t *testing.T) {
	// Setup a clean Echo instance
	e := echo.New()

	// Create a dummy HTTP GET request simulating an HTMX call to our fragment endpoint
	req := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Define the handler logic inline (exactly matching your server setup)
	handler := func(c echo.Context) error {
		return RenderTempl(c, http.StatusOK, views.HelloFragment())
	}

	// Execute the handler
	if err := handler(c); err != nil {
		t.Fatalf("handler returned unexpected error: %v", err)
	}

	// Assert Response Status Code is 200 OK
	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	// Assert Response Header is text/html
	contentType := rec.Header().Get(echo.HeaderContentType)
	if !strings.Contains(contentType, echo.MIMETextHTML) {
		t.Errorf("expected content type to be text/html, got %s", contentType)
	}

	// Assert Response Body contents
	body, _ := io.ReadAll(rec.Result().Body)
	bodyStr := string(body)
	if !strings.Contains(bodyStr, "animate-bounce") {
		t.Errorf("expected dynamic HTMX response fragment to contain styling, got: %s", bodyStr)
	}
}
