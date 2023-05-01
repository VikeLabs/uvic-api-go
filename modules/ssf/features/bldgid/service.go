package bldgid

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/VikeLabs/uvic-api-go/lib/api"
	"github.com/VikeLabs/uvic-api-go/modules/ssf/lib"
	"github.com/VikeLabs/uvic-api-go/modules/ssf/schemas"
	"github.com/go-chi/chi/v5"
)

func getBuildingSchedules(query *lib.TimeQueries, bldgID uint64) (*schemas.BuildingSummary, *api.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := newDB(ctx)

	bldg := schemas.Building{ID: bldgID}
	if err := db.getBuildingName(&bldg); err != nil {
		if errors.Is(err, ErrBadQuery) {
			return nil, api.ErrBadRequest(err, "Missing building ID")
		}
		return nil, api.ErrNotFound(err, "Building not found")
	}

	// Get all rooms
	var rooms []schemas.RoomSummary
	if err := db.getRooms(bldg.ID, &rooms); err != nil {
		log.Fatal(err) // TODO: handle err
	}

	// Get sessions per room

	var out []RoomSchedule
	for _, room := range rooms {
		log.Println(room)
		var buf RoomSchedule
		if err := db.getRoomSchedule(room.ID, query.Day, &buf); err != nil {
			log.Fatal(err) // TODO: handle err
		}

		out = append(out, buf)
	}

	log.Println(out)
	return nil, nil
}

func getBldgID(r *http.Request) (uint64, *api.Error) {
	param := chi.URLParam(r, "id")
	if param == "" {
		return 0, api.ErrBadRequest(nil, "Missing building path param.")
	}

	id, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return 0, api.ErrBadRequest(err, "Failed to parse building id.")
	}

	return id, nil
}
