package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithError function to send an error response as JSON
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON function to send a JSON response
// The interfac{} type can hold values of any type, which is flexible
// for functions that need to handle different kinds of data
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	if response, err := json.Marshal(payload); err != nil {
		log.Printf("Failed to marshall payload %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(response)
	}
}
