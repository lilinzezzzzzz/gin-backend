package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang-backend/internal/entity"
	"golang-backend/internal/services"
	"golang-backend/internal/utils/logger"
	"golang-backend/internal/utils/resp"
)

type AuthController struct {
	logger *logrus.Logger
	serv   *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		serv: services.NewAuthService(),
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
		resp.BadRequest(ctx, fmt.Sprintf("Invalid request parameters: %v", err))
		return
	}

	session, err := a.serv.LoginByAccount(ctx, loginReq.Account, loginReq.Password)
	if err != nil {
		resp.UNAUTHORIZED(ctx, err.Error())
		return
	}
	resp.Success(ctx, session)
}

// LoginOut 登出
func (a *AuthController) LoginOut(ctx *gin.Context) {
	err := a.serv.LoginOut(ctx)
	if err != nil {
		logger.Logger(ctx).Infof(fmt.Sprintf("LoginOut error: %v", err))
		return
	}

	resp.Success(ctx, 0)
}
