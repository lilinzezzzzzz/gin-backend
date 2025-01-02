package services

import "github.com/gin-gonic/gin"

type HelloService struct {
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (h *HelloService) Hello(ctx *gin.Context) (string, error) {
	return "hello", nil
}
