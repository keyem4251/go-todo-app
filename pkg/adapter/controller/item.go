package controller

import (
	"fmt"
	"net/http"
)

type ItemController struct {}

func (ctrl ItemController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get request.")
	switch r.Method {
	case http.MethodGet:
		fmt.Println(r)
	case http.MethodPost:
		fmt.Println(r)
	}
}
