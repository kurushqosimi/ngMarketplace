package category

import "encoding/json"

type getCategoryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type createCategoryRequest struct {
	CategoryName    string          `json:"category_name" binding:"required"`
	ParentID        int             `json:"parent_id"`
	AttributeSchema json.RawMessage `json:"attribute_schema"`
}
