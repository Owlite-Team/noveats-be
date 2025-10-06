package user

import (
	"errors"
	"noveats-be/internal/domain/entity"
	"noveats-be/internal/domain/repository"
	"time"

	"github.com/google/uuid"
)

type CreateUserUseCase struct {
	userRepository repository.UserRepository
}

func NewCreateUserUseCase(repository repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: repository,
	}
}

// Execute retrieves a user by ID and returns the domain entity
func (uc *CreateUserUseCase) Execute(email, name, password string) (*entity.User, error) {
	existingUser, _ := uc.userRepository.FindByEmail(email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Create new user entity
	user := &entity.User{
		ID:        uuid.New().String(),
		Email:     email,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Hash password
	if err := user.HashPassword(password); err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Save to database
	if err := uc.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// ExecuteAll retrieves all users and returns the domain entity
func (uc *CreateUserUseCase) ExecuteAll() ([]*entity.User, error) {
	users, err := uc.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}
