package domain

import (
	"time"
)

type ProductImage struct {
	ID              int       `json:"id"`
	Description     string    `json:"description"`
	ProductImageUrl string    `json:"product_image_url"`
	FileSize        int64     `json:"file_size"`
	ProductID       int       `json:"product_id"`
	Product         Product   `json:"product" gorm:"foreignKey:ProductID"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       int       `json:"created_by"`
	UpdatedAt       time.Time `json:"updated_at"`
	UpdatedBy       int       `json:"updated_by"`
	IsDeleted       bool      `json:"is_deleted"`
}
