package repository

import "github.com/keyem4251/go-todo-app/pkg/domain/model"

type ItemQueryRepository interface {
	FindById(id int) (*model.Item, error)
	// FindAll() (*model.Items, error) 複数返す実装
}

type ItemCommandRepository interface {
	Create(item *model.Item) error
	Update(item *model.Item) error
	Delete(item *model.Item) error
}

