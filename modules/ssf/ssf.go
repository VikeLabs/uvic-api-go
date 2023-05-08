package ssf

import (
	"net/http"

	"github.com/VikeLabs/uvic-api-go/modules/ssf/features/bldgid"
	bldg "github.com/VikeLabs/uvic-api-go/modules/ssf/features/bldgs"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Route("/v0", func(r chi.Router) {
		r.Use(jsonHeader)
		r.Use(cache)
		r.Handle("/bldgs", http.HandlerFunc(bldg.Controller))
		r.Handle("/bldgs/{id}", http.HandlerFunc(bldgid.Controller))
	})
}
