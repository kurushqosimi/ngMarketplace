package product

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
