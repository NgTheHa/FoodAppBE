package domain

type ProductPromotion struct {
	ID          int       `json:"id"`
	ProductID   int       `json:"product_id"`
	Product     Product   `json:"product" gorm:"foreignKey:ProductID"`
	PromotionID int       `json:"promotion_id"`
	Promotion   Promotion `json:"promotion" gorm:"foreignKey:PromotionID"`
	IsActive    bool      `json:"is_active"`
}
