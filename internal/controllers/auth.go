package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"innoversepm-backend/internal/services"
	"innoversepm-backend/pkg/logger"
	"net/http"
)

type AuthController struct {
	logger *logrus.Logger
	serv   *services.AuthServ
}

func NewAuthController() *AuthController {
	return &AuthController{
		logger: logger.Logger,
		serv:   services.NewAuthServ(),
	}
}

func (a *AuthController) Login(ctx *gin.Context) {
	userData, err := a.serv.LoginByAccount(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": userData})
}
