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

// AddUser 增加用户
func (u *UserController) AddUser(ctx *gin.Context) {
	err := u.srv.AddUser(ctx)
	if err != nil {
		return
	}
	resp.Success(ctx, nil)
}

func (u *UserController) UserList(ctx *gin.Context) {
	users, err := u.srv.UserList(ctx)
	if err != nil {
		return
	}
	resp.Success(ctx, users)
}

func (u *UserController) UserDetail(ctx *gin.Context) {
	user, err := u.srv.UserDetail(ctx, 0)
	if err != nil {
		return
	}
	resp.Success(ctx, user)
}
