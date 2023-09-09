package features

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/VikeLabs/uvic-api-go/database"
	"github.com/VikeLabs/uvic-api-go/lib/api"
	"github.com/VikeLabs/uvic-api-go/modules/ssf/schemas"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type timeQueries struct {
	time uint64
	day  string
}

type session struct {
	TimeStartStr string `json:"time_start_str"`
	ID           uint64 `json:"room_id"`
	Room         string `json:"room" gorm:"room"`
	Subject      string `json:"subject"`
	TimeStartInt uint64 `json:"-"`
	TimeEndInt   uint64 `json:"-"`
}

func (db *state) BuildingID(w http.ResponseWriter, r *http.Request) {
	bldgID := chi.URLParam(r, "id")
	if bldgID == "" {
		err := api.ErrBadRequest(nil, "missing url param: building id")
		api.ResponseBuilder(w).Error(err)
		return
	}

	q, err := parseQueries(r)
	if err != nil {
		api.ResponseBuilder(w).Error(err)
		return
	}

	var buf schemas.BuildingSummary
	if err := db.getBuildingSchedules(q, bldgID, &buf); err != nil {
		api.ResponseBuilder(w).Error(err)
		return
	}

	api.ResponseBuilder(w).
		Status(http.StatusOK).
		JSON(buf)
}

func parseQueries(r *http.Request) (*timeQueries, *api.Error) {
	hourQuery := r.URL.Query().Get("hour")
	if hourQuery == "" {
		return nil, api.ErrBadRequest(nil, "Missing query: hour")
	}
	hour, err := strconv.ParseUint(hourQuery, 10, 8)
	if err != nil || hour > 24 {
		return nil, api.ErrBadRequest(err, "Bad value: hour")
	}

	minuteQuery := r.URL.Query().Get("minute")
	if minuteQuery == "" {
		return nil, api.ErrBadRequest(nil, "Missing query: minute")
	}
	minute, err := strconv.ParseUint(minuteQuery, 10, 8)
	if err != nil || minute > 60 || (hour == 24 && minute != 0) {
		return nil, api.ErrBadRequest(err, "Bad value: minute")
	}

	dayQuery := r.URL.Query().Get("day")
	if dayQuery == "" {
		return nil, api.ErrBadRequest(nil, "Missing query: day")
	}
	day, err := strconv.ParseUint(dayQuery, 0, 8)
	if err != nil || day > 6 {
		return nil, api.ErrBadRequest(err, "Bad value: day")
	}

	timeToSecond := hour*3600 + minute*60
	return &timeQueries{timeToSecond, getDay(uint8(day))}, nil
}

func getDay(day uint8) string {
	switch day {
	case 0:
		return "sunday"
	case 1:
		return "monday"
	case 2:
		return "tuesday"
	case 3:
		return "wednesday"
	case 4:
		return "thursday"
	case 5:
		return "friday"
	case 6:
		return "saturday"
	default:
		panic(fmt.Sprintf("invalid value, `day` less than 6, got %d", day))
	}
}

func (db *state) getBuildingSchedules(q *timeQueries, bldgID string, buf *schemas.BuildingSummary) *api.Error {
	var result *gorm.DB

	// get building
	var bldg schemas.Building
	result = db.
		Table(schemas.TableBuildings).
		Where("id = ?", bldgID).
		First(&bldg)

	if result.Error != nil {
		return api.ErrInternalServer(result.Error)
	}

	if result.RowsAffected == 0 {
		return api.ErrNotFound(nil, "invalid building id")
	}

	// get all rooms in building
	type Room struct {
		ID   uint64
		Room string
	}
	var rooms []Room
	result = db.
		Table(database.Rooms).
		Select("rooms.id", "rooms.room").
		Joins("JOIN buildings ON rooms.building_id=buildings.id").
		Where("buildings.id=?", bldgID).
		Order("room ASC").
		Scan(&rooms)

	if result.Error != nil {
		return api.ErrInternalServer(result.Error)
	}

	roomMap := make(map[uint64]Room)
	// get sessions in rooms
	for _, room := range rooms {
		var sessions []session
		sel := []string{
			"sections.time_start_str",
			"rooms.id",
			"rooms.room",
			"subjects.subject",
			"sections.time_start_int",
			"sections.time_end_int",
		}
		filter := map[string]any{
			"sections.room_id": room.ID,
			q.day:              true,
		}

		result = db.Table(database.Sections).
			Select(sel).
			Joins("JOIN rooms ON sections.room_id=rooms.id").
			Joins("JOIN subjects ON sections.subject_id=subjects.id").
			Where(filter).
			Where("sections.time_start_int >= ? OR sections.time_end_int >= ?", q.time, q.time).
			Order("time_start_int ASC").
			Limit(2).
			Scan(&sessions)

		if result.Error != nil {
			return api.ErrInternalServer(result.Error)
		}

		if len(sessions) == 0 {
			roomMap[room.ID] = room
			continue
		}

		nextSession := getNextSession(sessions, q.time)
		if nextSession == nil {
			roomMap[room.ID] = room
		}
		log.Println(sessions)
	}

	log.Println(roomMap)
	r := make([]schemas.RoomSummary, 0, len(roomMap))
	buf.Building = bldg.Name
	buf.Data = r
	return nil
}

/*
NOTE: len of `sessions` is either 1 or 2, from sql query (limit 2) and empty
session check
*/
func getNextSession(sessions []session, currentTime uint64) *schemas.RoomSummary {
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

func isInRange(timeStart, timeEnd, currentTime uint64) bool {
	inRangeStart := timeStart <= currentTime
	inRangeEnd := timeEnd >= currentTime
	return inRangeStart && inRangeEnd
}

func buildRoomInfo(session session, freeEod bool) *schemas.RoomSummary {
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
