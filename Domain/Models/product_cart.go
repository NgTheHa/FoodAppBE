package domain

type ProductCart struct {
	ID        int     `json:"id"`
	Quantity  int     `json:"quantity"`
	ProductID int     `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	CartID    int     `json:"cart_id"`
	Cart      Cart    `json:"cart" gorm:"foreignKey:CartID"`
	IsDeleted bool    `json:"is_deleted"`
}
