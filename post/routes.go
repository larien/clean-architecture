package post

import (
	"larien/clean-architecture/helper"
	"larien/clean-architecture/helper/request"
	"net/http"
)

// NewRoutes creates a router for Post and sets the endpoints
func NewRoutes(c Controller) helper.Router {
	r := router.NewRouter()

	r.Post("/posts", create(c)) // POST /posts/

	return r
}

func create(c Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post Post
		if err := request.Read(r, post); err != nil {
			request.Error(w, http.StatusBadRequest, err)
			return
		}

		// if err := c.Create(); err != nil {
		// 	log.Printf("an error ocurred when creating a post: %v", err)
		// 	helper.Error(w, http.StatusBadRequest, err)
		// 	return
		// }

		request.Success(w, http.StatusOK, "Post created successfully")
	}
}
