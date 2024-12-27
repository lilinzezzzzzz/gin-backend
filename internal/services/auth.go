package services

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang-backend/internal/dao"
	"golang-backend/internal/entity"
	"golang-backend/internal/utils/ctxhelper"
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
	return nil, errors.New("not implement")
}

func (a *AuthService) LogOut(ctx *gin.Context) error {
	return errors.New("not implement")
}
