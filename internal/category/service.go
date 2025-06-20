package category

import (
	"context"
	"ngMarketplace/pkg/data"
)

type Storage interface {
	Create(ctx context.Context, category *Category) error
	GetAll(ctx context.Context) ([]*Category, error)
	FindOne(ctx context.Context, id string) (*Category, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id string) error
	GetByParentID(ctx context.Context, parentID string) ([]*Category, error)
	GetPaginated(ctx context.Context, categoryName string, filters data.Filters) ([]*Category, data.Metadata, error)
	Restore(ctx context.Context, categoryID string) error
}

type Service struct {
	Repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{Repository: repository}
}

func (u Service) Create() {
	//TODO implement me
	panic("implement me")
}

func (u Service) ListAll() {
	//TODO implement me
	panic("implement me")
}

func (u Service) GetTree() {
	//TODO implement me
	panic("implement me")
}
