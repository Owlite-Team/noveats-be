package dto

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"min=8"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}
