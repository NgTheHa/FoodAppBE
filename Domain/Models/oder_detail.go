package domain

type OrderDetail struct {
	ID        int     `json:"id"`
	Quantity  int     `json:"quantity"`
	OrderID   int     `json:"order_id"`
	Order     Order   `json:"order" gorm:"foreignKey:OrderID"`
	ProductID int     `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	IsDeleted bool    `json:"is_deleted"`
}
