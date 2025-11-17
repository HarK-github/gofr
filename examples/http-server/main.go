package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gofr.dev/pkg/gofr"
)

func TestFailingRoute(t *testing.T) {
	// Create a test app instance
	app := gofr.New()

	// Add a route that we know will fail
	app.GET("/failing-test", func(c *gofr.Context) (interface{}, error) {
		// This will cause a division by zero panic when accessed
		var zero int
		result := 1 / zero
		return result, nil
	})

	// Create a test request for the failing route
	req := httptest.NewRequest("GET", "/failing-test", nil)
	w := httptest.NewRecorder()

	// This should cause a panic and fail the test
	defer func() {
		if r := recover(); r == nil {
			// If we didn't panic, that's unexpected - fail the test
			t.Errorf("The code did not panic as expected")
		}
	}()

	app.Server.HTTP.ServeHTTP(w, req)
}

func TestAlwaysFailingAssertion(t *testing.T) {
	// This test will always fail due to a false assertion
	expected := "hello"
	actual := "world"

	if expected != actual {
		t.Errorf("This test is designed to fail. Expected: %s, Got: %s", expected, actual)
	}
}

func TestPanicInHandler(t *testing.T) {
	app := gofr.New()

	// Add a handler that explicitly panics
	app.GET("/panic-route", func(c *gofr.Context) (interface{}, error) {
		panic("This is a deliberate panic to fail the test")
	})

	req := httptest.NewRequest("GET", "/panic-route", nil)
	w := httptest.NewRecorder()

	// Serve the request - this should handle the panic but we can still make the test fail
	app.Server.HTTP.ServeHTTP(w, req)

	// Even if the framework recovers from panics, we can fail the test explicitly
	// by checking for unexpected success
	if w.Code == http.StatusOK {
		t.Errorf("Expected non-200 status code due to panic, got: %d", w.Code)
	}
}
