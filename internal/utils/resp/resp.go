package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 定义统一的响应结构
type Response struct {
	Code    string `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type Items struct {
	Total int64 `json:"total"`
	Items any   `json:"items"`
}

// Success 返回成功响应
func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Response{
		Code:    "200",
		Data:    data,
		Message: "",
	})
	return
}

// Failed 返回失败响应
func Failed(ctx *gin.Context, code string, message string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    nil,
		Message: message,
	})
	return
}

// Error 返回错误响应
func Error(ctx *gin.Context, httpStatus int, code string, message string) {
	ctx.JSON(httpStatus, Response{
		Code:    code,
		Data:    nil,
		Message: message,
	})
	ctx.Abort()
	return // 在这里中止请求
}

// BadRequest 返回 400 错误
func BadRequest(ctx *gin.Context, message string) {
	Error(ctx, http.StatusBadRequest, "400", message)
	return
}

// UNAUTHORIZED 返回 401 错误
func UNAUTHORIZED(ctx *gin.Context, message string) {
	Error(ctx, http.StatusUnauthorized, "401", message)
	return
}

// Forbidden 返回 403 错误
func Forbidden(ctx *gin.Context, message string) {
	Error(ctx, http.StatusForbidden, "403", message)
	return
}

// NotFound 返回 404 错误
func NotFound(ctx *gin.Context, message string) {
	Error(ctx, http.StatusNotFound, "404", message)
	return
}

// UnprocessableEntity 返回 422 错误
func UnprocessableEntity(ctx *gin.Context, message string) {
	Error(ctx, http.StatusUnprocessableEntity, "422", message)
	return
}

// InternalServerError 返回 500 错误
func InternalServerError(ctx *gin.Context, message string) {
	Error(ctx, http.StatusInternalServerError, "500", message)
	return
}
