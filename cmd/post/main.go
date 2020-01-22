package main

import (
	"fmt"
	"github.com/larien/clean-architecture/post"
	"net/http"
)

func main() {
	fmt.Println("Hello, Lauren!")

	repository := post.NewRepository()

	controller := post.NewController(repository)

	routes := post.NewRoutes(controller)

	panic(http.ListenAndServe(":8080", routes))
}
