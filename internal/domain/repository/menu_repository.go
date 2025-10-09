package repository

import "noveats-be/internal/domain/entity"

type MenuRepository interface {
	Create(menu *entity.MenuItem) error
	FindById(id string) (*entity.MenuItem, error)
	FindByName(name string) (*entity.MenuItem, error)
	FindAll() ([]*entity.MenuItem, error)
	Update(menu *entity.MenuItem) error
	Delete(id string) error
}
