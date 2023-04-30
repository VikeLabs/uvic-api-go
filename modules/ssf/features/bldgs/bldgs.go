package bldg

import (
	"encoding/json"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	bldgs, err := bldgsService()
	if err != nil {
		err.HandleError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&bldgs)
}
