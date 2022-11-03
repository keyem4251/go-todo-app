package service

import (
	"github.com/keyem4251/go-todo-app/pkg/domain/repository"
)

type ItemService struct {
	QueryRepository repository.ItemQueryRepository
}

func (is *ItemService) Exists(id int) bool {
	return is.QueryRepository.Exists(id)
}

func NewItemService(queryRepository repository.ItemQueryRepository) ItemService {
	return ItemService{
		QueryRepository: queryRepository,
	}
}
