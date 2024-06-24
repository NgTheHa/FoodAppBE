package Services

import (
	"errors"
	"go/foodappbe/Application/Services/DTO"
	domain "go/foodappbe/Domain/Models"
	"gorm.io/gorm"
	"time"
)

type CartService interface {
	AddProductToCart(productId int) (string, error)
	RemoveProductFromCart(productId int) error
	GetCartByUserId() (*DTO.CartResponseDto, error)
	UpdateQuantity(input DTO.UpdateCartDto) (string, error)
}

type cartServiceImpl struct {
	db *gorm.DB
}

func NewCartService(db *gorm.DB) CartService {
	return &cartServiceImpl{db: db}
}

func (s *cartServiceImpl) AddProductToCart(productId int) (string, error) {
	var currentUserId int = getCurrentUserId()
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	cart := &domain.Cart{}
	if err := s.db.First(&cart, "user_id = ?", currentUserId).Error; err != nil {
		cart = &domain.Cart{
			UserID:    currentUserId,
			CreatedBy: currentUserId,
			CreatedAt: time.Now(),
			IsDeleted: false,
		}
		if err := s.db.Create(&cart).Error; err != nil {
			tx.Rollback()
			return "", err
		}
	}

	productCart := &domain.ProductCart{}
	if err := s.db.First(&productCart, "cart_id = ? AND product_id = ? AND is_deleted = ?", cart.ID, productId, false).Error; err != nil {
		productCart = &domain.ProductCart{
			ProductID: productId,
			Quantity:  1,
			CartID:    cart.ID,
			IsDeleted: false,
		}
		if err := s.db.Create(&productCart).Error; err != nil {
			tx.Rollback()
			return "", err
		}
	} else {
		productCart.Quantity++
		if err := s.db.Save(&productCart).Error; err != nil {
			tx.Rollback()
			return "", err
		}
	}

	tx.Commit()
	return "Thêm sản phẩm vào giỏ hàng thành công", nil
}

func (s *cartServiceImpl) RemoveProductFromCart(productId int) error {
	productCart := &domain.ProductCart{}
	if err := s.db.First(&productCart, "product_id = ? AND is_deleted = ?", productId, false).Error; err != nil {
		return errors.New("Không tìm thấy sản phẩm trong giỏ hàng")
	}

	productCart.IsDeleted = true
	return s.db.Save(&productCart).Error
}

func (s *cartServiceImpl) GetCartByUserId() (*DTO.CartResponseDto, error) {
	var currentUserId int = getCurrentUserId()
	var shoppingCart DTO.CartResponseDto

	err := s.db.Table("carts").
		Select("carts.id AS cart_id, SUM(product_carts.quantity * (products.actual_price - products.actual_price * promotions.discount_percent / 100)) AS total_price").
		Joins("JOIN product_carts ON carts.id = product_carts.cart_id").
		Joins("JOIN products ON product_carts.product_id = products.id").
		Joins("JOIN product_promotions ON products.id = product_promotions.product_id").
		Joins("JOIN promotions ON product_promotions.promotion_id = promotions.id").
		Where("carts.user_id = ? AND products.is_actived = ? AND products.is_deleted = ? AND product_promotions.is_active = ? AND product_carts.is_deleted = ?", currentUserId, true, false, true, false).
		Group("carts.id").
		Scan(&shoppingCart).Error
	if err != nil {
		return nil, errors.New("Chưa có sản phẩm nào trong giỏ hàng")
	}

	var products []DTO.InfoProductCartDto
	s.db.Table("product_carts").
		Select("product_carts.id AS product_cart_id, products.id AS product_id, products.name, products.actual_price, products.category_id, products.description, products.price, products.is_actived, product_carts.quantity").
		Joins("JOIN products ON product_carts.product_id = products.id").
		Where("product_carts.cart_id = ? AND product_carts.is_deleted = ?", shoppingCart.CartID, false).
		Scan(&products)
	shoppingCart.Products = products

	return &shoppingCart, nil
}

func (s *cartServiceImpl) UpdateQuantity(input DTO.UpdateCartDto) (string, error) {
	var currentUserId int = getCurrentUserId()
	cart := &domain.Cart{}
	if err := s.db.First(&cart, "user_id = ?", currentUserId).Error; err != nil {
		return "", errors.New("Chưa tồn tại giỏ hàng!")
	}

	productCart := &domain.ProductCart{}
	if err := s.db.First(&productCart, "cart_id = ? AND product_id = ? AND is_deleted = ?", cart.ID, input.ProductID, false).Error; err != nil {
		return "", errors.New("Không tìm thấy sản phẩm có trong giỏ hàng!")
	}

	productCart.Quantity += input.Quantity
	if productCart.Quantity <= 0 {
		productCart.IsDeleted = true
		if err := s.db.Save(&productCart).Error; err != nil {
			return "", err
		}
		return "Sản phẩm đã được xóa khỏi giỏ hàng", nil
	}

	if err := s.db.Save(&productCart).Error; err != nil {
		return "", err
	}

	return "Cập nhật số lượng thành công.", nil
}

func getCurrentUserId() int {
	// Placeholder function to get the current user ID
	// Implement the logic to get the user ID from context or session
	return 1
}
