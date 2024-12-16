package services

import (
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/entity"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) UserDetail(ctx *gin.Context, userID int) (*entity.UserEntity, error) {
	return nil, nil
}
