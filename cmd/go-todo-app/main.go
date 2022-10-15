package main

import "fmt"
import "github.com/keyem4251/go-todo-app/pkg/domain/model"

func main() {
	item, _ := model.NewItem("Hello", "World")
	fmt.Printf("Item model: %s, %s\n", item.Title, item.Content)
}
