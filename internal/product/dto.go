package product

import "ngMarketplace/internal/common"

// createProductRequest represents a request body for creating a product
type createProductRequest struct {
	Price      float64 `json:"price" binding:"required,min=1"`
	Currency   string  `json:"currency" binding:"required"`
	CategoryID int     `json:"category_id" binding:"required,min=1"`
	UserID     int     `json:"user_id" binding:"required,min=1"` // todo user_id should be got from token
}

// getProductRequest represents the param request for getting a product
type getProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// updateProductRequest represents a request body for updating a product
type updateProductRequest struct {
	Price      *float64 `json:"price"`
	Currency   *string  `json:"currency"`
	CategoryID *int     `json:"category_id"`
}

// getProductsRequest represents a query for getting products by filters
type getProductsRequest struct {
	FromPrice  float64 `form:"from_price"`
	ToPrice    float64 `form:"to_price"`
	Currency   string  `form:"currency"`
	CategoryID int     `form:"category_id"`
	UserID     int     `form:"user_id"`
	common.Filters
}
