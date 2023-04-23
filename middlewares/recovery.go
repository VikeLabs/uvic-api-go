package middlewares

import (
	"log"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		if it, ok := recover().(string); ok {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(it)
		}
	})
}
