package ssf

import (
	"log"
	"net/http"

	"github.com/VikeLabs/uvic-api-go/database"
)

func jsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func cache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, err := database.GetVersion()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		resourceTag := r.Header.Get("If-None-Match")
		log.Println(resourceTag)
		if resourceTag == v.StudySpaceFinder && resourceTag != "" {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		w.Header().Set("ETag", v.StudySpaceFinder)
		next.ServeHTTP(w, r)
	})
}
