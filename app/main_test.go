package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Set up a dummy auth token for testing
	os.Setenv("AUTH_TOKEN", "dummy_token")
	code := m.Run()
	os.Unsetenv("AUTH_TOKEN")
	os.Exit(code)
}

func TestAuthorizedRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Token dummy_token")

	rr := httptest.NewRecorder()
	handler := catFactsHandler("dummy_token")

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedFacts := facts
	var actualFacts catFacts
	if err := json.NewDecoder(rr.Body).Decode(&actualFacts); err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	if len(actualFacts.Facts) != len(expectedFacts.Facts) {
		t.Errorf("handler returned unexpected number of facts: got %v want %v", len(actualFacts.Facts), len(expectedFacts.Facts))
	}
}

func TestUnauthorizedRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := catFactsHandler("dummy_token")

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	var errorResponse ErrorResponse
	if err := json.NewDecoder(rr.Body).Decode(&errorResponse); err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	expectedError := ErrorResponse{
		Type:    "error",
		Message: "Unauthorized. Please send the expected HTTP authorization header. The expected format is 'Authorization: Token token_value",
	}

	if errorResponse != expectedError {
		t.Errorf("handler returned unexpected error response: got %v want %v", errorResponse, expectedError)
	}
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	healthCheckHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
