package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		start := time.Now()

		next(rw, r)

		log.Printf(
			"%s %s -> %d (%s)",
			r.Method,
			r.URL.Path,
			rw.status,
			time.Since(start),
		)
	}
}
