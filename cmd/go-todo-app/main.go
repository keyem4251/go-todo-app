package main

import (
	"fmt"
	"net/http"

	"github.com/keyem4251/go-todo-app/pkg/domain/model"
	"github.com/keyem4251/go-todo-app/pkg/infrastructure"
)

func main() {
	item, _ := model.NewItem(1, "Hello", "World")
	fmt.Printf("Item model: %s, %s\n", item.Title, item.Content)

	infrastructure.Init()
	
	fmt.Println("start server")
	http.HandleFunc("/", sample_handler)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func sample_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get request.")
	switch r.Method {
	case http.MethodGet:
		fmt.Println(r)
	case http.MethodPost:
		fmt.Println(r)
	}
}
