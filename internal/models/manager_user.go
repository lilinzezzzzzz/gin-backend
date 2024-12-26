package models

import (
	"golang-backend/pkg/snowflake"
	"gorm.io/gorm"
	"time"
)

// ManageUser 模型
type ManageUser struct {
	gorm.Model        // 包含 ID、CreatedAt、UpdatedAt、DeletedAt 字段
	Account    string `gorm:"size:32"`  // 对应 account 字段，字符串长度为 32
	Username   string `gorm:"size:32"`  // 对应 username 字段，字符串长度为 32
	Password   string `gorm:"size:128"` // 对应 password 字段，字符串长度为 128
}

func (m *ManageUser) BeforeCreate(_ *gorm.DB) error {
	m.ID = snowflake.GenerateSnowflakeID()
	m.CreatedAt = time.Now().UTC()
	return nil
}

// BeforeUpdate 钩子，在更新记录前将 UpdatedAt 设置为 UTC 时间
func (m *ManageUser) BeforeUpdate(_ *gorm.DB) error {
	m.UpdatedAt = time.Now().UTC()
	return nil
}
