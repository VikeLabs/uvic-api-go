package main

import (
	"log"
	"net/http"

	"github.com/VikeLabs/uvic-api-go/modules/ssf"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	mux := chi.NewMux()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"If-None-Match"},
		ExposedHeaders: []string{"ETag"},
	}))

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
