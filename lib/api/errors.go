package api

import (
	"log"
	"net/http"
)

type Error struct {
	err        error
	statusCode int
	msg        string
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
