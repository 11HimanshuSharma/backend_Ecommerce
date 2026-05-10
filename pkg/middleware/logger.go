package middleware

import (
	"log"
	"net/http"
	"time"
)

// logging logs every request with method, path and how long  it took

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
		duration := time.Since(start)

		log.Printf("%s %s took %v", r.Method, r.URL.Path, duration
		)
	})
}