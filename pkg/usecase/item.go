package usecase

import (
	"github.com/keyem4251/go-todo-app/pkg/domain/model"
	"github.com/keyem4251/go-todo-app/pkg/domain/repository"
)

type getItemUseCase struct {
	ItemId int
	ItemRepository repository.ItemQueryRepository
}

func NewGetItemUseCase(itemId int, itemRepository repository.ItemQueryRepository) *getItemUseCase {
	return &getItemUseCase{
		ItemId: itemId,
		ItemRepository: itemRepository,
	}
}

func (ius getItemUseCase) Exec() (*model.Item, error) {
	item, err := ius.ItemRepository.FindById(ius.ItemId)
	if err != nil {
		return nil, err
	}
	return item, nil
}
