package product

import (
	"context"
	"errors"
	"fmt"
	"ngMarketplace/pkg/postgres"
)

type Repository struct {
	client *postgres.Postgres
}

func NewRepository(client *postgres.Postgres) *Repository {
	return &Repository{client: client}
}

// Create method creates a new product in db
func (r *Repository) Create(ctx context.Context, product *Product) error {
	const op = "Create"

	query := `
		INSERT INTO 
		    products (price, currency, category_id, user_id)
		VALUES 
		    ($1, $2, $3, $4)
		RETURNING product_id, created_at, active`

	args := []interface{}{
		product.Price,
		product.Currency,
		product.CategoryID,
		product.UserID,
	}

	if err := r.client.Pool.QueryRow(
		ctx,
		query,
		args...,
	).Scan(
		&product.ProductID,
		&product.CreatedAt,
		&product.Active,
	); err != nil {
		if postgres.IsPgErr(err) {
			err = postgres.Conv2CustomErr(err)
		}

		var pgErr *postgres.PostgresErr
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23503":
				return postgres.ErrDoQuery(op, ErrInvalidForeignKey)
			case "08000", "08001", "08003", "08006":
				return postgres.ErrDoQuery(op, ErrConnectionFailed)
			default:
				return postgres.ErrDoQuery(op, fmt.Errorf("unexpected database error: %w", err))
			}
		}
		return postgres.ErrDoQuery(op, err)
	}

	return nil
}

// GetByID method gets a product by ID
func (r *Repository) GetByID(ctx context.Context, id int64) (*Product, error) {
	const op = "GetByID"

	query := `
		SELECT 
		    product_id, price, currency, category_id, user_id, created_at, active, updated_at, deleted_at
		FROM 
		    products
		WHERE 
		    active = true 
		AND 
			product_id = $1
		LIMIT 1`

	var product Product

	if err := r.client.Pool.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&product.ProductID,
		&product.Price,
		&product.Currency,
		&product.CategoryID,
		&product.UserID,
		&product.CreatedAt,
		&product.Active,
		&product.UpdatedAt,
		&product.DeletedAt,
	); err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return nil, ErrProductNotFound
		}
		return nil, postgres.ErrDoQuery(op, err)
	}

	return &product, nil
}
