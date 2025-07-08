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
	productURL  = "/products/:id"
	productsURL = "/products"
)

type UseCase interface {
	CreateProduct(ctx context.Context, product *Product) error
	GetProduct(ctx context.Context, id int64) (*Product, error)
	UpdateProduct(ctx context.Context, id int64, request *updateProductRequest) (*Product, error)
	DeleteProduct(ctx context.Context, id int64) error
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
	router.GET(productURL, h.showProductHandler)
	router.PATCH(productURL, h.updateProductHandler)
	router.DELETE(productURL, h.deleteProductHandler)
	router.GET(productsURL, h.listProductsHandler)
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

// showProductHandler gets a product by product_id
func (h *Handler) showProductHandler(ctx *gin.Context) {
	const op = "getProductHandler"

	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindUri: %v", op, err)
		apperror.WriteBadRequestResponse(ctx, ErrInvalidID, "Provide correct product id")
		return
	}

	product, err := h.useCase.GetProduct(ctx, req.ID)
	if err != nil {
		h.logger.Error("%s: h.useCase.GetProduct: %v", op, err)
		switch {
		case errors.Is(err, ErrProductNotFound):
			apperror.WriteNotFoundResponse(ctx, err, "Product you are seeking does not exist")
		default:
			apperror.WriteInternalErrResponse(ctx, err, "Unexpected error occurred")
		}
		return
	}

	if err = router.WriteJSON(ctx, http.StatusOK, gin.H{"product": product}, nil); err != nil {
		h.logger.Warn("%s: router.WriteJSON: %v", op, err)
		ctx.JSON(http.StatusOK, gin.H{"product": product})
		return
	}
}

// updateProductHandler updates product currency, price, or category
func (h *Handler) updateProductHandler(ctx *gin.Context) {
	const op = "updateProductHandler"

	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindUri: %v", op, err)
		apperror.WriteBadRequestResponse(ctx, ErrInvalidID, "Provide correct product id")
		return
	}

	var input updateProductRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Error("%s: ctx.ShouldBindJSON: %v", op, err)
		apperror.WriteBadRequestResponse(ctx, ErrBindJSON, "Something is missing or was not sent correctly")
		return
	}

	product, err := h.useCase.UpdateProduct(ctx, req.ID, &input)
	if err != nil {
		h.logger.Error("%s: h.useCase.UpdateProduct: %v", op, err)
		switch {
		case errors.Is(err, ErrProductNotFound):
			apperror.WriteNotFoundResponse(ctx, err, "Product you are seeking to update does not exist")
		case errors.Is(err, ErrProductValidationFailed):
			apperror.WriteBadRequestResponse(ctx, err, err.Error())
		default:
			apperror.WriteInternalErrResponse(ctx, err, "Unexpected error occurred")
		}
		return
	}

	if err = router.WriteJSON(ctx, http.StatusOK, gin.H{"product": product}, nil); err != nil {
		h.logger.Warn("%s: router.WriteJSON: %v", op, err)
		ctx.JSON(http.StatusOK, gin.H{"product": product})
		return
	}
}

// deleteProductHandler makes Product's field active false
func (h *Handler) deleteProductHandler(ctx *gin.Context) {
	const op = "deleteProductHandler"

	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindUri: %v", op, err)
		apperror.WriteBadRequestResponse(ctx, ErrInvalidID, "Provide correct product id")
		return
	}

	if err := h.useCase.DeleteProduct(ctx, req.ID); err != nil {
		h.logger.Error("%s: h.useCase.DeleteProduct: %v", op, err)
		switch {
		case errors.Is(err, ErrProductNotFound):
			apperror.WriteNotFoundResponse(ctx, err, "Product you are seeking to delete does not exist")
		default:
			apperror.WriteInternalErrResponse(ctx, err, "Unexpected error occurred")
		}
		return
	}

	if err := router.WriteJSON(ctx, http.StatusOK, gin.H{"message": "product was successfully deleted"}, nil); err != nil {
		h.logger.Warn("%s: router.WriteJSON: %v", op, err)
		ctx.JSON(http.StatusOK, gin.H{"message": "product was successfully deleted"})
		return
	}
}

func (h *Handler) listProductsHandler(ctx *gin.Context) {

}
