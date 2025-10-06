package user

import (
	"errors"
	"noveats-be/internal/domain/entity"
	"noveats-be/internal/domain/repository"
	"time"
)

type UpdateUserUseCase struct {
	userRepository repository.UserRepository
}

func NewUpdateUserUseCase(repository repository.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepository: repository,
	}
}

// Execute updates a user's information
func (uc *UpdateUserUseCase) Execute(id, name string) (*entity.User, error) {
	user, err := uc.userRepository.FindById(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	user.Name = name
	user.UpdatedAt = time.Now()

	if err := uc.userRepository.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}
