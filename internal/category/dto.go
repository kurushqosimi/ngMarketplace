package category

import (
	"encoding/json"
	"ngMarketplace/internal/common"
)

// getCategoryRequest represents the param request for getting a category
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

// updateCategoryRequest represents the request body for updating a category
type updateCategoryRequest struct {
	CategoryName    *string         `json:"category_name"`
	ParentID        *int            `json:"parent_id"`
	Language        *string         `json:"language"`
	AttributeSchema json.RawMessage `json:"attribute_schema"`
}

// getCategoriesRequest represents the request query for getting the list of categories
type getCategoriesRequest struct {
	CategoryName string `form:"category_name"`
	Language     string `form:"language"`
	common.Filters
}

// getCategoryByParentIDRequest represents the param request for getting a category
type getCategoryByParentIDRequest struct {
	ParentID int64 `uri:"parent_id" binding:"required,min=1"`
}
