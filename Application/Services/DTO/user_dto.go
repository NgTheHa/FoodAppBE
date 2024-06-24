package DTO

import "time"

type UpdateUserDto struct {
	ID          int
	FirstName   string
	LastName    string
	PhoneNumber string
}

type UserResponseDto struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}

type CreateAddressDto struct {
	UserID        int
	AddressType   int
	Province      string
	District      string
	Ward          string
	StreetAddress string
	DetailAddress string
	Notes         string
}

type UpdateAddressDto struct {
	ID            int
	UserID        int
	AddressType   int
	Province      string
	District      string
	Ward          string
	StreetAddress string
	DetailAddress string
	Notes         string
}

type AddressResponseDto struct {
	ID            int
	Province      string
	District      string
	Ward          string
	StreetAddress string
	DetailAddress string
	Notes         string
	AddressType   int
	CreatedAt     time.Time
	CreatedBy     int
	UpdatedAt     time.Time
	UpdateBy      int
}

type AddressFilterDto struct {
	UserID    int
	PageIndex int
	PageSize  int
}

type UserPageResultDto struct {
	Items      []AddressResponseDto
	TotalItems int
}
