package Controllers

import (
	"log"
	"net/http"
	"time"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got request with path %s", r.URL.Path)
		log.Printf("Method: %s", r.Method)
		log.Println("Time: ", time.Now().Format(time.ANSIC))

		next.ServeHTTP(w, r)
	})
}
