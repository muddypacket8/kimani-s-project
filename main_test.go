package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test 1: Hello World endpoint
func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	helloHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, "Hello, World!") {
		t.Errorf("Expected 'Hello, World!' in response, got: %s", body)
	}
}

// Test 2: API endpoint returns JSON
func TestApiHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api", nil)
	w := httptest.NewRecorder()

	apiHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Expected JSON content type, got: %s", contentType)
	}

	body := w.Body.String()
	if !strings.Contains(body, "success") {
		t.Errorf("Expected 'success' status in response, got: %s", body)
	}
}

// Test 3: About endpoint
func TestAboutHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/about", nil)
	w := httptest.NewRecorder()

	aboutHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

// Test 4: 404 handler returns correct status code
func TestNotFoundHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/404", nil)
	w := httptest.NewRecorder()

	notFoundHandler(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}
