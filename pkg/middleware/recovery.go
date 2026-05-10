package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

//Recovery catches panics and prevents the entire server from crashing
//why ? imagine a user sends unexpected data that cause a nil pointer dereference
// without recovery, The server crashes. All users are affected
// with recovery: That one request gets a 500 error, The server keeps runnning
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// defer runs after the handler finishes (or panics)

		defer func(){
			if err := recover(); err != nil {
				log.Printf("PANIC: %v\n%s", err, debug.Stack())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError	)
			}
		}()
		next.ServeHTTP(w,r)
	})
}