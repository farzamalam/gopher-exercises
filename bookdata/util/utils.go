package util

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond ...
func Respond(w http.ResponseWriter, statusCode int, data map[string]interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
