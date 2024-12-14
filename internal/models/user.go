package models

import (
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

// User 模型
type User struct {
	gorm.Model             // 包含 ID、CreatedAt、UpdatedAt、DeletedAt 字段
	Account     string     `gorm:"size:20"`              // 对应 account 字段，字符串长度为 20
	Username    string     `gorm:"size:64"`              // 对应 username 字段，字符串长度为 64
	Phone       string     `gorm:"size:11"`              // 对应 phone 字段，字符串长度为 11
	AvatarURL   string     `gorm:"size:255"`             // 对应 avatar_url 字段，字符串长度为 255
	Status      string     `gorm:"size:16"`              // 对应 status 字段，字符串长度为 16
	Email       string     `gorm:"size:255"`             // 对应 email 字段，字符串长度为 255
	Company     string     `gorm:"size:16"`              // 对应 company 字段，字符串长度为 16
	Department  string     `gorm:"size:16"`              // 对应 department 字段，字符串长度为 16
	Role        string     `gorm:"size:16"`              // 对应 role 字段，字符串长度为 16
	Channel     string     `gorm:"size:16"`              // 对应 channel 字段，字符串长度为 16
	Category    string     `gorm:"size:16"`              // 对应 category 字段，字符串长度为 16
	VIPLevel    string     `gorm:"size:16"`              // 对应 vip_level 字段，字符串长度为 16
	LastLoginAt *time.Time `gorm:"column:last_login_at"` // 对应 last_login_at 字段，时间类型
}
