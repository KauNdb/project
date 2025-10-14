package product

import "github.com/lib/pq"

type ProductRequest struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Images      pq.StringArray `gorm:"type:text[]"`
}

type ProductResponse struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Images      pq.StringArray `gorm:"type:text[]"`
}
