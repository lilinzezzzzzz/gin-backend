package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang-backend/internal/infra"
	"golang-backend/internal/models"
	"golang-backend/internal/utils/logger"

	"gorm.io/gorm"
)

type User interface {
	CreateUser(ctx *gin.Context, user *models.User) error
	GetUserByID(ctx *gin.Context, id uint) (*models.User, error)
}

// UserDao 结构体
type UserDao struct {
	db     func(ctx *gin.Context) *gorm.DB
	logger func(ctx *gin.Context) *logrus.Entry
}

// NewUserDao 创建 UserDao 实例
func NewUserDao() *UserDao {
	return &UserDao{
		db:     infra.NewDB,
		logger: logger.Logger,
	}
}

// CreateUser 创建用户
func (u *UserDao) CreateUser(ctx *gin.Context, user *models.User) error {
	if err := u.db(ctx).Create(user).Error; err != nil {
		u.logger(ctx).Errorf("Error creating user: %v", err)
		return err
	}
	return nil
}

// GetUserByID 根据 ID 查询用户
func (u *UserDao) GetUserByID(ctx *gin.Context, id uint) (*models.User, error) {
	var user models.User
	if err := u.db(ctx).First(&user, id).Error; err != nil {
		u.logger(ctx).Errorf("Error fetching user by ID: %v", err)
		return nil, err
	}
	return &user, nil
}

// GetUserByAccount 根据账号查询用户
func (u *UserDao) GetUserByAccount(ctx *gin.Context, account string) (*models.User, error) {
	var user models.User
	if err := u.db(ctx).Where("account = ?", account).First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "UserDao.GetUserByAccount")
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func (u *UserDao) UpdateUser(ctx *gin.Context, user *models.User) error {
	if err := u.db(ctx).Save(user).Error; err != nil {
		logger.Logger(ctx).Errorf("Error updating user: %+v", err)
		return err
	}
	return nil
}

// DeleteUserByID 软删除用户
func (u *UserDao) DeleteUserByID(ctx *gin.Context, id uint) error {
	if err := u.db(ctx).Delete(&models.User{}, id).Error; err != nil {
		u.logger(ctx).Errorf("Error deleting user: %+v", err)
		return err
	}
	return nil
}
