package bldgid

import (
	"context"
	"errors"
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

	// Get building name
	bldg := schemas.Building{ID: bldgID}
	if err := db.getBuildingName(&bldg); err != nil {
		return nil, api.ErrBadRequest(err, "Building not found")
	}

	// Get all rooms
	var rooms []schemas.RoomSummary
	if err := db.getRooms(bldgID, &rooms); err != nil {
		return nil, api.ErrInternalServer(err)
	}

	// Get sessions per room
	for _, room := range rooms {
		var sessions []Session
		if err := db.getRoomSchedule(room.ID, query, &sessions); err != nil {
			if errors.Is(err, ErrNoData) {
				continue
			}
			return nil, api.ErrInternalServer(err)
		}

		nextSession := getNextSession(sessions, query.Time)
		if nextSession != nil {
			rooms = append(rooms, *nextSession)
		}
	}

	return &schemas.BuildingSummary{Building: bldg.Name, Data: rooms}, nil
}

/*
NOTE: len of `sessions` is either 1 or 2, from sql query (limit 2) and empty
session check
*/
func getNextSession(sessions []Session, currentTime uint64) *schemas.RoomSummary {
	// one class found
	if len(sessions) == 1 {
		session := sessions[0]
		if isInRange(session.TimeStartInt, session.TimeEndInt, currentTime) {
			return nil
		}
		freeEod := currentTime > session.TimeEndInt
		return buildRoomInfo(session, freeEod)
	}

	prevSession := sessions[0]
	if isInRange(prevSession.TimeStartInt, prevSession.TimeEndInt, currentTime) {
		return nil
	}

	nextSession := sessions[1]
	inBetweenSessions := isInRange(
		prevSession.TimeEndInt,
		nextSession.TimeStartInt,
		currentTime,
	)

	if inBetweenSessions {
		return buildRoomInfo(nextSession, false)
	}

	if nextSession.TimeEndInt <= currentTime {
		return buildRoomInfo(nextSession, true)
	}

	return nil
}

func buildRoomInfo(session Session, freeEod bool) *schemas.RoomSummary {
	nextSession := &session.TimeStartStr
	subject := &session.Subject

	if freeEod {
		nextSession = nil
		subject = nil
	}

	return &schemas.RoomSummary{
		ID:        session.ID,
		Room:      session.Room,
		NextClass: nextSession,
		Subject:   subject,
	}
}

func isInRange(timeStart, timeEnd, currentTime uint64) bool {
	inRangeStart := timeStart <= currentTime
	inRangeEnd := timeEnd >= currentTime
	return inRangeStart && inRangeEnd
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
