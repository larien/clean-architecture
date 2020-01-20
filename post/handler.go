package post

import "net/http"

// HandlerFunc defines the expected return for an endpoint from handler layer.
type HandlerFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int)

// Handler exposes the defined feature for handler layer.
func Handler() HandlerFunc { // TODO - receive Controller as parameter
	return func(w http.ResponseWriter, r *http.Request) (interface{}, int) {
		// receives params
		// calls controller layer
		// if fails, returns 500
		// if successes, returns 200
		return "", 0
	}
}