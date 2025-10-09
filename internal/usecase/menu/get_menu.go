package menu

import (
	"errors"
	"noveats-be/internal/domain/entity"
	"noveats-be/internal/domain/repository"
)

type GetMenuUseCase struct {
	menuRepository repository.MenuRepository
}

func NewGetMenuUseCase(repository repository.MenuRepository) *GetMenuUseCase {
	return &GetMenuUseCase{
		menuRepository: repository,
	}
}

func (uc *GetMenuUseCase) Execute(id string) (*entity.MenuItem, error) {
	item, err := uc.menuRepository.FindById(id)
	if err != nil {
		return nil, errors.New("menu not found")
	}

	return item, nil
}

func (uc *GetMenuUseCase) ExecuteAll() ([]*entity.MenuItem, error) {
	itemList, err := uc.menuRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return itemList, nil
}
