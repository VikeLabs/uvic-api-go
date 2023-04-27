package main

import (
	"log"
	"net/http"

	"github.com/VikeLabs/uvic-api-go/modules/ssf"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	mux := chi.NewMux()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)

	mux.Group(func(r chi.Router) {
		r.Route("/ssf", ssf.Router)
	})

	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
