package menu

import (
	"noveats-be/internal/domain/entity"
	"noveats-be/internal/domain/repository"
	"time"

	"github.com/google/uuid"
)

type CreateMenuUseCase struct {
	menuRepository repository.MenuRepository
}

func NewCreateMenuUseCase(repository repository.MenuRepository) *CreateMenuUseCase {
	return &CreateMenuUseCase{
		menuRepository: repository,
	}
}

func (uc *CreateMenuUseCase) Execute(
	name, desc string,
	price float64,
	category, imageURL string,
) (*entity.MenuItem, error) {
	menu := &entity.MenuItem{
		ID:          uuid.New().String(),
		Name:        name,
		Description: desc,
		Price:       price,
		Category:    entity.Category(category),
		ImageURL:    imageURL,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := uc.menuRepository.Create(menu); err != nil {
		return nil, err
	}

	return menu, nil
}
