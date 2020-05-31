// Package utils contains handy utils functions to build json message and return json response.
package utils

import (
	"encoding/json"
	"net/http"
)

// Message is used to set the status and message in the response
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond is used to write the message and statusCode and Content-type to the response.
func Respond(w http.ResponseWriter, statusCode int, data map[string]interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
