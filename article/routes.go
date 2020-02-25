package article

import (
	"fmt"
	"log"
	"net/http"

	"github.com/larien/clean-architecture/helper/request"
	"github.com/larien/clean-architecture/helper/router"
)

// NewRoutes creates a router for Article and sets the endpoints
func NewRoutes(c Controller) router.Router {
	r := router.New()

	r.Post("/articles", create(c))                    // POST /articles
	r.Get("/articles", list(c))                       // GET /articles
	r.Get("/articles/{article_id}/detail", detail(c)) // GET /articles/{article_id}/detail

	// r.Route("/articles", func(r chi.Router) {
	// 	r.Get("/{article_id}/detail", detail(c))
	// })

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
			log.Printf("an error occurred when creating the article: %v", err)
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

		if articles == nil {
			request.Write(w, http.StatusNotFound, nil)
			return
		}

		request.Write(w, http.StatusOK, articles)
	}
}

// detail is the handler for Article's detail and handles GET /articles/{article_id}/detail
func detail(controller Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("ID: %+v\n", r.URL.Query())
		articleID, err := request.ParseID(router.GetParamFromURL(r, "article_id"))
		if err != nil {
			log.Printf("couldn't parse article ID: %v", err)
			request.Error(w, http.StatusBadRequest, err)
			return
		}

		article, err := controller.Detail(articleID)
		if err != nil {
			log.Printf("an error occurred when detailing article %d: %v", articleID, err)
			request.Error(w, http.StatusInternalServerError, err)
			return
		}

		if article == nil {
			request.Write(w, http.StatusNotFound, nil)
			return
		}

		request.Write(w, http.StatusOK, article)
	}
}
