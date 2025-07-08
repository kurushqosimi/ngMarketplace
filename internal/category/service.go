package category

import (
	"context"
	"fmt"
	"ngMarketplace/internal/common"
	"ngMarketplace/pkg/validator"
)

type Storage interface {
	Create(ctx context.Context, category *Category) error
	GetByID(ctx context.Context, id int64) (*Category, error)
	Update(ctx context.Context, category *Category) error
	SoftDelete(ctx context.Context, id int64) error
	GetPaginated(ctx context.Context, categoryName string, language string, filters common.Filters) ([]*Category, int, error)
	GetByParentID(ctx context.Context, parentID int64) ([]*Category, error)
	Restore(ctx context.Context, categoryID int64) error
}

type Service struct {
	Repository Storage
}

func NewUseCase(repository Storage) *Service {
	return &Service{Repository: repository}
}

func (s *Service) Create(ctx context.Context, category *Category) error {
	v := validator.New()

	if validateCategory(v, category); !v.Valid() {
		return fmt.Errorf("%w: %w", ErrCategoryValidationFailed, v.Errors)
	}

	if err := s.Repository.Create(ctx, category); err != nil {
		return fmt.Errorf("failed to create a category: %w", err)
	}

	return nil
}

func (s *Service) GetCategory(ctx context.Context, categoryID int64) (*Category, error) {
	return s.Repository.GetByID(ctx, categoryID)
}

func (s *Service) UpdateCategory(ctx context.Context, categoryID int64, newCategory *updateCategoryRequest) (*Category, error) {
	category, err := s.Repository.GetByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	if newCategory.CategoryName != nil {
		category.CategoryName = *newCategory.CategoryName
	}

	if newCategory.ParentID != nil {
		category.ParentID = newCategory.ParentID
	}

	if newCategory.Language != nil {
		category.Language = *newCategory.Language
	}

	if newCategory.AttributeSchema != nil {
		category.AttributeSchema = newCategory.AttributeSchema
	}

	v := validator.New()

	if validateCategory(v, category); !v.Valid() {
		return nil, fmt.Errorf("%w: %w", ErrCategoryValidationFailed, v.Errors)
	}

	if err = s.Repository.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) DeleteCategory(ctx context.Context, categoryID int64) error {
	return s.Repository.SoftDelete(ctx, categoryID)
}

func (s *Service) GetCategories(ctx context.Context, filters getCategoriesRequest) ([]*Category, common.Metadata, error) {
	if filters.Page == 0 {
		filters.Page = 1
	}

	if filters.PageSize == 0 {
		filters.PageSize = 20
	}

	if filters.Sort == "" {
		filters.Sort = "category_name"
	}

	if filters.Language == "" {
		filters.Language = "ru"
	}

	filters.SortSafeList = []string{"category_id", "category_name", "parent_id", "-category_id", "-category_name", "-parent_id"}

	v := validator.New()

	v.Check(validator.In(filters.Language, "ru", "tj", "en"), "language", "language must be one of [tj ru en]")

	if common.ValidateFilters(v, filters.Filters); !v.Valid() {
		return nil, common.Metadata{}, fmt.Errorf("%w: %w", common.ErrFilterValidationFailed, v.Errors)
	}

	categories, totalRecords, err := s.Repository.GetPaginated(ctx, filters.CategoryName, filters.Language, filters.Filters)
	if err != nil {
		return nil, common.Metadata{}, err
	}

	metadata := common.CalculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return categories, metadata, nil
}

func (s *Service) GetCategoryByParentID(ctx context.Context, parentID int64) ([]*Category, error) {
	return s.Repository.GetByParentID(ctx, parentID)
}
