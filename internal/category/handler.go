package category

import (
	"github.com/gin-gonic/gin"
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
}

type Handler struct {
	UseCase UseCase
	logger  *logger.Logger
}

func (h Handler) Register(router *gin.Engine) {
	router.GET(categoryURL, h.GetCategoryHandler)
}

func (h Handler) GetCategoryHandler(c *gin.Context) {

}
