package post

import (
	"log"
	"net/http"

	"github.com/larien/clean-architecture/helper/request"
	"github.com/larien/clean-architecture/helper/router"
)

// NewRoutes creates a router for Post and sets the endpoints
func NewRoutes(c Controller) router.Router {
	r := router.New()

	r.Post("/posts", create(c)) // POST /posts/

	return r
}

func create(c Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post *Post
		if err := request.Read(r, post); err != nil {
			log.Printf("an error occurred when parsing JSON: %v", err)
			request.Error(w, http.StatusBadRequest, err)
			return
		}

		if err := c.Create(post); err != nil {
			log.Printf("an error occurred when creating a post: %v", err)
			request.Error(w, http.StatusInternalServerError, err)
			return
		}

		request.Success(w, http.StatusOK, "Post created successfully")
	}
}
