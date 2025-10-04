package handler

import (
	"net/http"
	"noveats-be/internal/adapter/http/dto"
	"noveats-be/internal/usecase/user"
	"noveats-be/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	createUserUC *user.CreateUserUseCase
	getUserUC    *user.GetUserUseCase
	updateUserUC *user.UpdateUserUseCase
	deleteUserUC *user.DeleteUserUseCase
	logger       *zap.Logger
}

func NewUserHandler(
	createUserUC *user.CreateUserUseCase,
	getUserUC *user.GetUserUseCase,
	updateUserUC *user.UpdateUserUseCase,
	deleteUserUC *user.DeleteUserUseCase,
	logger *zap.Logger,
) *UserHandler {
	return &UserHandler{
		createUserUC: createUserUC,
		getUserUC:    getUserUC,
		updateUserUC: updateUserUC,
		deleteUserUC: deleteUserUC,
		logger:       logger,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body", zap.Error(err))
		response.Error(c, http.StatusBadRequest, err)

		return
	}

	u, err := h.createUserUC.Execute(req.Name, req.Email, req.Password)
	if err != nil {
		h.logger.Error("Failed to create u", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err)

		return
	}

	res := dto.ToUserResponse(u)
	response.Success(c, http.StatusCreated, "User created successfully", res)
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	u, err := h.getUserUC.ExecuteAll()
	if err != nil {
		h.logger.Error("Failed to get all users", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err)
	}

	res := dto.ToUserResponseList(u)
	response.Success(c, http.StatusOK, "Users retrieve successfully", res)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	u, err := h.getUserUC.Execute(id)
	if err != nil {
		h.logger.Error("Failed to get u", zap.String("id", id), zap.Error(err))
		response.Error(c, http.StatusNotFound, err)

		return
	}

	res := dto.ToUserResponse(u)
	response.Success(c, http.StatusOK, "User retrieve successfully", res)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body", zap.Error(err))
		response.Error(c, http.StatusBadRequest, err)

		return
	}

	u, err := h.updateUserUC.Execute(id, req.Name)
	if err != nil {
		h.logger.Error("Failed to update u", zap.String("id", id), zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err)

		return
	}

	res := dto.ToUserResponse(u)
	response.Success(c, http.StatusOK, "User updated successfully", res)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := h.deleteUserUC.Execute(id); err != nil {
		h.logger.Error("Failed to get user", zap.String("id", id), zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err)

		return
	}

	response.Success(c, http.StatusOK, "User deleted successfully", nil)
}
