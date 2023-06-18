package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	http.ResponseWriter
	code int
}

func ResponseBuilder(w http.ResponseWriter) *response {
	return &response{w, 0}
}

func (r *response) Status(code int) *response {
	r.code = code
	return r
}

func (r *response) JSON(value interface{}) {
	j, err := json.Marshal(value)
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	r.Header().Add("Content-Type", "application/json")
	r.WriteHeader(r.code)
	r.Write(j)
}
