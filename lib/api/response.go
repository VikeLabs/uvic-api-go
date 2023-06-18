package api

import (
	"encoding/json"
	"net/http"
	"os"
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
	if r.code == 0 {
		panic("status code not set.")
	}

	j, err := json.Marshal(value)
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		os.Stderr.WriteString(err.Error())
		return
	}

	r.Header().Add("Content-Type", "application/json")
	r.WriteHeader(r.code)
	r.Write(j)
}

func (r *response) Error(apiErr *Error) {
	if apiErr.err != nil {
		_, err := os.Stderr.WriteString(apiErr.err.Error())
		if err != nil {
			r.WriteHeader(http.StatusInternalServerError)
			os.Stderr.WriteString(err.Error())
			return
		}
	}

	j, err := json.Marshal(map[string]string{"error": apiErr.msg})
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		os.Stderr.WriteString(err.Error())
		return
	}

	r.Header().Add("Content-Type", "application/json")
	if r.code != 0 {
		r.WriteHeader(apiErr.statusCode)
	} else {
		r.WriteHeader(r.code)
	}

	r.Write(j)
}
