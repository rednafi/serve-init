package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type catFact struct {
	FactNo int    `json:"fact_no"`
	Ref    string `json:"ref"`
	Fact   string `json:"fact"`
}

type catFacts struct {
	Facts []catFact `json:"facts"`
}

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

var facts = catFacts{
	Facts: []catFact{
		{FactNo: 1, Fact: `Cats have five toes on their front paws, but only
four toes on their back paws.`},
		{FactNo: 2, Fact: "A group of cats is called a clowder."},
		{FactNo: 3, Fact: "Cats can rotate their ears 180 degrees."},
		{FactNo: 4, Fact: "The oldest cat on record lived to be 38 years old."},
		{FactNo: 5, Fact: "Cats have over 20 muscles that control their ears."},
	},
}

func catFactsHandler(authToken string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log the request details
		log.Printf("Received request: %s %s from %s", r.Method, r.URL, r.RemoteAddr)

		w.Header().Set("Content-Type", "application/json")
		token := r.Header.Get("Authorization")

		if token != "Token "+authToken {
			errorResponse := ErrorResponse{
				Type:    "error",
				Message: "Unauthorized. Please send the expected HTTP authorization header." +
				" The expected format is 'Authorization: Token token_value",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		// Encode and return the facts in JSON format
		if err := json.NewEncoder(w).Encode(facts); err != nil {
			errorResponse := ErrorResponse{
				Type:    "error",
				Message: err.Error(),
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
		}
	}
}

func main() {
	authToken := os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		log.Fatalf("AUTH_TOKEN is not set in the environment")
	}

	http.HandleFunc("/", catFactsHandler(authToken))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
