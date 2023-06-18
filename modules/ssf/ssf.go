package ssf

import (
	"net/http"

	"github.com/VikeLabs/uvic-api-go/modules/ssf/features"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	handlers, err := features.New()
	if err != nil {
		panic(err)
	}

	r.Handle("/buildings", http.HandlerFunc(handlers.Buildings))
	r.Handle("/buildings/{id}", http.HandlerFunc(handlers.BuildingID))
	r.Handle("/rooms/{id}", http.HandlerFunc(handlers.GetRoomSchedule))
}
