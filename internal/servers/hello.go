package servers

import (
	"github.com/gin-gonic/gin"
	"golang-backend/internal/services"
	"golang-backend/internal/utils/resp"
)

type HelloServer struct {
	srv *services.HelloService
}

func NewHelloServer() *HelloServer {
	return &HelloServer{
		srv: services.NewHelloService(),
	}
}

func (h *HelloServer) Hello(ctx *gin.Context) {
	data, err := h.srv.Hello(ctx)
	if err != nil {
		resp.InternalServerError(ctx, err.Error())
		return
	}

	resp.Success(ctx, data)
	return
}
