package order

import (
	"project/internal/product"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	PhoneId  uint              `json:"phone_id"`
	Products []product.Product `json:"products" gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
