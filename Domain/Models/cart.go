package domain

import (
	"time"
)

type Cart struct {
	ID           int           `json:"id"`
	UserID       int           `json:"user_id"`
	ProductCarts []ProductCart `json:"product_carts" gorm:"foreignKey:CartID"`
	Products     []Product     `json:"products" gorm:"many2many:cart_products"`
	CreatedAt    time.Time     `json:"created_at"`
	CreatedBy    int           `json:"created_by"`
	IsDeleted    bool          `json:"is_deleted"`
}
