package ssf

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func bldgs(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"hello": "world"})
}

func Router(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Handle("/bldgs", http.HandlerFunc(bldgs))
	})
}
