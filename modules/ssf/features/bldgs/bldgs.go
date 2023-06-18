package bldg

import (
	"encoding/json"
	"net/http"

	"github.com/VikeLabs/uvic-api-go/lib/api"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	bldgs, err := bldgsService()
	if err != nil {
		api.ResponseBuilder(w).Error(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&bldgs)
}
