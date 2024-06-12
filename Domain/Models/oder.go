package domain

import (
	"time"
)

type Order struct {
	ID            int           `json:"id"`
	UserID        int           `json:"user_id"`
	User          User          `json:"user" gorm:"foreignKey:UserID"`
	Status        int           `json:"status"`
	PaymentMethod int           `json:"payment_method"`
	Province      string        `json:"province"`
	District      string        `json:"district"`
	Ward          string        `json:"ward"`
	StreetAddress string        `json:"street_address"`
	DetailAddress string        `json:"detail_address" gorm:"size:250" binding:"required"`
	Notes         string        `json:"notes" gorm:"size:250"`
	AddressType   int           `json:"address_type"`
	Products      []Product     `json:"products" gorm:"many2many:order_products"`
	OrderDetails  []OrderDetail `json:"order_details" gorm:"foreignKey:OrderID"`
	IsPaid        bool          `json:"is_paid"`
	IsDeleted     bool          `json:"is_deleted"`
	CreatedAt     time.Time     `json:"created_at"`
	CreatedBy     int           `json:"created_by"`
	UpdatedAt     time.Time     `json:"updated_at"`
	UpdatedBy     int           `json:"updated_by"`
}
