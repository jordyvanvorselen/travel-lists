package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/jordyvanvorselen/travel-lists/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	http.HandleFunc("/", api.Handler)

	log.Println("Started server and proxy (frontend with live updates).\n\nBackend: http://localhost:3000\nFrontend: http://localhost:8080")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
