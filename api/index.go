package handler

import (
	"net/http"

	"github.com/VikeLabs/uvic-api-go/modules/ssf"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mux := chi.NewMux()

	mux.Use(middleware.Recoverer, middleware.Logger)

	mux.Route("/ssf", ssf.Router)

	mux.ServeHTTP(w, r)
}
