package apperror

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ngMarketplace/internal/transport/http/router"
)

// Codes for error response
var (
	badRequestCode         = "BAD_REQUEST"
	conflictCode           = "CONFLICT"
	serviceUnavailableCode = "SERVICE_UNAVAILABLE"
	serverErrorCode        = "INTERNAL_SERVER_ERROR"
)

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Error   string `json:"error"`
	Details string `json:"details"`
}

func WriteBadRequestResponse(ctx *gin.Context, err error, details string) {
	if details == "" {
		details = "Something wrong you have sent"
	}

	errResp := &ErrorResponse{
		Status:  http.StatusBadRequest,
		Code:    badRequestCode,
		Error:   err.Error(),
		Details: details,
	}

	writeError(ctx, errResp)
}

func WriteConflictResponse(ctx *gin.Context, err error, details string) {
	errResp := &ErrorResponse{
		Status:  http.StatusConflict,
		Code:    conflictCode,
		Error:   err.Error(),
		Details: details,
	}

	writeError(ctx, errResp)
}

func WriteSrvUnResponse(ctx *gin.Context, err error, details string) {
	errResp := &ErrorResponse{
		Status:  http.StatusServiceUnavailable,
		Code:    serviceUnavailableCode,
		Error:   err.Error(),
		Details: details,
	}

	writeError(ctx, errResp)
}

func WriteInternalErrResponse(ctx *gin.Context, err error, details string) {
	errResp := &ErrorResponse{
		Status:  http.StatusInternalServerError,
		Code:    serverErrorCode,
		Error:   err.Error(),
		Details: details,
	}

	writeError(ctx, errResp)
}

func writeError(ctx *gin.Context, errResp *ErrorResponse) {
	err := router.WriteJSON(ctx, errResp.Status, gin.H{"error-response": errResp}, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error-response": errResp})
	}
}
