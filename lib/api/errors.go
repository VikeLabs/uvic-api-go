package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Err        error
	StatusCode int
	Msg        string
}

func newErr(err error, statusCode int, msg string) *Error {
	if err != nil {
		log.Println(err)
	}

	return &Error{err, statusCode, msg}
}

func ErrNotFound(err error, msg string) *Error {
	return newErr(err, http.StatusNotFound, msg)
}

func ErrBadRequest(err error, msg string) *Error {
	return newErr(err, http.StatusBadRequest, msg)
}

func ErrInternalServer(err error) *Error {
	return newErr(err, http.StatusInternalServerError, "")
}

func (err *Error) HandleError(w http.ResponseWriter) {
	w.WriteHeader(err.StatusCode)

	errMsg := map[string]string{"error": "err.Msg"}
	if err := json.NewEncoder(w).Encode(&errMsg); err != nil {
		panic(err)
	}
}
