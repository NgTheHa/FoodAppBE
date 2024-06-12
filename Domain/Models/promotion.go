package domain

import (
	"time"
)

type Promotion struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	PromotionCode     string             `json:"promotion_code"`
	Description       string             `json:"description"`
	DiscountPercent   float64            `json:"discount_percent"`
	StartTime         time.Time          `json:"start_time"`
	EndTime           time.Time          `json:"end_time"`
	IsActive          bool               `json:"is_active"`
	ProductPromotions []ProductPromotion `json:"product_promotions" gorm:"foreignKey:PromotionID"`
	CreatedAt         time.Time          `json:"created_at"`
	CreatedBy         int                `json:"created_by"`
	UpdatedAt         time.Time          `json:"updated_at"`
	UpdatedBy         int                `json:"updated_by"`
	IsDeleted         bool               `json:"is_deleted"`
}
