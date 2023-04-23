package middlewares

import (
	"fmt"
	"net/http"
)

func SetCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		maxAge := 10 * 60 // setting a 10 minute cache
		w.Header().Set("Cache-Control", "max-age="+fmt.Sprintf("%d", maxAge))
	})
}
