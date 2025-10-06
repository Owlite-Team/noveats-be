package user

import (
	"errors"
	"noveats-be/internal/domain/entity"
	"noveats-be/internal/domain/repository"
)

type GetUserUseCase struct {
	userRepository repository.UserRepository
}

func NewGetUserUseCase(repository repository.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{
		userRepository: repository,
	}
}

func (uc *GetUserUseCase) Execute(id string) (*entity.User, error) {
	user, err := uc.userRepository.FindById(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (uc *GetUserUseCase) ExecuteAll() ([]*entity.User, error) {
	users, err := uc.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}
