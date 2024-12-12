package services

import (
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/entity"
)

type AuthServ struct {
}

func NewAuthServ() *AuthServ {
	return &AuthServ{}
}

func (a *AuthServ) LoginByAccount(ctx *gin.Context) (*entity.UserSessionData, error) {
	userData := entity.GetUserData(ctx)
	return userData, nil
}
