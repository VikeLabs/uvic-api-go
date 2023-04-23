package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/VikeLabs/uvic-api-go/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
}

func main() {
	mux := chi.NewRouter()

	// middlewares
	mux.Use(middlewares.Recovery)
	mux.Use(middlewares.SetContent)
	mux.Use(middlewares.SetCache)
	mux.Use(middleware.Logger)

	fmt.Println("Listening on port", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
