package infrastructure

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/keyem4251/go-todo-app/pkg/domain/repository"
	"github.com/keyem4251/go-todo-app/pkg/domain/model"
)

type ItemRepository struct {
	Conn *gorm.DB
}

func NewItemRepository() repository.ItemRepository {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &ItemRepository {
		Conn: db,
	}
}

func (ir *ItemRepository) Create(item *model.Item) (*model.Item, error) {
	if err := ir.Conn.Create(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (ir *ItemRepository) FindById(id int) (*model.Item, error) {
	item := &model.Item{
		Id: id,
	}

	if err := ir.Conn.First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (ir *ItemRepository) Update(item *model.Item) error {
	if err := ir.Conn.Model(&item).Updates(&item).Error; err != nil {
		return err
	}

	return nil
}

func (ir *ItemRepository) Delete(item *model.Item) error {
	if err := ir.Conn.Delete(&item).Error; err != nil {
		return err
	}
	return nil
}
