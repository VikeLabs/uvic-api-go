package main

import (
	"log"
	"net/http"

	api "github.com/VikeLabs/uvic-api-go/api"
)

// NOTE: This file is for running docker on development purposes only, don't edit
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(api.Handler))
	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
