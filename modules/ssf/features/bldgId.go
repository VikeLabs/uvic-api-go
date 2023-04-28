package features

import (
	"log"
	"net/http"
	"strconv"

	"github.com/VikeLabs/uvic-api-go/lib/api"
	"github.com/go-chi/chi/v5"
)

func BldgIDController(w http.ResponseWriter, r *http.Request) {
	bldgID, err := getBldgID(r)
	if err != nil {
		err.HandleError(w)
		return
	}

	val, err := getQuery("hour", r)
	if err != nil {
		err.HandleError(w)
		return
	}

	log.Println(bldgID, val)
}

func getQuery(key string, r *http.Request) (string, *api.Error) {
	val := r.URL.Query().Get(key)
	if val == "" {
		return "", api.ErrBadRequest(nil, "Missing query: "+key)
	}

	return val, nil
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
