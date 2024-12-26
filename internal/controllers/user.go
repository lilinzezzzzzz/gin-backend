package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-backend/internal/services"
	"golang-backend/internal/utils/resp"
)

type UserController struct {
	srv *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		srv: services.NewUserService(),
	}
}

func (u *UserController) UserDetail(ctx *gin.Context) {
	user, err := u.srv.UserDetail(ctx, 0)
	if err != nil {
		return
	}
	resp.Success(ctx, user)
}
