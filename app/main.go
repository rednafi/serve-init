package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type CatFact struct {
	FactNo int    `json:"fact_no"`
	Ref    string `json:"ref"`
	Fact   string `json:"fact"`
}

type CatFacts struct {
	Facts []CatFact `json:"facts"`
}

func generateRef() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

var facts = CatFacts{
	Facts: []CatFact{
		{FactNo: 1, Ref: generateRef(), Fact: "Cats have five toes on their front paws, but only four toes on their back paws."},
		{FactNo: 2, Ref: generateRef(), Fact: "A group of cats is called a clowder."},
		{FactNo: 3, Ref: generateRef(), Fact: "Cats can rotate their ears 180 degrees."},
		{FactNo: 4, Ref: generateRef(), Fact: "The oldest cat on record lived to be 38 years old."},
		{FactNo: 5, Ref: generateRef(), Fact: "Cats have over 20 muscles that control their ears."},
	},
}

func main() {
	authToken := os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		log.Fatalf("AUTH_TOKEN is not set in the environment")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		// Send error in JSON format
		if token != "Token "+authToken {
			http.Error(
				w,
				"Unauthorized. Please send the expected HTTP authorization header.\n"+
					"The expected format is 'Authorization: Token <token>'",
				http.StatusUnauthorized,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(facts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
