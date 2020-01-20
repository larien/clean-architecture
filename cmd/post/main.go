package post

import (
	"fmt"
	"github.com/larien/clean-architecture/post"
)

func main(){
	fmt.Println("Hello, Lauren!")

	repository := NewRepository()

	controller := NewController(repository)

	_ := Handler(controller)
}