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
	logger func(ctx *gin.Context) *logrus.Entry
	srv    *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		srv:    services.NewAuthService(),
		logger: logger.Logger,
	}
}

func (a *AuthController) AuthMe(ctx *gin.Context) {
	userData, err := a.srv.UserSessionData(ctx)
	if err != nil {
		logger.Logger(ctx).Errorf("AuthMe.UserSessionData err: %v", err)
		resp.UNAUTHORIZED(ctx, err.Error())
		return
	}

	resp.Success(ctx, userData)
}

func (a *AuthController) UserLogin(ctx *gin.Context) {
	// 绑定请求体到 LoginRequest 结构体
	var req entity.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(ctx, fmt.Sprintf("Invalid request parameters: %v", err))
		return
	}

	session, err := a.srv.LoginByAccount(ctx, req.Account, req.Password)
	if err != nil {
		logger.Logger(ctx).Errorf("UserLogin.LoginByAccount err: %v", err)
		resp.InternalServerError(ctx, err.Error())
		return
	}
	resp.Success(ctx, session)
}

// UserLoginOut 登出
func (a *AuthController) UserLoginOut(ctx *gin.Context) {
	err := a.srv.LogOut(ctx)
	if err != nil {
		logger.Logger(ctx).Errorf("AuthController.UserLoginOut err: %v", err)
		resp.InternalServerError(ctx, err.Error())
		return
	}

	resp.Success(ctx, nil)
}
