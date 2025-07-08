package product

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"ngMarketplace/internal/apperror"
	"ngMarketplace/internal/transport/http/router"
	"ngMarketplace/pkg/logger"
)

const (
	productURL  = "/product/:id"
	productsURL = "/products"
)

type UseCase interface {
	CreateProduct(ctx context.Context, product *Product) error
}

type Handler struct {
	useCase UseCase
	logger  logger.Logger
}

func NewHandler(useCase UseCase, logger logger.Logger) *Handler {
	return &Handler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *Handler) Register(router *gin.Engine) {
	router.POST(productsURL, h.createProductHandler)
}

// createProductHandler creates a new Product in Marketplace
func (h *Handler) createProductHandler(ctx *gin.Context) {
	const op = "CreateProductHandler"

	var req createProductRequest
	if err := ctx.BindJSON(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindJSON: %v", op, err)
		apperror.WriteBadRequestResponse(ctx, ErrBindJSON, "Something is missing or was not sent correctly")
		return
	}

	product := &Product{
		Price:      req.Price,
		Currency:   req.Currency,
		CategoryID: req.CategoryID,
		UserID:     req.UserID,
	}

	if err := h.useCase.CreateProduct(ctx, product); err != nil {
		h.logger.Error("%s: h.useCase.Create: %v", op, err)
		switch {
		case errors.Is(err, ErrProductValidationFailed):
			apperror.WriteBadRequestResponse(ctx, err, err.Error())
		case errors.Is(err, ErrInvalidForeignKey):
			apperror.WriteBadRequestResponse(ctx, err, "Entered wrong category, or please sign out and sign in again")
		case errors.Is(err, ErrConnectionFailed):
			apperror.WriteSrvUnResponse(ctx, err, "Database connection failed")
		default:
			apperror.WriteInternalErrResponse(ctx, err, "Internal server error")
		}
		return
	}

	if err := router.WriteJSON(ctx, http.StatusCreated, gin.H{"product": product}, nil); err != nil {
		h.logger.Warn("%s: router.WriteJSON: %v", op, err)
		ctx.JSON(http.StatusCreated, gin.H{"product": product})
		return
	}
}
