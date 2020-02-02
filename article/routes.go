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

	r.Post("/articles", create(c)) // POST /articles
	r.Get("/articles", list(c))    // GET /articles

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

// list is the handler for Article's list and handles GET /articles
func list(controller Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := controller.List()
		if err != nil {
			log.Printf("an error occurred when listing the articles: %v", err)
			request.Error(w, http.StatusInternalServerError, err)
			return
		}

		request.Write(w, http.StatusOK, articles)
	}
}
