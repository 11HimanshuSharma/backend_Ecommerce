package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// Defie a custom type for our context key
// why a custom type ? if two packages use the same string key like "id",
// they collide and overwrite each other, using custom type prevents this

type contextKey string


const RequestIDKey contextKey = "request_id"


// RequestId assigns a unique ID to every incoming requestt
// Why ? when you have 100000 requests/min and user says "my order failed"
// you search longs for their specific request_id to find exactly what happened


func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		id := uuid.New().String()

		w.Header().Set("X-Request-ID", id)

		// store it the request context so handlers can access it later
		ctx := context.WithValue(r.Context(), RequestIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}