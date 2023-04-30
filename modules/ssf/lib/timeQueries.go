package lib

import (
	"net/http"
	"strconv"

	"github.com/VikeLabs/uvic-api-go/lib/api"
)

type TimeQueries struct {
	Hour   uint64
	Minute uint64
	Day    uint64
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

	return &TimeQueries{hour, minute, day}, nil
}
