package category

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
	categoriesURL = "/categories"
	categoryURL   = "/categories/:id"
)

type UseCase interface {
	Create()
	ListAll()
	GetTree()
	GetCategory(ctx context.Context, categoryID int64) (*Category, error)
}

type Handler struct {
	useCase UseCase
	logger  logger.Logger
}

func NewHandler(usecase UseCase, logger logger.Logger) *Handler {
	return &Handler{
		useCase: usecase,
		logger:  logger,
	}
}

func (h *Handler) Register(router *gin.Engine) {
	router.GET(categoryURL, h.ShowCategoryHandler)
	router.POST(categoriesURL, h.CreateCategoryHandler)
}

func (h *Handler) ShowCategoryHandler(ctx *gin.Context) {
	const op = "GetCategoryHandler"

	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindUri: %v", op, err)
		apperror.WriteError(ctx, apperror.ErrInvalidID)
		return
	}

	category, err := h.useCase.GetCategory(ctx, req.ID)
	if err != nil {
		h.logger.Error("%s: h.useCase.GetCategory: %v", op, err)
		switch {
		case errors.Is(err, ErrNotFound):
			apperror.WriteError(ctx, apperror.ErrNotFound)
		default:
			apperror.WriteError(ctx, apperror.ErrInternal)
		}
		return
	}

	err = router.WriteJSON(ctx, http.StatusOK, gin.H{"category": category}, nil)
	if err != nil {
		h.logger.Error("%s: router.WriteJSON: %v", op, err)
		apperror.WriteError(ctx, apperror.ErrInternal)
		return
	}
}

func (h *Handler) CreateCategoryHandler(ctx *gin.Context) {
	const op = "CreateCategoryHandler"

	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindJSON(&req): %v", op, err)
		apperror.WriteBadRequestError(ctx, err, "Something is missing or was not sent correctly")
		return
	}

	_ = &Category{
		CategoryName:    req.CategoryName,
		ParentID:        req.ParentID,
		AttributeSchema: req.AttributeSchema,
	}

}
