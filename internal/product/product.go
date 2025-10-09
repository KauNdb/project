package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
}

func NewProduct(name string, des string) *Product {
	return &Product{
		Name:        name,
		Description: des,
	}
}
