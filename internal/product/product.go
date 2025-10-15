package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      pq.StringArray `gorm:"type:text[]"`
}

func NewProduct(name string, des string, images pq.StringArray) *Product {
	return &Product{
		Name:        name,
		Description: des,
		Images:      images,
	}
}
