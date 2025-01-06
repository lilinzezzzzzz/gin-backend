package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-backend/internal/services"
	"golang-backend/internal/utils/resp"
)

type HelloController struct {
	srv *services.HelloService
}

func NewHelloController() *HelloController {
	return &HelloController{
		srv: services.NewHelloService(),
	}
}

func (h *HelloController) Hello(ctx *gin.Context) {
	data, err := h.srv.Hello(ctx)
	if err != nil {
		resp.InternalServerError(ctx, err.Error())
		return
	}

	resp.Success(ctx, data)
	return
}
