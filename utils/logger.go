package utils

import (
	"log"
	"net/http"
)

// middleware function that takes in a http.Handler function, logs what occured, then
// actually execute the http.Handler
func DisplayLog(incomingf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, r.Method)
		incomingf(w, r)
	}
}
