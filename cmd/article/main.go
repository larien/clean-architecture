package main

import (
	"fmt"
	"net/http"

	"github.com/larien/clean-architecture/article"
)

func main() {
	fmt.Println("Hello, Lauren!")

	repository := article.NewRepository("localhost", "larien", "clean_architecture", "")

	controller := article.NewController(repository)

	routes := article.NewRoutes(controller)

	panic(http.ListenAndServe(":8080", routes))
}
