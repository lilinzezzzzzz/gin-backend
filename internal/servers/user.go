package servers

import (
	"github.com/gin-gonic/gin"
	"golang-backend/internal/services"
	"golang-backend/internal/utils/resp"
)

type UserServer struct {
	srv *services.UserService
}

func NewUserServer() *UserServer {
	return &UserServer{
		srv: services.NewUserService(),
	}
}

// AddUser 增加用户
func (u *UserServer) AddUser(ctx *gin.Context) {
	err := u.srv.AddUser(ctx)
	if err != nil {
		return
	}
	resp.Success(ctx, nil)
}

func (u *UserServer) UserList(ctx *gin.Context) {
	users, err := u.srv.UserList(ctx)
	if err != nil {
		return
	}
	resp.Success(ctx, users)
}

func (u *UserServer) UserDetail(ctx *gin.Context) {
	user, err := u.srv.UserDetail(ctx, 0)
	if err != nil {
		return
	}
	resp.Success(ctx, user)
}
