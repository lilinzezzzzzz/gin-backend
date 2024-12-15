package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/dao/mysql"
	"innoversepm-backend/internal/entity"
	"innoversepm-backend/pkg/bcrypt"
	"innoversepm-backend/pkg/ctxHelper"
	"innoversepm-backend/pkg/resp"
)

type AuthServ struct {
	userDao *mysql.UserDao
}

func NewAuthServ() *AuthServ {
	return &AuthServ{
		userDao: mysql.NewUserDao(),
	}
}

func (a *AuthServ) UserSessionData(ctx *gin.Context) (*entity.UserSessionData, error) {
	userData, err := ctxhelper.GetUserData(ctx)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (a *AuthServ) LoginByAccount(ctx *gin.Context, account string, password string) (*entity.UserSessionData, error) {
	user, err := a.userDao.GetUserByAccount(ctx, account)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.VerifyPassword(password, user.Password); err != nil {
		resp.UNAUTHORIZED(ctx, err.Error())
		return nil, err
	}

	return &entity.UserSessionData{}, nil
}

func (a *AuthServ) LoginByPhone(ctx *gin.Context, phone string) (*entity.UserSessionData, error) {
	return nil, nil
}

func (a *AuthServ) LoginOut(ctx *gin.Context) error {
	resp.BadRequest(ctx, "not implement")
	return errors.New("not implement")
}
