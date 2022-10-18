package infrastructure

import (
	"gorm.io/gorm"

	"github.com/keyem4251/go-todo-app/pkg/domain/repository"
	"github.com/keyem4251/go-todo-app/pkg/domain/model"
)

type ItemQueryRepositoryImpl struct {
	Conn *gorm.DB
}

func NewItemQueryRepository() repository.ItemQueryRepository {
	db := GetDB()

	return &ItemQueryRepositoryImpl {
		Conn: db,
	}
}

func (iqr *ItemQueryRepositoryImpl) FindById(id int) (*model.Item, error) {
	item := &model.Item{
		Id: id,
	}

	if err := iqr.Conn.First(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}


type ItemCommandRepositoryImpl struct {
	Conn *gorm.DB
}

func NewItemCommandRepository() repository.ItemCommandRepository {
	db := GetDB()

	return &ItemCommandRepositoryImpl {
		Conn: db,
	}
}

func (icr *ItemCommandRepositoryImpl) Create(item *model.Item) (*model.Item, error) {
	if err := icr.Conn.Create(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}


func (icr *ItemCommandRepositoryImpl) Update(item *model.Item) error {
	if err := icr.Conn.Model(&item).Updates(&item).Error; err != nil {
		return err
	}

	return nil
}

func (icr *ItemCommandRepositoryImpl) Delete(item *model.Item) error {
	if err := icr.Conn.Delete(&item).Error; err != nil {
		return err
	}
	return nil
}
