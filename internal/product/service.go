package product

import (
	"context"
	"fmt"
	"ngMarketplace/pkg/validator"
)

type Storage interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id int64) (*Product, error)
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
