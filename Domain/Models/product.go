package domain

import "time"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" gorm:"size:250;not null" binding:"required"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	ActualPrice float64 `json:"actual_price" gorm:"type:decimal(10,2);not null"`
	Description string  `json:"description" gorm:"size:500;not null" binding:"required"`
	CategoryID  int     `json:"category_id"`
	//Category          Category           `json:"category" gorm:"foreignKey:CategoryID"`
	OrderDetails []OrderDetail `json:"order_details" gorm:"foreignKey:ProductID"`
	//ProductImages     []ProductImage     `json:"product_images" gorm:"foreignKey:ProductID"`
	Orders []Order `json:"orders" gorm:"many2many:order_products"`
	//ProductPromotions []ProductPromotion `json:"product_promotions" gorm:"foreignKey:ProductID"`
	//Carts             []Cart             `json:"carts" gorm:"many2many:cart_products"`
	IsActived bool      `json:"is_actived"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}
