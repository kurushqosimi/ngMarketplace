package product

import (
	"errors"
	"ngMarketplace/pkg/validator"
	"time"
)

// Product represents a product in marketplace
type Product struct {
	ProductID  int        `json:"product_id"`
	Price      float64    `json:"price"`
	Currency   string     `json:"currency"`
	CategoryID int        `json:"category_id"`
	UserID     int        `json:"user_id"`
	Active     bool       `json:"-"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"-"`
}

func validateProduct(v *validator.Validator, product *Product) {
	v.Check(validator.In(product.Currency, "TJS", "RUB", "USD"), "currency", "must be TJS, RUB, or USD")
}

// Repository Errors
var (
	ErrDuplicateProduct  = errors.New("product already exists")
	ErrInvalidForeignKey = errors.New("category_id or user_id does not exist")
	ErrConnectionFailed  = errors.New("database connection failed")
	ErrProductNotFound   = errors.New("product not found")
)

// Service Errors
var (
	ErrProductValidationFailed = errors.New("product validation failed")
)

// Handler Errors
var (
	ErrBindJSON    = errors.New("failed binding json")
	ErrInvalidID   = errors.New("invalid category id was sent")
	ErrFailedQuery = errors.New("failed to parse query")
)
