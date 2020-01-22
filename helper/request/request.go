package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Write(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func Read(r *http.Request, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(&body)
}

func Success(w http.ResponseWriter, statusCode int, message string) {
	Write(w, statusCode, struct {
		Message string `json:"message"`
	}{message})
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	Write(w, statusCode, struct {
		Error string `json:"error"`
	}{err.Error()})
}
