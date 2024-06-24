package DTO

import (
	"time"
)

// CartResponseDto represents the response structure for a cart
type CartResponseDto struct {
	CartID     int                  `json:"cart_id"`
	TotalPrice float64              `json:"total_price"`
	Products   []InfoProductCartDto `json:"products"`
}

// CreateCartDto represents the structure for creating a new cart
type CreateCartDto struct {
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	IsDeleted bool      `json:"is_deleted"`
}

// InfoProductCartDto represents the structure for a product in a cart
type InfoProductCartDto struct {
	ProductCartID   int     `json:"product_cart_id"`
	ProductID       int     `json:"product_id"`
	Quantity        int     `json:"quantity"`
	Name            string  `json:"name"`
	ActualPrice     float64 `json:"actual_price"`
	CategoryID      int     `json:"category_id"`
	Description     string  `json:"description"`
	ProductImageUrl string  `json:"product_image_url"`
	Price           float64 `json:"price"`
	IsActive        bool    `json:"is_active"`
}

// UpdateCartDto represents the structure for updating the cart
type UpdateCartDto struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
type CreateCategoryDto struct {
	Name        string `json:"name" binding:"required,max=250"`
	Description string `json:"description"`
}

type UpdateCategoryDto struct {
	CreateCategoryDto
	Id int `json:"id"`
}

type CategoryResponseDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryFilterDto struct {
	Name      string `json:"name"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

type PageResultDto struct {
	Items     []CategoryResponseDto `json:"items"`
	TotalItem int                   `json:"totalItem"`
}

type DashboardResponseDto struct {
	Id           int                  `json:"id"`
	UpdatedDate  time.Time            `json:"updatedDate"`
	OrderRevenue float64              `json:"orderRevenue"`
	Products     []InfoProductCartDto `json:"products"`
}

type DashBoardInfoProductCartDto struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	ActualPrice     float64 `json:"actualPrice"`
	CategoryId      int     `json:"categoryId"`
	Description     string  `json:"description"`
	ProductImageUrl string  `json:"productImageUrl"`
	Quantity        int     `json:"quantity"`
	CreateBy        int     `json:"createBy"`
	Price           float64 `json:"price"`
	IsActive        bool    `json:"isActive"`
}

type DashboardPageResultDto struct {
	Item         []DashboardResponseDto `json:"item"`
	TotalItem    int                    `json:"totalItem"`
	TotalRevenue float64                `json:"totalRevenue"`
}

type DashboardFilterDto struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	PageIndex int       `json:"pageIndex"`
	PageSize  int       `json:"pageSize"`
}

type DashboardFilterByDayDto struct {
	Year      int `json:"year"`
	Month     int `json:"month"`
	Day       int `json:"day"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type DashboardFilterByMonthDto struct {
	Year      int `json:"year"`
	Month     int `json:"month"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type DashboardResponseTopProductsDto struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	Price           float64   `json:"price"`
	Description     string    `json:"description"`
	CategoryId      int       `json:"categoryId"`
	IsActived       bool      `json:"isActived"`
	CreatedAt       time.Time `json:"createdAt"`
	ProductImageUrl string    `json:"productImageUrl"`
	TotalQuantity   int       `json:"totalQuantity"`
	TotalRevenue    float64   `json:"totalRevenue"`
}

type DashboardProductsFilterDto struct {
	Amount    int `json:"amount"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}
