package article

import (
	"log"
	"net/http"

	"github.com/larien/clean-architecture/helper/request"
	"github.com/larien/clean-architecture/helper/router"
)

// NewRoutes creates a router for Article and sets the endpoints
func NewRoutes(c Controller) router.Router {
	r := router.New()

	r.Post("/articles", create(c)) // POST /articles/

	return r
}

// create is the handler for Article creation and handles POST /articles
func create(c Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var article *Article
		if err := request.Read(r, article); err != nil {
			log.Printf("an error occurred when parsing JSON: %v", err)
			request.Error(w, http.StatusBadRequest, err)
			return
		}

		if err := c.Create(article); err != nil {
			log.Printf("an error occurred when creating a article: %v", err)
			request.Error(w, http.StatusInternalServerError, err)
			return
		}

		request.Success(w, http.StatusOK, "Article created successfully")
	}
}
