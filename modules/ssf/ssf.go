package ssf

import (
	"encoding/json"
	"net/http"

	"github.com/VikeLabs/uvic-api-go/modules/ssf/features"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Route("/v0", func(r chi.Router) {
		r.Use(jsonHeader)
		r.Handle("/bldgs", http.HandlerFunc(bldgs))
	})
}

func bldgs(w http.ResponseWriter, r *http.Request) {
	bldgs, err := features.GetBuildings()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "eroo"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&bldgs)
}
