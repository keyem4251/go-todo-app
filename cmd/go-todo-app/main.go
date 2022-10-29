package main

import (
	"fmt"
	"net/http"

	"github.com/keyem4251/go-todo-app/pkg/domain/model"
	"github.com/keyem4251/go-todo-app/pkg/infrastructure"
	"github.com/keyem4251/go-todo-app/pkg/adapter/controller"
)

func main() {
	item, _ := model.NewItem(1, "Hello", "World")
	fmt.Printf("Item model: %s, %s\n", item.Title, item.Content)

	infrastructure.Init()
	
	fmt.Println("start server")
	itemController := controller.ItemController{}

	http.HandleFunc("/", itemController.Get)
	fmt.Println(http.ListenAndServe(":8080", nil))
}
