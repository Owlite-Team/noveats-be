package entity

import "time"

type Category string

const (
	RICE     Category = "rice"
	NOODLES  Category = "noodles"
	BEVERAGE Category = "beverage"
	UNKNOWN  Category = "unknown"
)

var categoryMap = map[string]Category{
	"rice":     RICE,
	"noodles":  NOODLES,
	"beverage": BEVERAGE,
}

type MenuItem struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    Category  `json:"category"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ParseCategoryString(s string) Category {
	cat := categoryMap[s]
	return cat
}
