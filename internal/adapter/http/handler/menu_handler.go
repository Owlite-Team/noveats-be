package handler

import (
	"net/http"
	"noveats-be/internal/adapter/http/dto"
	"noveats-be/internal/usecase/menu"
	"noveats-be/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuHandler struct {
	createMenuUC *menu.CreateMenuUseCase
	getMenuUC    *menu.GetMenuUseCase
	updateMenuUC *menu.UpdateMenuUseCase
	deleteMenuUC *menu.DeleteMenuUseCase
	logger       *zap.Logger
}

func NewMenuHandler(
	createMenuUC *menu.CreateMenuUseCase,
	getMenuUC *menu.GetMenuUseCase,
	updateMenuUC *menu.UpdateMenuUseCase,
	deleteMenuUC *menu.DeleteMenuUseCase,
	logger *zap.Logger,
) *MenuHandler {
	return &MenuHandler{
		createMenuUC: createMenuUC,
		getMenuUC:    getMenuUC,
		updateMenuUC: updateMenuUC,
		deleteMenuUC: deleteMenuUC,
		logger:       logger,
	}
}

// func (h *MenuHandler) CreateMenu(c *gin.Context) {
//
// }

func (h *MenuHandler) GetAllMenu(c *gin.Context) {
	m, err := h.getMenuUC.ExecuteAll()
	if err != nil {
		h.logger.Error("Failed to get all menu", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, err)
	}

	res := dto.ToMenuResponseList(m)
	response.Success(c, http.StatusOK, "Menu retrieved successfully", res)
}

func (h *MenuHandler) GetMenu(c *gin.Context) {
	id := c.Param("id")

	m, err := h.getMenuUC.Execute(id)
	if err != nil {
		h.logger.Error("Failed to get menu", zap.String("id", id), zap.Error(err))
		response.Error(c, http.StatusNotFound, err)

		return
	}

	res := dto.ToMenuResponse(m)
	response.Success(c, http.StatusOK, "Menu retrieve successfully", res)
}

// func (h *MenuHandler) UpdateMenu(c *gin.Context) {
//
// }

// func (h *MenuHandler) DeleteMenu(c *gin.Context) {
//
// }
