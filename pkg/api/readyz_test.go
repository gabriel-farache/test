package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestReadyz(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/readyz", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create a context for the request
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the function we want to test
	Readyz(c)

	// Check the status code is what we expect
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	// Check the response body is what we expect
	expected := `{"ready":true}`
	if w.Body.String() != expected {
		t.Fatalf("Expected to get '%s' but instead got '%s'\n", expected, w.Body.String())
	}
}
