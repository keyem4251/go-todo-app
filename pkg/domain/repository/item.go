package repository

import "github.com/keyem4251/go-todo-app/pkg/domain/model"

type ItemRepository interface {
	Create(item *model.Item) (*model.Item, error)
	FindById(id int) (*model.Item, error)
	Update(item *model.Item) error
	Delete(item *model.Item) error
}
