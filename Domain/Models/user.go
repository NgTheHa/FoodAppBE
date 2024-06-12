package domain

import "time"

type User struct {
	ID                     int           `json:"id"`
	FirstName              string        `json:"first_name" gorm:"size:250"`
	LastName               string        `json:"last_name" gorm:"size:250"`
	FullName               string        `json:"full_name" gorm:"-"`
	PhoneNumber            string        `json:"phone_number" gorm:"size:20"`
	Email                  string        `json:"email" gorm:"size:250"`
	Password               string        `json:"password"`
	UserType               int           `json:"user_type"`
	RefreshToken           string        `json:"refresh_token"`
	UserAddresses          []UserAddress `json:"user_addresses" gorm:"foreignKey:UserID"`
	RefreshTokenExpiryTime time.Time     `json:"refresh_token_expiry_time"`
	Orders                 []Order       `json:"orders" gorm:"foreignKey:UserID"`
	CreatedAt              time.Time     `json:"created_at"`
	CreatedBy              int           `json:"created_by"`
}
