package ssf

import (
	"net/http"

	"github.com/VikeLabs/uvic-api-go/modules/ssf/features"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Route("/v0", func(r chi.Router) {
		r.Use(jsonHeader)
		r.Handle("/bldgs", http.HandlerFunc(features.BldgsController))
	})
}
