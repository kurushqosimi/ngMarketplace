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
	Create(ctx context.Context, category *Category) error
	GetCategory(ctx context.Context, categoryID int64) (*Category, error)
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
	router.GET(categoryURL, h.ShowCategoryHandler)
	router.POST(categoriesURL, h.CreateCategoryHandler)
}

// CreateCategoryHandler creates a new category in the marketplace
func (h *Handler) CreateCategoryHandler(ctx *gin.Context) {
	const op = "CreateCategoryHandler"

	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindJSON: %v", op, err)
		apperror.WriteBadRequestResponse(ctx, ErrBindJSON, "Something is missing or was not sent correctly")
		return
	}

	category := &Category{
		CategoryName:    req.CategoryName,
		ParentID:        req.ParentID,
		Language:        req.Language,
		AttributeSchema: req.AttributeSchema,
	}

	err := h.useCase.Create(ctx, category)
	if err != nil {
		h.logger.Error("%s: h.useCase.Create: %v", op, err)
		switch {
		case errors.Is(err, ErrValidationFailed):
			apperror.WriteBadRequestResponse(ctx, err, err.Error())
		case errors.Is(err, ErrDuplicateCategory):
			apperror.WriteConflictResponse(ctx, err, "Category with this name, parent, and language already exists")
		case errors.Is(err, ErrInvalidParentID):
			apperror.WriteBadRequestResponse(ctx, err, "Parent category does not exists")
		case errors.Is(err, ErrConnectionFailed):
			apperror.WriteSrvUnResponse(ctx, err, "Database connection failed")
		default:
			apperror.WriteInternalErrResponse(ctx, err, "Internal server error")
		}
		return
	}

	err = router.WriteJSON(ctx, http.StatusCreated, gin.H{"category": category}, nil)
	if err != nil {
		h.logger.Warn("%s: router.WriteJSON: %v", op, err)
		ctx.JSON(http.StatusCreated, gin.H{"category": category})
		return
	}
}

// ShowCategoryHandler gets category by id
func (h *Handler) ShowCategoryHandler(ctx *gin.Context) {
	const op = "GetCategoryHandler"

	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindUri: %v", op, err)
		apperror.WriteBadRequestResponse(ctx, ErrInvalidID, "Provide correct category id")
		return
	}

	category, err := h.useCase.GetCategory(ctx, req.ID)
	if err != nil {
		h.logger.Error("%s: h.useCase.GetCategory: %v", op, err)
		switch {
		case errors.Is(err, ErrNotFound):
			apperror.WriteNotFoundResponse(ctx, err, "Category you are seeking does not exist")
		default:
			apperror.WriteInternalErrResponse(ctx, err, "Unexpected error occurred")
		}
		return
	}

	err = router.WriteJSON(ctx, http.StatusOK, gin.H{"category": category}, nil)
	if err != nil {
		h.logger.Warn("%s: router.WriteJSON: %v", op, err)
		ctx.JSON(http.StatusOK, gin.H{"category": category})
		return
	}
}
