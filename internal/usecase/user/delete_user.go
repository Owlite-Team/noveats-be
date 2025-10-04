package user

import (
	"errors"
	"noveats-be/internal/domain/repository"
)

type DeleteUserUseCase struct {
	userRepository repository.UserRepository
}

func NewDeleteUserUseCase(repository repository.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepository: repository,
	}
}

func (uc *DeleteUserUseCase) Execute(id string) error {
	_, err := uc.userRepository.FindById(id)
	if err != nil {
		return errors.New("user not found")
	}

	return uc.userRepository.Delete(id)
}
