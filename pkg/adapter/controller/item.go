package controller

import (
	"fmt"
	"net/http"

	"github.com/keyem4251/go-todo-app/pkg/infrastructure/repositoryImpl"
	"github.com/keyem4251/go-todo-app/pkg/usecase"
)

type ItemController struct {}

func (ctrl ItemController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get request.")
	switch r.Method {
	case http.MethodGet:
		itemRepository := repositoryImpl.NewItemQueryRepository()
		getItemUseCase := usecase.NewGetItemUseCase(1, itemRepository)
		item, err := getItemUseCase.Exec()
		if err != nil {
			fmt.Println("Item Not Found")
		}
		w.Write([]byte(item.Title))
		w.Write([]byte(item.Content))
	}
}

func (ctrl ItemController) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post request.")
	switch r.Method {
	case http.MethodPost:
		itemRepository := repositoryImpl.NewItemCommandRepository()
		getItemUseCase := usecase.NewCreateItemUseCase("title", "content", itemRepository)
		err := getItemUseCase.Exec()
		if err != nil {
			fmt.Println("Create Item Failed")
			return
		}
		fmt.Println("Create Item Success")
	}
}
