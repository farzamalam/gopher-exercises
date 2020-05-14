package utils

import (
	"encoding/json"
	"net/http"
)

// Message is a utility function return the status and message of response.
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond is used to write the REST response.
func Respond(w http.ResponseWriter, statusCode int, data map[string]interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}
