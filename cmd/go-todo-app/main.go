package main

import (
	"fmt"
	"net/http"

	"github.com/keyem4251/go-todo-app/pkg/infrastructure"
	"github.com/keyem4251/go-todo-app/pkg/adapter/controller"
)

func main() {
	infrastructure.Init()
	
	fmt.Println("start server")
	itemController := controller.ItemController{}

	http.HandleFunc("/", itemController.Get)
	http.HandleFunc("/crate_item", itemController.Post)
	fmt.Println(http.ListenAndServe(":8080", nil))
}
