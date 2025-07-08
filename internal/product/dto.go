package product

// createProductRequest represents a request body for creating a product
type createProductRequest struct {
	Price      float64 `json:"price" binding:"required,min=1"`
	Currency   string  `json:"currency" binding:"required"`
	CategoryID int     `json:"category_id" binding:"required,min=1"`
	UserID     int     `json:"user_id" binding:"required,min=1"`
}
