package converter

import (
	"github.com/jinzhu/copier"
	"golang-backend/internal/entity"
	"golang-backend/internal/models"
	"time"
)

// UserToEntity 将 models.User 转换为 entity.UserEntity，所有时间字段转换为 ISO 8601 格式
// UserToEntity 将 models.User 转换为 entity.UserEntity
func UserToEntity(m *models.User) *entity.UserEntity {
	userEntity := &entity.UserEntity{}

	// 使用 copier 复制相同字段
	err := copier.Copy(userEntity, m)
	if err != nil {
		return nil
	}

	// 手动处理时间字段转换为 ISO 8601 格式
	if m.LastLoginAt != nil {
		formatted := m.LastLoginAt.Format(time.RFC3339)
		userEntity.LastLoginAt = &formatted
	}

	if m.DeletedAt.Valid {
		formatted := m.DeletedAt.Time.Format(time.RFC3339)
		userEntity.DeletedAt = &formatted
	}

	userEntity.CreatedAt = m.CreatedAt.Format(time.RFC3339)
	userEntity.UpdatedAt = m.UpdatedAt.Format(time.RFC3339)

	return userEntity
}

// UserToModel 将 entity.UserEntity 转换为 models.User
func UserToModel(e *entity.UserEntity) *models.User {
	userModel := &models.User{}

	// 使用 copier 复制相同字段
	err := copier.Copy(userModel, e)
	if err != nil {
		return nil
	}

	// 手动处理时间字段，如果需要的话
	// 如果 LastLoginAt 存在，则解析为 time.Time
	if e.LastLoginAt != nil {
		if parsedTime, err := time.Parse(time.RFC3339, *e.LastLoginAt); err == nil {
			userModel.LastLoginAt = &parsedTime
		}
	}

	return userModel
}
