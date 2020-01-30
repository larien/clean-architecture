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
func create(controller Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		article := &Article{}
		if err := article.Decode(r); err != nil {
			log.Printf("an error occurred when parsing JSON: %v", err)
			request.Error(w, http.StatusBadRequest, err)
			return
		}

		if err := controller.Create(article); err != nil {
			log.Printf("an error occurred when creating a article: %v", err)
			request.Error(w, http.StatusInternalServerError, err)
			return
		}

		request.Write(w, http.StatusOK, article)
	}
}
