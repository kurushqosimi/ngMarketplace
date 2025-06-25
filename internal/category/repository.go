package category

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"ngMarketplace/pkg/data"
	"ngMarketplace/pkg/postgres"
	"time"
)

type Repository struct {
	client postgres.Postgres
}

func NewRepository(client postgres.Postgres) *Repository {
	return &Repository{client: client}
}

func (r *Repository) Create(ctx context.Context, category *Category) error {
	const op = "Create"

	query := `
		INSERT INTO categories (category_name, parent_id, language,attribute_schema)
		VALUES ($1, $2, $3, $4)
		RETURNING category_id, created_at, active`

	args := []interface{}{
		category.CategoryName,
		category.ParentID,
		category.Language,
		category.AttributeSchema,
	}

	if err := r.client.Pool.QueryRow(
		ctx,
		query,
		args...,
	).Scan(
		&category.CategoryID,
		&category.CreatedAt,
		&category.Active,
	); err != nil {
		if postgres.IsPgErr(err) {
			err = postgres.Conv2CustomErr(err)
		}

		var pgErr *postgres.PostgresErr
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				return postgres.ErrDoQuery(op, ErrDuplicateCategory)
			case "23503":
				return postgres.ErrDoQuery(op, ErrInvalidParentID)
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

func (r *Repository) GetAll(ctx context.Context) ([]*Category, error) {
	const op = "GetAll"

	query := `
		SELECT category_id, category_name, parent_id, attribute_schema, created_at, active, updated_at, deleted_at
		FROM categories
		WHERE active = true`

	rows, err := r.client.Pool.Query(ctx, query)
	if err != nil {
		return nil, postgres.ErrDoQuery(op, err)
	}
	defer rows.Close()

	categories := []*Category{}

	for rows.Next() {
		var category Category
		err = rows.Scan(
			&category.CategoryID,
			&category.CategoryName,
			&category.ParentID,
			&category.AttributeSchema,
			&category.CreatedAt,
			&category.Active,
			&category.UpdatedAt,
			&category.DeletedAt,
		)
		if err != nil {
			return nil, postgres.ErrScan(op, err)
		}

		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		return nil, postgres.ErrReadRows(op, err)
	}

	return categories, nil
}

func (r *Repository) FindOne(ctx context.Context, id string) (*Category, error) {
	const op = "FindOne"

	query := `
		SELECT category_id, category_name, parent_id, attribute_schema, created_at, active, updated_at, deleted_at
		FROM categories
		WHERE active = true AND category_id = $1
		LIMIT 1`

	var category Category

	if err := r.client.Pool.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&category.CategoryID,
		&category.CategoryName,
		&category.ParentID,
		&category.AttributeSchema,
		&category.CreatedAt,
		&category.Active,
		&category.UpdatedAt,
		&category.DeletedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, postgres.ErrDoQuery(op, err)
	}

	return &category, nil
}

func (r *Repository) Update(ctx context.Context, category *Category) error {
	const op = "Update"

	query := `
		UPDATE categories
		SET category_name = $1, parent_id = $2, attribute_schema = $3
		WHERE category_id = $4 AND active = true
		RETURNING updated_at`

	args := []interface{}{
		category.CategoryName,
		category.ParentID,
		category.AttributeSchema,
		category.CategoryID,
	}

	if err := r.client.Pool.QueryRow(
		ctx,
		query,
		args...,
	).Scan(&category.UpdatedAt); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return ErrUpdate
		default:
			return postgres.ErrDoQuery(op, err)
		}
	}

	return nil
}

func (r *Repository) SoftDelete(ctx context.Context, id string) error {
	const op = "Delete"

	query := `
		UPDATE categories
		SET deleted_at = now(), active = false
		WHERE category_id = $1 AND active = true`

	var deletedAt time.Time
	err := r.client.Pool.QueryRow(ctx, query, id).Scan(&deletedAt)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return ErrDelete
		default:
			return postgres.ErrDoQuery(op, err)
		}
	}

	return nil
}

func (r *Repository) GetByParentID(ctx context.Context, parentID string) ([]*Category, error) {
	const op = "GetByParentID"

	query := `
		SELECT category_id, category_name, parent_id, attribute_schema, created_at, active, updated_at, deleted_at
		FROM categories
		WHERE active = true AND parent_id = $1`

	categories := []*Category{}

	rows, err := r.client.Pool.Query(ctx, query, parentID)
	if err != nil {
		return nil, postgres.ErrDoQuery(op, err)
	}

	for rows.Next() {
		var category Category
		err = rows.Scan(
			&category.CategoryID,
			&category.CategoryName,
			&category.ParentID,
			&category.AttributeSchema,
			&category.CreatedAt,
			&category.Active,
			&category.UpdatedAt,
			&category.DeletedAt,
		)
		if err != nil {
			return nil, postgres.ErrScan(op, err)
		}

		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		return nil, postgres.ErrReadRows(op, err)
	}

	return categories, nil
}

func (r *Repository) GetPaginated(ctx context.Context, categoryName string, filters data.Filters) ([]*Category, data.Metadata, error) {
	const op = "GetPaginated"

	query := fmt.Sprintf(`
		SELECT count(*) OVER(), category_id, category_name, parent_id, attribute_schema, created_at, active, updated_at, deleted_at
		FROM categories
		WHERE (to_tsvector('simple', category_name) @@ plainto_tsquery('simple', $1) OR $1 = '')
		ORDER BY %s %s, id ASC
		LIMIT $2 OFFSET $3`, filters.SortColumn(), filters.SortDirection())

	args := []interface{}{categoryName, filters.Limit(), filters.Offset()}

	rows, err := r.client.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, data.Metadata{}, postgres.ErrDoQuery(op, err)
	}
	defer rows.Close()

	totalRecords := 0
	categories := []*Category{}

	for rows.Next() {
		var category Category
		err = rows.Scan(
			&totalRecords,
			&category.CategoryID,
			&category.CategoryName,
			&category.ParentID,
			&category.AttributeSchema,
			&category.CreatedAt,
			&category.Active,
			&category.UpdatedAt,
			&category.DeletedAt,
		)
		if err != nil {
			return nil, data.Metadata{}, postgres.ErrScan(op, err)
		}

		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		return nil, data.Metadata{}, postgres.ErrReadRows(op, err)
	}

	// TODO implement it in service
	metadata := data.CalculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return categories, metadata, nil
}

func (r *Repository) Restore(ctx context.Context, categoryID string) error {
	const op = "Restore"

	query := `
		UPDATE categories
		SET deleted_at = NULL, active = true
		WHERE category_id = $1`

	result, err := r.client.Pool.Exec(ctx, query, categoryID)
	if err != nil {
		return postgres.ErrExec(op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

// TODO implement it in service GetCategoryTree
