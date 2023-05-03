package bldgid

import (
	"context"
	"encoding/json"
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
	// TODO: handle error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := newDB(ctx)

	bldg := schemas.Building{ID: bldgID}
	if err := db.getBuildingName(&bldg); err != nil {
		panic(err)
	}

	// Get all rooms
	var rooms []schemas.RoomSummary
	if err := db.getRooms(bldgID, &rooms); err != nil {
		panic(err)
	}

	// Get sessions per room
	var out []RoomSchedule
	for _, room := range rooms {
		var buf []RoomSchedule
		if err := db.getRoomSchedule(room.ID, lib.GetDay(query.Day), &buf); err != nil {
			if errors.Is(err, ErrNoData) {
				continue
			}
			panic(err)
		}
		log.Println(buf)
		// buf.RoomID = room.ID
		// buf.RoomName = room.Room
		// out = append(out, buf)
	}

	for _, session := range out {
		buf, _ := json.MarshalIndent(session, "", "  ")
		log.Println(string(buf))
	}

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
