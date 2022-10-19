package dto

import (
	"github.com/keyem4251/go-todo-app/pkg/domain/model"
)

type Item struct {
	Id      int
	Title   string
	Content string
	Status  model.Status
}

func (i *Item) ConvertToModel() *model.Item {
	return &model.Item{
		Id: i.Id,
		Title: i.Title,
		Content: i.Content,
		Status: i.Status,
	}
}
