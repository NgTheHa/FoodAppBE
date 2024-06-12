package domain

import (
	"time"
)

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"products" gorm:"foreignKey:CategoryID"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int       `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   int       `json:"updated_by"`
}
