package ssf

import (
	"net/http"

	"github.com/VikeLabs/uvic-api-go/config"
)

func jsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func cache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, err := config.GetVersions()
		if err != nil {
			panic(err)
		}

		cxTag := r.Header.Get("If-None-Match")
		if cxTag == v.Database.Banner {
			w.WriteHeader(http.StatusNotModified)
			return
		}

		w.Header().Set("ETag", v.Database.Banner)
		next.ServeHTTP(w, r)
	})
}
