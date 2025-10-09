package menu

import (
	"errors"
	"noveats-be/internal/domain/repository"
)

type DeleteMenuUseCase struct {
	menuRepository repository.MenuRepository
}

func NewDeleteMenuUseCase(repository repository.
	MenuRepository) *DeleteMenuUseCase {
	return &DeleteMenuUseCase{
		menuRepository: repository,
	}
}

func (uc *DeleteMenuUseCase) Execute(id string) error {
	_, err := uc.menuRepository.FindById(id)
	if err != nil {
		return errors.New("menu not found")
	}

	return uc.menuRepository.Delete(id)
}
