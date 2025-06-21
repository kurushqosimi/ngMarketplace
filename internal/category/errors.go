package category

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type envelope map[string]interface{}

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

func WriteError(ctx *gin.Context, err *ErrorResponse) {
	ctx.JSON(err.Status, envelope{"error-response": err})
}
