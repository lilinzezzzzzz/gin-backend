package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 定义统一的响应结构
type Response struct {
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Success 返回成功响应
func Success(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    "200",
		Data:    data,
		Message: message,
	})
}

// Error 返回错误响应
func Error(ctx *gin.Context, code string, httpStatus int, message string) {
	ctx.JSON(httpStatus, Response{
		Code:    code,
		Data:    nil,
		Message: message,
	})
}

// BadRequest 返回 400 错误
func BadRequest(ctx *gin.Context, message string) {
	Error(ctx, "400", http.StatusBadRequest, message)
}

// UNAUTHORIZED 返回 401 错误
func UNAUTHORIZED(ctx *gin.Context, message string) {
	Error(ctx, "401", http.StatusUnauthorized, message)
}

// Forbidden 返回 403 错误
func Forbidden(ctx *gin.Context, message string) {
	Error(ctx, "403", http.StatusForbidden, message)
}

// NotFound 返回 404 错误
func NotFound(ctx *gin.Context, message string) {
	Error(ctx, "404", http.StatusNotFound, message)
}

// UnprocessableEntity 返回 422 错误
func UnprocessableEntity(ctx *gin.Context, message string) {
	Error(ctx, "422", http.StatusUnprocessableEntity, message)
}

// InternalServerError 返回 500 错误
func InternalServerError(ctx *gin.Context, message string) {
	Error(ctx, "500", http.StatusInternalServerError, message)
}
