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
	AddressType   string
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
	AddressType   string
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
	AddressType   string
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

type PageResultDto struct {
	Items      []AddressResponseDto
	TotalItems int
}
