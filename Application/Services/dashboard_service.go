package Services

import (
	"context"
	"go/foodappbe/Application/Services/DTO"
	domain "go/foodappbe/Domain/Models"
	"gorm.io/gorm"
	"sort"
	"time"
)

type IDashboardService interface {
	GetOrderStatistics(ctx context.Context, input DTO.DashboardFilterDto) (DTO.DashboardPageResultDto, error)
	GetOrderStatisticsByDay(ctx context.Context, input DTO.DashboardFilterByDayDto) (DTO.DashboardPageResultDto, error)
	GetOrderStatisticsByMonth(ctx context.Context, input DTO.DashboardFilterByMonthDto) (DTO.DashboardPageResultDto, error)
	GetTopProducts(ctx context.Context, input DTO.DashboardProductsFilterDto) (DTO.DashboardPageResultDto, error)
}

type DashboardService struct {
	db *gorm.DB
}

func NewDashboardService(db *gorm.DB) IDashboardService {
	return &DashboardService{db: db}
}

func (s *DashboardService) GetOrderStatistics(ctx context.Context, input DTO.DashboardFilterDto) (DTO.DashboardPageResultDto, error) {
	var result DTO.DashboardPageResultDto
	var orders []DTO.DashboardResponseDto

	query := s.db.Model(&domain.Order{}).
		Joins("JOIN order_details od ON orders.id = od.order_id").
		Joins("JOIN products p ON od.product_id = p.id").
		Joins("JOIN product_promotions pp ON p.id = pp.product_id").
		Joins("JOIN promotions pro ON pp.promotion_id = pro.id").
		Where("p.is_actived = ? AND p.is_deleted = ? AND pp.is_active = ? AND orders.status = ? AND orders.updated_at BETWEEN ? AND ?",
			true, false, true, "SUCCESS", input.StartDate, input.EndDate).
		Group("orders.id").
		Select("orders.id, orders.updated_at, SUM(od.quantity * (p.actual_price - (p.actual_price * pro.discount_percent / 100))) as order_revenue").
		Find(&orders).Error

	if query != nil {
		return DTO.DashboardPageResultDto{}, query
	}

	result.TotalItem = len(orders)
	for _, order := range orders {
		result.TotalRevenue += order.OrderRevenue
	}

	result.Item = orders[(input.PageIndex-1)*input.PageSize : input.PageIndex*input.PageSize]
	return result, nil
}

func (s *DashboardService) GetOrderStatisticsByDay(ctx context.Context, input DTO.DashboardFilterByDayDto) (DTO.DashboardPageResultDto, error) {
	var result DTO.DashboardPageResultDto
	var orders []DTO.DashboardResponseDto

	query := s.db.Model(&domain.Order{}).
		Joins("JOIN order_details od ON orders.id = od.order_id").
		Joins("JOIN products p ON od.product_id = p.id").
		Joins("JOIN product_promotions pp ON p.id = pp.product_id").
		Joins("JOIN promotions pro ON pp.promotion_id = pro.id").
		Where("p.is_actived = ? AND p.is_deleted = ? AND pp.is_active = ? AND orders.status = ? AND YEAR(orders.updated_at) = ? AND MONTH(orders.updated_at) = ? AND DAY(orders.updated_at) = ?",
			true, false, true, "SUCCESS", input.Year, input.Month, input.Day).
		Group("orders.id").
		Select("orders.id, orders.updated_at, SUM(od.quantity * (p.actual_price - (p.actual_price * pro.discount_percent / 100))) as order_revenue").
		Find(&orders).Error

	if query != nil {
		return DTO.DashboardPageResultDto{}, query
	}

	result.TotalItem = len(orders)
	for _, order := range orders {
		result.TotalRevenue += order.OrderRevenue
	}

	result.Item = orders[(input.PageIndex-1)*input.PageSize : input.PageIndex*input.PageSize]
	return result, nil
}

func (s *DashboardService) GetOrderStatisticsByMonth(ctx context.Context, input DTO.DashboardFilterByMonthDto) (DTO.DashboardPageResultDto, error) {
	var result DTO.DashboardPageResultDto
	var orders []DTO.DashboardResponseDto

	query := s.db.Model(&domain.Order{}).
		Joins("JOIN order_details od ON orders.id = od.order_id").
		Joins("JOIN products p ON od.product_id = p.id").
		Joins("JOIN product_promotions pp ON p.id = pp.product_id").
		Joins("JOIN promotions pro ON pp.promotion_id = pro.id").
		Where("p.is_actived = ? AND p.is_deleted = ? AND pp.is_active = ? AND orders.status = ? AND YEAR(orders.updated_at) = ? AND MONTH(orders.updated_at) = ?",
			true, false, true, "SUCCESS", input.Year, input.Month).
		Group("orders.id").
		Select("orders.id, orders.updated_at, SUM(od.quantity * (p.actual_price - (p.actual_price * pro.discount_percent / 100))) as order_revenue").
		Find(&orders).Error

	if query != nil {
		return DTO.DashboardPageResultDto{}, query
	}

	result.TotalItem = len(orders)
	for _, order := range orders {
		result.TotalRevenue += order.OrderRevenue
	}

	result.Item = orders[(input.PageIndex-1)*input.PageSize : input.PageIndex*input.PageSize]
	return result, nil
}

