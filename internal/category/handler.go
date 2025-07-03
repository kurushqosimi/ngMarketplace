package category

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"ngMarketplace/internal/apperror"
	"ngMarketplace/internal/common"
	"ngMarketplace/internal/transport/http/router"
	"ngMarketplace/pkg/logger"
	"strconv"
)

const (
	categoriesURL         = "/categories"
	categoryURL           = "/categories/:id"
	categoriesByParentURL = "/categories/parent/:parent_id"
)

type UseCase interface {
	Create(ctx context.Context, category *Category) error
	GetCategory(ctx context.Context, categoryID int64) (*Category, error)
	UpdateCategory(ctx context.Context, categoryID int64, category *updateCategoryRequest) (*Category, error)
	DeleteCategory(ctx context.Context, categoryID int64) error
	GetCategories(ctx context.Context, filters getCategoriesRequest) ([]*Category, common.Metadata, error)
	GetCategoryByParentID(ctx context.Context, parentID string) ([]*Category, error)
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
	router.GET(categoryURL, h.showCategoryHandler)
	router.POST(categoriesURL, h.createCategoryHandler)
	router.PATCH(categoryURL, h.updateCategoryHandler)
	router.DELETE(categoryURL, h.deleteCategoryHandler)
	router.GET(categoriesURL, h.listCategoriesHandler)
	router.GET(categoriesByParentURL, h.getByParentIDHandler)
}

// CreateCategoryHandler creates a new category in the marketplace
func (h *Handler) createCategoryHandler(ctx *gin.Context) {
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
		case errors.Is(err, ErrCategoryValidationFailed):
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

	if err = router.WriteJSON(ctx, http.StatusCreated, gin.H{"category": category}, nil); err != nil {
		h.logger.Warn("%s: router.WriteJSON: %v", op, err)
		ctx.JSON(http.StatusCreated, gin.H{"category": category})
		return
	}
}

// ShowCategoryHandler gets category by id
func (h *Handler) showCategoryHandler(ctx *gin.Context) {
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

	if err = router.WriteJSON(ctx, http.StatusOK, gin.H{"category": category}, nil); err != nil {
		h.logger.Warn("%s: router.WriteJSON: %v", op, err)
		ctx.JSON(http.StatusOK, gin.H{"category": category})
		return
	}
}

// updateCategoryHandler updates category by id
func (h *Handler) updateCategoryHandler(ctx *gin.Context) {
	const op = "updateCategoryHandler"

	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindUri: %v", op, err)
		apperror.WriteBadRequestResponse(ctx, ErrInvalidID, "Provide correct category id")
		return
	}

	var input updateCategoryRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Error("%s: ctx.ShouldBindJSON: %v", op, err)
		apperror.WriteBadRequestResponse(ctx, ErrBindJSON, "Something is missing or was not sent correctly")
		return
	}

	category, err := h.useCase.UpdateCategory(ctx, req.ID, &input)
	if err != nil {
		h.logger.Error("%s: h.useCase.UpdateCategory: %v", op, err)
		switch {
		case errors.Is(err, ErrNotFound) || errors.Is(err, ErrNotFoundForUpdate):
			apperror.WriteNotFoundResponse(ctx, err, "Category you are seeking to update does not exist")
		case errors.Is(err, ErrCategoryValidationFailed):
			apperror.WriteBadRequestResponse(ctx, err, err.Error())
		default:
			apperror.WriteInternalErrResponse(ctx, err, "Unexpected error occurred")
		}
		return
	}

	if err = router.WriteJSON(ctx, http.StatusOK, gin.H{"category": category}, nil); err != nil {
		h.logger.Warn("%s: router.WriteJSON: %v", op, err)
		ctx.JSON(http.StatusOK, gin.H{"category": category})
		return
	}
}

// deleteCategoryHandler deletes category by id
func (h *Handler) deleteCategoryHandler(ctx *gin.Context) {
	const op = "deleteCategoryHandler"

	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindUri: %v", op, err)
		apperror.WriteBadRequestResponse(ctx, ErrInvalidID, "Provide correct category id")
		return
	}

	if err := h.useCase.DeleteCategory(ctx, req.ID); err != nil {
		h.logger.Error("%s: h.useCase.DeleteCategory: %v", op, err)
		switch {
		case errors.Is(err, ErrNotFoundForDelete):
			apperror.WriteNotFoundResponse(ctx, err, "Category you are seeking to delete does not exist")
		default:
			apperror.WriteInternalErrResponse(ctx, err, "Unexpected error occurred")
		}
		return
	}

	if err := router.WriteJSON(ctx, http.StatusOK, gin.H{"message": "category was successfully deleted"}, nil); err != nil {
		h.logger.Warn("%s: router.WriteJSON: %v", op, err)
		ctx.JSON(http.StatusOK, gin.H{"message": "category was successfully deleted"})
		return
	}
}

// listCategoriesHandler get categories
func (h *Handler) listCategoriesHandler(ctx *gin.Context) {
	const op = "listCategoriesHandler"

	var req getCategoriesRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindQuery: %v", op, err)
		apperror.WriteBadRequestResponse(ctx, err, "some filter was sent with incorrect type")
		return
	}

	categories, metadata, err := h.useCase.GetCategories(ctx, req)
	if err != nil {
		switch {
		case errors.Is(err, common.ErrFilterValidationFailed):
			apperror.WriteBadRequestResponse(ctx, err, "check filter parameters")
		default:
			apperror.WriteInternalErrResponse(ctx, err, "Unexpected error occurred")
		}
		return
	}

	if err := router.WriteJSON(ctx, http.StatusOK, gin.H{"categories": categories, "metadata": metadata}, nil); err != nil {
		h.logger.Warn("%s: router.WriteJSON: %v", op, err)
		ctx.JSON(http.StatusOK, gin.H{"categories": categories, "metadata": metadata})
		return
	}
}

// getByParentIDHandler returns a list of categories by parent_id
func (h *Handler) getByParentIDHandler(ctx *gin.Context) {
	const op = "getByParentIDHandler"

	var req getCategoryByParentIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindUri: %v", op, err)
		apperror.WriteBadRequestResponse(ctx, ErrInvalidID, "Provide correct category id")
		return
	}

	parentID := strconv.Itoa(int(req.ParentID))

	categories, err := h.useCase.GetCategoryByParentID(ctx, parentID)
	if err != nil {
		h.logger.Error("%s: h.useCase.GetCategoryByParentID: %v", op, err)
		apperror.WriteInternalErrResponse(ctx, err, "Unexpected error occurred")
		return
	}

	if err := router.WriteJSON(ctx, http.StatusOK, gin.H{"categories": categories}, nil); err != nil {
		h.logger.Warn("%s: router.WriteJSON: %v", op, err)
		ctx.JSON(http.StatusOK, gin.H{"categories": categories})
		return
	}
}
