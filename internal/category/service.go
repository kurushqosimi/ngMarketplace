package category

import (
	"context"
	"fmt"
	"ngMarketplace/pkg/data"
	"ngMarketplace/pkg/validator"
	"strconv"
)

type Storage interface {
	Create(ctx context.Context, category *Category) error
	GetByID(ctx context.Context, id string) (*Category, error)
	GetAll(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category *Category) error
	SoftDelete(ctx context.Context, id string) error
	GetByParentID(ctx context.Context, parentID string) ([]*Category, error)
	GetPaginated(ctx context.Context, categoryName string, filters data.Filters) ([]*Category, data.Metadata, error)
	Restore(ctx context.Context, categoryID string) error
}

type Service struct {
	Repository Repository
}

func NewUseCase(repository Repository) *Service {
	return &Service{Repository: repository}
}

func (s *Service) Create(ctx context.Context, category *Category) error {
	v := validator.New()

	if validateCategory(v, category); !v.Valid() {
		return fmt.Errorf("%w: %w", v.Errors, ErrValidationFailed)
	}

	if err := s.Repository.Create(ctx, category); err != nil {
		return fmt.Errorf("failed to create a category: %w", err)
	}

	return nil
}

func (s *Service) GetCategory(ctx context.Context, categoryID int64) (*Category, error) {
	categoryIDStr := strconv.Itoa(int(categoryID))
	return s.Repository.GetByID(ctx, categoryIDStr)
}
