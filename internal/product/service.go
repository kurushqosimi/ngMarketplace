package product

import (
	"context"
	"fmt"
	"ngMarketplace/pkg/validator"
)

type Storage interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id int64) (*Product, error)
	Update(ctx context.Context, product *Product) error
	SoftDelete(ctx context.Context, id int64) error
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
