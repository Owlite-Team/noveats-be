package dto

import (
	"noveats-be/internal/domain/entity"
	"time"
)

type MenuResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"image_url"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func ToMenuResponse(menu *entity.MenuItem) *MenuResponse {
	return &MenuResponse{
		ID:          menu.ID,
		Name:        menu.Name,
		Description: menu.Description,
		Price:       menu.Price,
		Category:    string(menu.Category),
		ImageURL:    menu.ImageURL,
		CreatedAt:   menu.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   menu.CreatedAt.Format(time.RFC3339),
	}
}

func ToMenuResponseList(items []*entity.MenuItem) []*MenuResponse {
	menuList := make([]*MenuResponse, len(items))
	for i, items := range items {
		menuList[i] = ToMenuResponse(items)
	}
	return menuList
}
