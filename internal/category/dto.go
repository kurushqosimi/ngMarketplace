package category

import "encoding/json"

type getCategoryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// createCategoryRequest represents the request body for creating a category
type createCategoryRequest struct {
	CategoryName    string          `json:"category_name" binding:"required"`
	ParentID        *int            `json:"parent_id"`
	Language        string          `json:"language" binding:"required,oneof=tj ru en"`
	AttributeSchema json.RawMessage `json:"attribute_schema"`
}
