package product

import (
	"context"
	"errors"
	"fmt"
	"ngMarketplace/pkg/postgres"
	"time"
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

// Update method updates product
func (r *Repository) Update(ctx context.Context, product *Product) error {
	const op = "Update"

	query := `
		UPDATE 
		    products
		SET 
		    price = $1, 
		    currency = $2, 
		    category_id = $3
		WHERE 
		    product_id = $4
		AND 
		    active = true
		RETURNING updated_at`

	args := []interface{}{
		product.Price,
		product.Currency,
		product.CategoryID,
		product.ProductID,
	}

	if err := r.client.Pool.QueryRow(
		ctx,
		query,
		args...,
	).Scan(&product.UpdatedAt); err != nil {
		switch {
		case errors.Is(err, postgres.ErrNoRows):
			return ErrProductNotFound
		default:
			return postgres.ErrDoQuery(op, err)
		}
	}

	return nil
}

// SoftDelete method deletes product softly, meaning that it makes active false and that's it
func (r *Repository) SoftDelete(ctx context.Context, id int64) error {
	const op = "Delete"

	query := `
		UPDATE 
		    products
		SET 
		    deleted_at = now(), 
		    active = false
		WHERE 
		    product_id = $1 
		AND 
		    active = true
		RETURNING deleted_at`

	var deletedAt *time.Time
	err := r.client.Pool.QueryRow(ctx, query, id).Scan(&deletedAt)
	if err != nil {
		switch {
		case errors.Is(err, postgres.ErrNoRows):
			return ErrProductNotFound
		default:
			return postgres.ErrDoQuery(op, err)
		}
	}

	return nil
}
