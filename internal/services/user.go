package services

import (
	"github.com/gin-gonic/gin"
	"golang-backend/internal/entity"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) UserDetail(ctx *gin.Context, userID int) (*entity.UserEntity, error) {
	return nil, nil
}

func (u *UserService) AddUser(ctx *gin.Context) interface{} {
	panic("not implemented")
}

func (u *UserService) UserList(ctx *gin.Context) (interface{}, interface{}) {
	panic("not implemented")
}
