package main

import (
	"fmt"
	"github.com/larien/clean-architecture/article"
	"net/http"
)

func main() {
	fmt.Println("Hello, Lauren!")

	repository := article.NewRepository()

	controller := article.NewController(repository)

	routes := article.NewRoutes(controller)

	panic(http.ListenAndServe(":8080", routes))
}
