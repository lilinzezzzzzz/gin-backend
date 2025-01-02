package dao

import (
	"github.com/gin-gonic/gin"
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
	db     *gorm.DB
	logger func(ctx *gin.Context) *logrus.Entry
}

// NewUserDao 创建 UserDao 实例
func NewUserDao() *UserDao {
	return &UserDao{
		db:     infra.DB,
		logger: logger.Logger,
	}
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(ctx *gin.Context, user *models.User) error {
	if err := dao.db.Create(user).Error; err != nil {
		dao.logger(ctx).Errorf("Error creating user: %v", err)
		return err
	}
	return nil
}

// GetUserByID 根据 ID 查询用户
func (dao *UserDao) GetUserByID(ctx *gin.Context, id uint) (*models.User, error) {
	var user models.User
	if err := dao.db.First(&user, id).Error; err != nil {
		dao.logger(ctx).Errorf("Error fetching user by ID: %v", err)
		return nil, err
	}
	return &user, nil
}

// GetUserByAccount 根据账号查询用户
func (dao *UserDao) GetUserByAccount(ctx *gin.Context, account string) (*models.User, error) {
	var user models.User
	if err := dao.db.Where("account = ?", account).First(&user).Error; err != nil {
		dao.logger(ctx).Errorf("fetching user by account: %+v", err)
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func (dao *UserDao) UpdateUser(ctx *gin.Context, user *models.User) error {
	if err := dao.db.Save(user).Error; err != nil {
		logger.Logger(ctx).Errorf("Error updating user: %+v", err)
		return err
	}
	return nil
}

// DeleteUserByID 软删除用户
func (dao *UserDao) DeleteUserByID(ctx *gin.Context, id uint) error {
	if err := dao.db.Delete(&models.User{}, id).Error; err != nil {
		dao.logger(ctx).Errorf("Error deleting user: %+v", err)
		return err
	}
	return nil
}

type ManagerUserDao struct {
	db     *gorm.DB
	logger func(ctx *gin.Context) *logrus.Entry
}

// NewManagerUserDao 创建 UserDao 实例
func NewManagerUserDao() *UserDao {
	return &UserDao{
		db:     infra.DB,
		logger: logger.Logger,
	}
}
