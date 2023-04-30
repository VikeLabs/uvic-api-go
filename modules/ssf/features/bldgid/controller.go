package bldgid

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VikeLabs/uvic-api-go/modules/ssf/lib"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	bldgID, err := getBldgID(r)
	if err != nil {
		err.HandleError(w)
		return
	}

	val, err := lib.ParseQueries(r)
	if err != nil {
		err.HandleError(w)
		return
	}

	data, err := getBuildingSchedules(val, bldgID)
	if err != nil {
		err.HandleError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
