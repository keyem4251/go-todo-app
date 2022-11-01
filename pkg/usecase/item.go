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

type createItemUseCase struct {
	Title string
	Content string
	ItemRepository repository.ItemCommandRepository
}

func NewCreateItemUseCase(title string, content string, itemRepository repository.ItemCommandRepository) *createItemUseCase {
	return &createItemUseCase{
		Title: title,
		Content: content,
		ItemRepository: itemRepository,
	}
}

func (ius createItemUseCase) Exec() error {
	item := model.NewItem(1, ius.Title, ius.Content)
	err := ius.ItemRepository.Create(item)
	if err != nil {
		return err
	}
	return nil
}
