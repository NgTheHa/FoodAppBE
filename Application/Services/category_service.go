package Services

import (
	"context"
	"errors"
	"go/foodappbe/Application/Services/DTO"
	domain "go/foodappbe/Domain/Models"
	"go/foodappbe/Utils"
	"gorm.io/gorm"
	"time"
)

type ICategoryService interface {
	Create(ctx context.Context, input DTO.CreateCategoryDto) error
	GetCategoryPaging(ctx context.Context, input DTO.CategoryFilterDto) (DTO.PageResultDto, error)
	GetCategoryById(ctx context.Context, id int) (DTO.CategoryResponseDto, error)
	UpdateCategory(ctx context.Context, input DTO.UpdateCategoryDto) error
	DeleteCategory(ctx context.Context, id int) error
}

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) ICategoryService {
	return &CategoryService{db: db}
}

func (s *CategoryService) Create(ctx context.Context, input DTO.CreateCategoryDto) error {
	var count int64
	s.db.Model(&domain.Category{}).Where("name = ?", input.Name).Count(&count)
	if count > 0 {
		return errors.New("category already exists")
	}

	category := &domain.Category{
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now(),
		CreatedBy:   Utils.GetUserId(ctx),
	}

	return s.db.Create(category).Error
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id int) error {
	var category domain.Category
	if err := s.db.First(&category, id).Error; err != nil {
		return errors.New("category not found")
	}
	category.IsDeleted = true
	category.UpdatedAt = time.Now()
	category.UpdatedBy = Utils.GetUserId(ctx)

	// Delete related products
	s.db.Model(&domain.Product{}).Where("category_id = ?", id).Update("is_deleted", true)
	s.db.Model(&domain.ProductImage{}).Where("category_id = ?", id).Update("is_deleted", true)
	s.db.Model(&domain.ProductPromotion{}).Where("category_id = ?", id).Update("is_active", false)

	return s.db.Save(&category).Error
}

func (s *CategoryService) GetCategoryById(_ context.Context, id int) (DTO.CategoryResponseDto, error) {
	var category domain.Category
	if err := s.db.First(&category, id).Error; err != nil || category.IsDeleted {
		return DTO.CategoryResponseDto{}, errors.New("category not found")
	}
	return DTO.CategoryResponseDto{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, input DTO.UpdateCategoryDto) error {
	var category domain.Category
	if err := s.db.First(&category, input.Id).Error; err != nil || category.IsDeleted {
		return errors.New("category not found")
	}
	category.Name = input.Name
	category.Description = input.Description
	category.UpdatedAt = time.Now()
	category.UpdatedBy = Utils.GetUserId(ctx)

	return s.db.Save(&category).Error
}

func (s *CategoryService) GetCategoryPaging(_ context.Context, input DTO.CategoryFilterDto) (DTO.PageResultDto, error) {
	var categories []domain.Category
	query := s.db.Model(&domain.Category{}).Where("is_deleted = ?", false)
	if input.Name != "" {
		query = query.Where("name ILIKE ?", "%"+input.Name+"%")
	}
	totalItem := int64(0)
	query.Count(&totalItem)

	query = query.Offset((input.PageIndex - 1) * input.PageSize).Limit(input.PageSize)
	query.Find(&categories)

	categoryDtos := make([]DTO.CategoryResponseDto, len(categories))
	for i, category := range categories {
		categoryDtos[i] = DTO.CategoryResponseDto{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
	}

	return DTO.PageResultDto{
		Items:     categoryDtos,
		TotalItem: int(totalItem),
	}, nil
}
