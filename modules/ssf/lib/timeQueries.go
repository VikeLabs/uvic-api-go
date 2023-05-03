package lib

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/VikeLabs/uvic-api-go/lib/api"
)

type TimeQueries struct {
	Time uint64
	Day  string
}

func ParseQueries(r *http.Request) (*TimeQueries, *api.Error) {
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
	return &TimeQueries{timeToSecond, getDay(uint8(day))}, nil
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
