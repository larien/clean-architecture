package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Read decodes the received body in a request
func Read(r *http.Request, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(&body)
}

// Write encodes the response for a request with its status code
func Write(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// Success returns a success message for a request
func Success(w http.ResponseWriter, statusCode int, message string) {
	Write(w, statusCode, struct {
		Message string `json:"message"`
	}{message})
}

// Error returns the message with reason for failure for a request
func Error(w http.ResponseWriter, statusCode int, err error) {
	Write(w, statusCode, struct {
		Error string `json:"error"`
	}{err.Error()})
}
