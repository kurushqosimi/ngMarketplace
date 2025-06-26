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
	unacceptableCode       = "UNACCEPTABLE"
	notFoundCode           = "NOT_FOUND"
)

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Error   string `json:"error"`
	Details string `json:"details"`
}

// WriteBadRequestResponse - answers with bad request status (400)
func WriteBadRequestResponse(ctx *gin.Context, err error, details string) {
	if details == "" {
		details = "Something wrong with what you have sent"
	}

	errResp := &ErrorResponse{
		Status:  http.StatusBadRequest,
		Code:    badRequestCode,
		Error:   err.Error(),
		Details: details,
	}

	writeError(ctx, errResp)
}

func WriteNotFoundResponse(ctx *gin.Context, err error, details string) {
	if details == "" {
		details = "Not found"
	}

	errResp := &ErrorResponse{
		Status:  http.StatusNotFound,
		Code:    notFoundCode,
		Error:   err.Error(),
		Details: details,
	}

	writeError(ctx, errResp)
}

// WriteNotAcceptableResponse - answers with not acceptable status (406)
func WriteNotAcceptableResponse(ctx *gin.Context, err error, details string) {
	if details == "" {
		details = "Unacceptable entity was sent"
	}

	errResp := &ErrorResponse{
		Status:  http.StatusNotAcceptable,
		Code:    unacceptableCode,
		Error:   err.Error(),
		Details: details,
	}

	writeError(ctx, errResp)
}

// WriteConflictResponse - answers with conflict status (409)
func WriteConflictResponse(ctx *gin.Context, err error, details string) {
	if details == "" {
		details = "Something is conflicting with the actual situation"
	}

	errResp := &ErrorResponse{
		Status:  http.StatusConflict,
		Code:    conflictCode,
		Error:   err.Error(),
		Details: details,
	}

	writeError(ctx, errResp)
}

// WriteInternalErrResponse - answers with internal server error status (500)
func WriteInternalErrResponse(ctx *gin.Context, err error, details string) {
	if details == "" {
		details = "Internal server error"
	}

	errResp := &ErrorResponse{
		Status:  http.StatusInternalServerError,
		Code:    serverErrorCode,
		Error:   err.Error(),
		Details: details,
	}

	writeError(ctx, errResp)
}

// WriteSrvUnResponse - answers with service unavailable status (503)
func WriteSrvUnResponse(ctx *gin.Context, err error, details string) {
	if details == "" {
		details = "Service unavailable at the moment"
	}

	errResp := &ErrorResponse{
		Status:  http.StatusServiceUnavailable,
		Code:    serviceUnavailableCode,
		Error:   err.Error(),
		Details: details,
	}

	writeError(ctx, errResp)
}

func writeError(ctx *gin.Context, errResp *ErrorResponse) {
	err := router.WriteJSON(ctx, errResp.Status, gin.H{"error-response": errResp}, nil)
	if err != nil {
		ctx.JSON(errResp.Status, gin.H{"error-response": errResp})
	}
}
