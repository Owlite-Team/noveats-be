package repository

import "noveats-be/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindById(id string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	Update(user *entity.User) error
	Delete(id string) error
}
