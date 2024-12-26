package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-backend/internal/dao"
	"golang-backend/internal/entity"
	"golang-backend/internal/utils/ctxhelper"
	"golang-backend/internal/utils/resp"
	"golang-backend/pkg/bcrypt"
)

type AuthService struct {
	userDao *dao.UserDao
}

func NewAuthService() *AuthService {
	return &AuthService{
		userDao: dao.NewUserDao(),
	}
}

func (a *AuthService) UserSessionData(ctx *gin.Context) (*entity.UserSessionData, error) {
	userData, err := ctxhelper.GetUserData(ctx)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (a *AuthService) LoginByAccount(ctx *gin.Context, account string, password string) (*entity.UserSessionData, error) {
	user, err := a.userDao.GetUserByAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.VerifyPassword(password, user.Password); err != nil {
		return nil, err
	}

	return &entity.UserSessionData{}, nil
}

func (a *AuthService) LoginByPhone(ctx *gin.Context, phone string) (*entity.UserSessionData, error) {
	return nil, nil
}

func (a *AuthService) LoginOut(ctx *gin.Context) error {
	resp.BadRequest(ctx, "not implement")
	return errors.New("not implement")
}
