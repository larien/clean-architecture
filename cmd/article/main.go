package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/larien/clean-architecture/article"
	"github.com/larien/clean-architecture/helper/database"
)

func main() {
	fmt.Println("Hello, Lauren!")

	db, err := database.New("localhost", "larien", "clean_architecture", "")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	db.DropTableIfExists(&article.Article{}) // remove

	repository := article.NewRepository(db)

	controller := article.NewController(repository)

	routes := article.NewRoutes(controller)

	panic(http.ListenAndServe(":8080", routes))
}
