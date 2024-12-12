package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware logs every HTTP request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		// Log the request details
		log.Printf(
			"Method: %s, Path: %s, ",
			r.Method, r.URL.Path,
		)
	})
}
