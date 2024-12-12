package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"innoversepm-backend/internal/entity"
	"innoversepm-backend/internal/services"
	"innoversepm-backend/pkg/logger"
	"innoversepm-backend/pkg/resp"
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

func (a *AuthController) UserSessionData(ctx *gin.Context) {
	userData, err := a.serv.UserSessionData(ctx)
	if err != nil {
		resp.UNAUTHORIZED(ctx, err.Error())
		return
	}

	resp.Success(ctx, userData)
}

func (a *AuthController) ManagerLogin(ctx *gin.Context) {
	// 绑定请求体到 LoginRequest 结构体
	var loginReq entity.LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		resp.BadRequest(ctx, "Invalid request parameters")
		return
	}

	session, err := a.serv.LoginByAccount(ctx, "", "")
	if err != nil {
		resp.UNAUTHORIZED(ctx, err.Error())
		return
	}
	resp.Success(ctx, session)
}
