package apperror

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ngMarketplace/internal/transport/http/router"
)

// ErrorResponse представляет собой стандартный формат ответа api
type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Error   string `json:"error"`
	Details string `json:"details"`
}

func WriteBadRequestError(ctx *gin.Context, err error, details string) {
	errResp := &ErrorResponse{
		Status:  http.StatusBadRequest,
		Code:    "BAD_REQUEST",
		Error:   err.Error(),
		Details: details,
	}

	writeError(ctx, errResp)
}

func WriteInternalServerError(ctx *gin.Context, err error, details string) {
	errResp := &ErrorResponse{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_ERROR",
		Error:   err.Error(),
		Details: "faced an unexpected error",
	}

	if details != "" {
		errResp.Details = details
	}

	writeError(ctx, errResp)
}

func writeError(ctx *gin.Context, errResp *ErrorResponse) {
	err := router.WriteJSON(ctx, errResp.Status, gin.H{"error-response": errResp}, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error-response": errResp})
	}
}
