package post

import (
	"fmt"
	"github.com/larien/clean-architecture/post"
)

func main(){
	fmt.Println("Hello, Lauren!")

	repository := post.NewRepository()

	controller := post.NewController(repository)

	_ := post.Handler(controller)
}