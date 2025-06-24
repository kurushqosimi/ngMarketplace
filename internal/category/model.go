package category

import (
	"encoding/json"
	"errors"
	"ngMarketplace/pkg/validator"
	"time"
)

type Category struct {
	CategoryID      int             `json:"category_id"`
	CategoryName    string          `json:"category_name"`
	ParentID        int             `json:"parent_id"`
	AttributeSchema json.RawMessage `json:"attribute_schema"`
	CreatedAt       time.Time       `json:"-"`
	Active          bool            `json:"-"`
	UpdatedAt       time.Time       `json:"-"`
	DeletedAt       time.Time       `json:"-"`
}

func ValidateCategory(v *validator.Validator, category *Category) {
	v.Check(len(category.CategoryName) <= 50, "category_name", "must not be more than 50 bytes long")
	if len(category.AttributeSchema) > 0 {

	}
}

var (
	ErrUpdate   = errors.New("no active category to update")
	ErrDelete   = errors.New("no active category to delete")
	ErrNotFound = errors.New("category not found")
)
