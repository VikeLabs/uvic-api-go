package features

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/VikeLabs/uvic-api-go/lib/api"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type timeQueries struct {
	time uint64
	day  string
}

type session struct {
	ID           uint64 `json:"id"`
	Subject      string `json:"subject"`
	Description  string `json:"description"`
	TimeStartStr string `json:"time_start"`
	TimeEndStr   string `json:"time_end"`
	Room         string `json:"-" gorm:"-"`
	TimeStartInt uint64 `json:"-"`
	TimeEndInt   uint64 `json:"-"`
}

type roomInfo struct {
	ID           uint64   `json:"id"`
	Room         string   `json:"room"`
	NextClass    *session `json:"next_class" gorm:"-"`
	CurrentClass *session `json:"current_class" gorm:"-"`
}

func (db *state) BuildingID(w http.ResponseWriter, r *http.Request) {
	bldgID, _err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 8)
	if _err != nil {
		api.ResponseBuilder(w).Error(api.ErrInternalServer(_err))
		return
	}

	q, err := parseQueries(r)
	if err != nil {
		api.ResponseBuilder(w).Error(api.ErrBadRequest(err, err.Error()))
		return
	}

	// get building room ids
	var rooms []roomInfo
	result := db.
		Table(tableRooms).
		Select("rooms.id", "rooms.room").
		Joins("JOIN buildings ON rooms.building_id=buildings.id").
		Where("buildings.id=?", bldgID).
		Order("room ASC").
		Scan(&rooms)
	if err := result.Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		api.ResponseBuilder(w).Error(api.ErrInternalServer(err))
	}

	// get session in room
	for i, room := range rooms {
		var sessions []session
		sel := []string{
			"sections.time_start_str",
			"sections.time_end_str",
			"sections.id",
			"rooms.room",
			"subjects.subject",
			"subjects.description",
			"sections.time_start_int",
			"sections.time_end_int",
		}
		filter := map[string]any{
			"sections.room_id": room.ID,
			q.day:              true,
		}

		result = db.Table(tableSections).
			Select(sel).
			Joins("JOIN rooms ON sections.room_id=rooms.id").
			Joins("JOIN subjects ON sections.subject_id=subjects.id").
			Where(filter).
			Where("sections.time_start_int >= ? OR sections.time_end_int >= ?", q.time, q.time).
			Order("time_start_int ASC").
			Limit(2).
			Scan(&sessions)

		// if no session is found, free til eod
		// if time_end_int > current_time -> busy until time_end_int
		// if time_start_int > current_time -> free until time_start_int
		if len(sessions) == 0 {
			continue
		}

		if len(sessions) == 1 {
			if sessions[0].TimeStartInt > q.time {
				rooms[i].NextClass = &sessions[0]
			} else {
				rooms[i].CurrentClass = &sessions[0]
			}
			continue
		}

		if sessions[0].TimeStartInt >= q.time {
			rooms[i].NextClass = &sessions[0]
			continue
		}

		if sessions[0].TimeEndInt >= q.time {
			rooms[i].CurrentClass = &sessions[0]
		} else if sessions[1].TimeStartInt >= q.time {
			rooms[i].NextClass = &sessions[1]
		} else {
			rooms[i].CurrentClass = &sessions[1]
		}
	}

	// sort by room
	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i].Room < rooms[j].Room
	})

	api.ResponseBuilder(w).Status(http.StatusOK).JSON(&rooms)
}

func parseQueries(r *http.Request) (*timeQueries, error) {
	hourQuery := r.URL.Query().Get("hour")
	hour, err := strconv.ParseUint(hourQuery, 10, 8)
	if err != nil || hour > 24 {
		return nil, err
	}

	minuteQuery := r.URL.Query().Get("minute")
	minute, err := strconv.ParseUint(minuteQuery, 10, 8)
	if err != nil || minute > 60 || (hour == 24 && minute != 0) {
		return nil, err
	}

	dayQuery := r.URL.Query().Get("day")
	day, err := strconv.ParseUint(dayQuery, 0, 8)
	if err != nil {
		return nil, err
	}

	if day > 6 {
		return nil, errors.New("invalid day query")
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
