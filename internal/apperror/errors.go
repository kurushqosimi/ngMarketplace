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

var (
	ErrInvalidID = &ErrorResponse{
		Status:  http.StatusBadRequest,
		Code:    "INVALID_ID",
		Error:   "Invalid ID parameter",
		Details: "The provided ID is invalid or malformed",
	}
	ErrNotFound = &ErrorResponse{
		Status:  http.StatusNotFound,
		Code:    "NOT_FOUND",
		Error:   "Resource not found",
		Details: "The requested resource was not found",
	}
	ErrInternal = &ErrorResponse{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_ERROR",
		Error:   "Internal server error",
		Details: "An unexpected error occurred",
	}
)

func WriteBadRequestError(ctx *gin.Context, err error, details string) {
	errResp := &ErrorResponse{
		Status:  http.StatusBadRequest,
		Code:    "BAD_REQUEST",
		Error:   err.Error(),
		Details: details,
	}
	WriteError(ctx, errResp)
}

func WriteError(ctx *gin.Context, errResp *ErrorResponse) {
	err := router.WriteJSON(ctx, errResp.Status, gin.H{"error-response": errResp}, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error-response": errResp})
	}
}
