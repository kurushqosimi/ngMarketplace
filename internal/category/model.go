package category

import (
	"encoding/json"
	"errors"
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

var (
	ErrUpdate   = errors.New("no active category to update")
	ErrDelete   = errors.New("no active category to delete")
	ErrNotFound = errors.New("category not found")
)
