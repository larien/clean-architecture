package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	// GetParamFromURL returns the URL parameter from a http.Request object
	GetParamFromURL = chi.URLParam
)

// Router represents the web framework router
type Router chi.Router

// New creates a new router and sets basic middleware
func New() Router {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	return router
}
