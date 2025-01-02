package services

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang-backend/internal/converter"
	"golang-backend/internal/core"
	"golang-backend/internal/dao"
	"golang-backend/internal/entity"
	"golang-backend/internal/utils/ctxhelper"
	"golang-backend/pkg/bcrypt"
)

type AuthService struct {
	userDao *dao.UserDao
	cache   *dao.Cache
}

func NewAuthService() *AuthService {
	return &AuthService{
		userDao: dao.NewUserDao(),
		cache:   dao.NewCache(),
	}
}

func (a *AuthService) UserSessionData(ctx *gin.Context) (*entity.UserSessionData, error) {
	userData, err := ctxhelper.GetUserData(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "UserSessionData.GetUserData")
	}
	return userData, nil
}

func (a *AuthService) LoginByAccount(ctx *gin.Context, account string, password string) (string, error) {
	user, err := a.userDao.GetUserByAccount(ctx, account)
	if err != nil {
		return "", errors.Wrap(err, "AuthService.LoginByAccount")
	}

	if err := bcrypt.VerifyPassword(password, user.Password); err != nil {
		return "", err
	}

	session := core.GenerateSession()
	userJson, err := converter.UserToJSONString(user)
	if err != nil {
		return "", err
	}

	// 设置session
	if err := a.cache.SetSession(ctx, session, userJson); err != nil {
		return "", err
	}

	if err := a.cache.SetSessionList(ctx, user.ID, session); err != nil {
		return "", err
	}

	return session, nil
}

func (a *AuthService) LoginByPhone(ctx *gin.Context, phone string) (*entity.UserSessionData, error) {
	return nil, errors.New("not implement")
}

func (a *AuthService) LogOut(ctx *gin.Context) error {
	return errors.New("not implement")
}
