package menu

import (
	"errors"
	"noveats-be/internal/domain/entity"
	"noveats-be/internal/domain/repository"
	"time"
)

type UpdateMenuUseCase struct {
	menuRepository repository.MenuRepository
}

func NewUpdateMenuUseCase(repository repository.MenuRepository) *UpdateMenuUseCase {
	return &UpdateMenuUseCase{
		menuRepository: repository,
	}
}

func (uc *UpdateMenuUseCase) Execute(
	id string, name, desc string,
	price float64,
	category, imageURL string) (*entity.MenuItem, error) {
	menu, err := uc.menuRepository.FindById(id)
	if err != nil {
		return nil, errors.New("menu not found")
	}

	menu.Name = name
	menu.Description = desc
	menu.Price = price
	menu.Category = entity.ParseCategoryString(category)
	menu.ImageURL = imageURL
	menu.UpdatedAt = time.Now()

	if err := uc.menuRepository.Update(menu); err != nil {
		return nil, err
	}

	return menu, nil
}
