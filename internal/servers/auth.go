package servers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang-backend/internal/entity"
	"golang-backend/internal/services"
	"golang-backend/internal/utils/logger"
	"golang-backend/internal/utils/resp"
)

type AuthServer struct {
	logger func(ctx *gin.Context) *logrus.Entry
	srv    *services.AuthService
}

func NewAuthServer() *AuthServer {
	return &AuthServer{
		srv:    services.NewAuthService(),
		logger: logger.Logger,
	}
}

func (a *AuthServer) UserSessionData(ctx *gin.Context) {
	userData, err := a.srv.UserSessionData(ctx)
	if err != nil {
		resp.UNAUTHORIZED(ctx, err.Error())
		return
	}

	resp.Success(ctx, userData)
}

func (a *AuthServer) ManagerLogin(ctx *gin.Context) {
	// 绑定请求体到 LoginRequest 结构体
	var loginReq entity.LoginRequest
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		resp.BadRequest(ctx, fmt.Sprintf("Invalid request parameters: %v", err))
		return
	}

	session, err := a.srv.LoginByAccount(ctx, loginReq.Account, loginReq.Password)
	if err != nil {
		resp.Failed(ctx, "", err.Error())
		return
	}
	resp.Success(ctx, session)
}

// LoginOut 登出
func (a *AuthServer) LoginOut(ctx *gin.Context) {
	err := a.srv.LogOut(ctx)
	if err != nil {
		resp.Failed(ctx, "", err.Error())
		return
	}

	resp.Success(ctx, 0)
}
