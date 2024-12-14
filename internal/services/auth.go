package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/entity"
	"innoversepm-backend/pkg/ctxHelper"
	"innoversepm-backend/pkg/resp"
)

type AuthServ struct {
}

func NewAuthServ() *AuthServ {
	return &AuthServ{}
}

func (a *AuthServ) UserSessionData(ctx *gin.Context) (*entity.UserSessionData, error) {
	userData, err := ctxHelper.GetUserData(ctx)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (a *AuthServ) LoginByAccount(ctx *gin.Context, account string, password string) (*entity.UserSessionData, error) {
	userData, err := ctxHelper.GetUserData(ctx)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (a *AuthServ) LoginByPhone(ctx *gin.Context, phone string) (*entity.UserSessionData, error) {
	return nil, nil
}

func (a *AuthServ) LoginOut(ctx *gin.Context) error {
	resp.BadRequest(ctx, "not implement")
	return errors.New("not implement")
}
