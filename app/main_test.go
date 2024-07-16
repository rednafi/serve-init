package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unicode/utf8"
)

// TestHandler tests the /catfacts endpoint
func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/catfacts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(facts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the Content-Type header
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("handler returned wrong Content-Type: got %v want %v",
			contentType, "application/json")
	}

	// Check the response body
	var response CatFacts
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("could not parse response: %v", err)
	}

	// Check the number of facts
	if len(response.Facts) != len(facts.Facts) {
		t.Errorf("handler returned wrong number of facts: got %v want %v",
			len(response.Facts), len(facts.Facts))
	}

	// Check each fact
	for i, fact := range response.Facts {
		if fact.FactNo != facts.Facts[i].FactNo {
			t.Errorf("handler returned wrong fact_no: got %v want %v",
				fact.FactNo, facts.Facts[i].FactNo)
		}
		if fact.Ref == "" {
			t.Errorf("handler returned empty ref for fact_no %v", fact.FactNo)
		}
		if fact.Fact != facts.Facts[i].Fact {
			t.Errorf("handler returned wrong fact: got %v want %v",
				fact.Fact, facts.Facts[i].Fact)
		}
	}
}

// TestGenerateRef tests the generateRef function
func TestGenerateRef(t *testing.T) {
	ref := generateRef()

	// Check the length of the ref
	if utf8.RuneCountInString(ref) != 8 {
		t.Errorf("generateRef returned wrong length: got %v want 8", len(ref))
	}

	// Check if the ref contains only valid characters
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, char := range ref {
		if !strings.ContainsRune(charset, char) {
			t.Errorf("generateRef returned invalid character: %v", char)
		}
	}
}
