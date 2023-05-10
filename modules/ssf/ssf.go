package ssf

import (
	"net/http"

	"github.com/VikeLabs/uvic-api-go/modules/ssf/features/bldgid"
	bldg "github.com/VikeLabs/uvic-api-go/modules/ssf/features/bldgs"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Use(jsonHeader)
	r.Use(cache)
	r.Handle("/buildings", http.HandlerFunc(bldg.Controller))
	r.Handle("/buildings/{id}", http.HandlerFunc(bldgid.Controller))
}
