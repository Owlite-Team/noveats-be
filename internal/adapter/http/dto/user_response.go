package dto

import (
	"noveats-be/internal/domain/entity"
	"time"
)

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func ToUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}

func ToUserResponseList(users []*entity.User) []*UserResponse {
	userList := make([]*UserResponse, len(users))
	for i, user := range users {
		userList[i] = ToUserResponse(user)
	}
	return userList
}
