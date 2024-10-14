package main

import (
	"github.com/jordyvanvorselen/travel-lists/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.Handler)

	log.Println("Started server on http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