func (s *DashboardService) GetTopProducts(ctx context.Context, input DTO.DashboardProductsFilterDto) (DTO.DashboardPageResultDto, error) {
	var products []struct {
		Id              int       `json:"id"`
		OrderStatus     string    `json:"order_status"`
		Quantity        int       `json:"quantity"`
		ProductId       int       `json:"product_id"`
		Name            string    `json:"name"`
		Price           float64   `json:"price"`
		ActualPrice     float64   `json:"actual_price"`
		TotalPrice      float64   `json:"total_price"`
		Description     string    `json:"description"`
		CategoryId      int       `json:"category_id"`
		IsActived       bool      `json:"is_actived"`
		CreatedAt       time.Time `json:"created_at"`
		ProductImageUrl string    `json:"product_image_url"`
	}

	// SQL query to fetch the necessary details
	query := `
		SELECT od.id, od.quantity, od.product_id, p.name, p.price, 
			(p.actual_price - (p.actual_price * prom.discount_percent / 100)) AS actual_price,
			((p.actual_price - (p.actual_price * prom.discount_percent / 100)) * od.quantity) AS total_price, 
			p.description, p.category_id, p.is_actived, p.created_at, pi.product_image_url 
		FROM orders ord 
		JOIN order_details od ON ord.id = od.order_id 
		JOIN products p ON od.product_id = p.id 
		JOIN product_images pi ON p.id = pi.product_id 
		JOIN product_promotions pp ON p.id = pp.product_id
		JOIN promotions prom ON pp.promotion_id = prom.id
		WHERE ord.status = ?`

	// Execute the query
	if err := s.db.Raw(query, "SUCCESS").Scan(&products).Error; err != nil {
		return DTO.DashboardPageResultDto{}, err
	}

	// Group by Product ID and calculate totals
	productMap := make(map[int]*DTO.DashboardResponseTopProductsDto)
	for _, p := range products {
		if _, exists := productMap[p.ProductId]; !exists {
			productMap[p.ProductId] = &DTO.DashboardResponseTopProductsDto{
				Id:              p.ProductId,
				Name:            p.Name,
				Price:           p.Price,
				Description:     p.Description,
				CategoryId:      p.CategoryId,
				IsActived:       p.IsActived,
				CreatedAt:       p.CreatedAt,
				ProductImageUrl: p.ProductImageUrl,
				TotalQuantity:   0,
				TotalRevenue:    0,
			}
		}
		productMap[p.ProductId].TotalQuantity += p.Quantity
		productMap[p.ProductId].TotalRevenue += p.TotalPrice
	}

	// Convert map to slice and sort by TotalQuantity
	var topProducts []DTO.DashboardResponseTopProductsDto
	for _, product := range productMap {
		topProducts = append(topProducts, *product)
	}

	sort.Slice(topProducts, func(i, j int) bool {
		return topProducts[i].TotalQuantity > topProducts[j].TotalQuantity
	})

	// Apply pagination
	start := (input.PageIndex - 1) * input.PageSize
	end := start + input.PageSize
	if start > len(topProducts) {
		start = len(topProducts)
	}
	if end > len(topProducts) {
		end = len(topProducts)
	}
	paginatedProducts := topProducts[start:end]

	paginatedDashboardResponse := convertToDashboardResponse(paginatedProducts)

	// Prepare result
	result := DTO.DashboardPageResultDto{
		TotalItem: len(topProducts),
		Item:      paginatedDashboardResponse,
	}

	return result, nil
}

func convertToDashboardResponse(products []DTO.DashboardResponseTopProductsDto) []DTO.DashboardResponseDto {
	var result []DTO.DashboardResponseDto

	for _, p := range products {
		item := DTO.DashboardResponseDto{
			Id:           p.Id,
			UpdatedDate:  p.CreatedAt,    // Example field mapping
			OrderRevenue: p.TotalRevenue, // Example field mapping
			Products: []DTO.InfoProductCartDto{ // Example field mapping
				{
					ProductID:       p.Id,
					Name:            p.Name,
					ActualPrice:     p.Price,
					CategoryID:      p.CategoryId,
					Description:     p.Description,
					ProductImageUrl: p.ProductImageUrl,
					Quantity:        p.TotalQuantity,
					Price:           p.Price,
					IsActive:        p.IsActived,
				},
			},
		}
		result = append(result, item)
	}

	return result
}
