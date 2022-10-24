package main

import "fmt"
import "github.com/keyem4251/go-todo-app/pkg/domain/model"
import "github.com/keyem4251/go-todo-app/pkg/infrastructure"

func main() {
	item, _ := model.NewItem(1, "Hello", "World")
	fmt.Printf("Item model: %s, %s\n", item.Title, item.Content)

	infrastructure.Init()
}
