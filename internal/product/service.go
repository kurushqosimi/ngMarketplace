package product

import (
	"context"
	"fmt"
	"ngMarketplace/internal/common"
	"ngMarketplace/pkg/validator"
)

type Storage interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id int64) (*Product, error)
	Update(ctx context.Context, product *Product) error
	SoftDelete(ctx context.Context, id int64) error
	GetPaginated(ctx context.Context, currency string, categoryID int, userID int, fromPrice float64, toPrice float64, filters common.Filters) ([]*Product, int, error)
}

type Service struct {
	Repository Storage
}

func NewUseCase(repository Storage) *Service {
	return &Service{Repository: repository}
}

func (s *Service) CreateProduct(ctx context.Context, product *Product) error {
	v := validator.New()

	if validateProduct(v, product); !v.Valid() {
		return fmt.Errorf("%w: %w", ErrProductValidationFailed, v.Errors)
	}

	if err := s.Repository.Create(ctx, product); err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}

func (s *Service) GetProduct(ctx context.Context, id int64) (*Product, error) {
	return s.Repository.GetByID(ctx, id)
}

func (s *Service) UpdateProduct(ctx context.Context, id int64, request *updateProductRequest) (*Product, error) {
	product, err := s.Repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if request.CategoryID != nil {
		product.CategoryID = *request.CategoryID
	}

	if request.Price != nil {
		product.Price = *request.Price
	}

	if request.Currency != nil {
		product.Currency = *request.Currency
	}

	v := validator.New()

	if validateProduct(v, product); !v.Valid() {
		return nil, fmt.Errorf("%w: %w", ErrProductValidationFailed, v.Errors)
	}

	if err = s.Repository.Update(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) DeleteProduct(ctx context.Context, id int64) error {
	return s.Repository.SoftDelete(ctx, id)
}

func (s *Service) GetProducts(ctx context.Context, filters getProductsRequest) ([]*Product, common.Metadata, error) {
	if filters.Page == 0 {
		filters.Page = 1
	}

	if filters.PageSize == 0 {
		filters.PageSize = 20
	}

	if filters.Sort == "" {
		filters.Sort = "product_id"
	}

	if filters.Currency == "" {
		filters.Currency = "TJS"
	}

	filters.SortSafeList = []string{"product_id", "price", "-product_id", "-price"}

	v := validator.New()

	v.Check(filters.ToPrice >= 0, "to_price", "to_price cannot be negative")
	v.Check(filters.FromPrice >= 0, "from_price", "from_price cannot be negative")
	v.Check(validator.In(filters.Currency, "TJS", "RUB", "USD"), "currency", "currency must be one of TJS, RUB, USD")
	v.Check(filters.CategoryID >= 0, "category_id", "category_id cannot be negative")
	v.Check(filters.UserID >= 0, "user_id", "user_id cannot be negative")

	if common.ValidateFilters(v, filters.Filters); !v.Valid() {
		return nil, common.Metadata{}, fmt.Errorf("%w: %w", common.ErrFilterValidationFailed, v.Errors)
	}

	products, totalRecords, err := s.Repository.GetPaginated(
		ctx,
		filters.Currency,
		filters.CategoryID,
		filters.UserID,
		filters.FromPrice,
		filters.ToPrice,
		filters.Filters,
	)
	if err != nil {
		return nil, common.Metadata{}, err
	}

	metadata := common.CalculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return products, metadata, nil
}
