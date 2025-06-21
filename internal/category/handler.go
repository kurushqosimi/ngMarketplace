package category

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
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
}

func (h *Handler) ShowCategoryHandler(ctx *gin.Context) {
	const op = "GetCategoryHandler"

	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		h.logger.Error("%s: ctx.ShouldBindUri: %v", op, err)
		WriteError(ctx, ErrInvalidID)
		return
	}

	category, err := h.useCase.GetCategory(ctx, req.ID)
	if err != nil {
		h.logger.Error("%s: h.useCase.GetCategory: %v", op, err)
		switch {
		case errors.Is(err, ErrNotFound):
			WriteError(ctx, ErrNotFound)
		default:
			WriteError(ctx, ErrInternal)
		}
		return
	}

	err = router.WriteJSON(ctx, http.StatusOK)
}
