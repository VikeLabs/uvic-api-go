package ssf

import (
	"net/http"

	"github.com/VikeLabs/uvic-api-go/modules/ssf/features"
	// "github.com/VikeLabs/uvic-api-go/modules/ssf/features/bldgid"

	// bldg "github.com/VikeLabs/uvic-api-go/modules/ssf/features/bldgs"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	handlers, err := features.New()
	if err != nil {
		panic(err)
	}

	r.Use(jsonHeader)
	// r.Handle("/buildings", http.HandlerFunc(bldg.Controller))
	// r.Handle("/buildings", http.HandlerFunc(handlers.BuildingID))
	// r.Handle("/buildings/{id}", http.HandlerFunc(bldgid.Controller))
	r.Handle("/buildings/{id}", http.HandlerFunc(handlers.BuildingID))
	r.Handle("/rooms/{id}", http.HandlerFunc(handlers.GetRoomSchedule))
}
