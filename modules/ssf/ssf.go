package ssf

import (
	"context"
	"net/http"

	"github.com/VikeLabs/uvic-api-go/modules/ssf/features"
	"github.com/VikeLabs/uvic-api-go/modules/ssf/features/bldgid"
	bldg "github.com/VikeLabs/uvic-api-go/modules/ssf/features/bldgs"
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	handlers := features.New(context.Background())

	r.Use(jsonHeader)
	r.Handle("/buildings", http.HandlerFunc(bldg.Controller))
	r.Handle("/buildings/{id}", http.HandlerFunc(bldgid.Controller))
	r.Handle("/rooms/{id}", http.HandlerFunc(handlers.GetRoomSchedule))
}
