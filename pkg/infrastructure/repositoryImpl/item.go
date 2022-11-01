package repositoryImpl

import (
	"gorm.io/gorm"

	"github.com/keyem4251/go-todo-app/pkg/domain/model"
	"github.com/keyem4251/go-todo-app/pkg/domain/repository"
	"github.com/keyem4251/go-todo-app/pkg/infrastructure"
	"github.com/keyem4251/go-todo-app/pkg/infrastructure/dto"
)

type ItemQueryRepositoryImpl struct {
	Conn *gorm.DB
}

func NewItemQueryRepository() repository.ItemQueryRepository {
	db := infrastructure.GetDB()

	return &ItemQueryRepositoryImpl {
		Conn: db,
	}
}

func (iqr *ItemQueryRepositoryImpl) FindById(id int) (*model.Item, error) {
	itemDao := dto.Item{}

	if err := iqr.Conn.Find(&itemDao, dto.Item{Id: id}).Error; err != nil {
		return nil, err
	}

	return itemDao.ConvertToModel(), nil
}


type ItemCommandRepositoryImpl struct {
	Conn *gorm.DB
}

func NewItemCommandRepository() repository.ItemCommandRepository {
	db := infrastructure.GetDB()

	return &ItemCommandRepositoryImpl {
		Conn: db,
	}
}

func (icr *ItemCommandRepositoryImpl) Create(item *model.Item) error {
	itemDao := dto.Item{
		Id: item.Id,
		Title: item.Title,
		Content: item.Content,
		Status: item.Status,
	}
	
	if err := icr.Conn.Create(&itemDao).Error; err != nil {
		return err
	}

	return nil
}


func (icr *ItemCommandRepositoryImpl) Update(item *model.Item) error {
	itemDao := dto.Item{}

	if err := icr.Conn.First(&itemDao, dto.Item{Id: item.Id}).Error; err != nil {
		return err
	}

	itemDao.Content = item.Content
	itemDao.Title = item.Title
	itemDao.Status = item.Status

	icr.Conn.Save(&itemDao)

	return nil
}

func (icr *ItemCommandRepositoryImpl) Delete(item *model.Item) error {
	itemDao := dto.Item{}

	if err := icr.Conn.First(&itemDao, dto.Item{Id: item.Id}).Error; err != nil {
		return err
	}

	icr.Conn.Delete(&itemDao)

	return nil
}
