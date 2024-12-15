package entity

import (
	"innoversepm-backend/internal/models"
	"time"
)

// UserDTO 用于传输 User 数据的结构体
type UserDTO struct {
	ID          uint    `json:"id"`            // 用户 ID
	Account     string  `json:"account"`       // 账号
	Username    string  `json:"username"`      // 用户名
	Phone       string  `json:"phone"`         // 电话号码
	Password    string  `json:"password"`      // 密码（通常不返回给前端）
	AvatarURL   string  `json:"avatar_url"`    // 头像 URL
	Status      string  `json:"status"`        // 状态
	Email       string  `json:"email"`         // 邮箱
	Company     string  `json:"company"`       // 公司
	Department  string  `json:"department"`    // 部门
	Role        string  `json:"role"`          // 角色
	Channel     string  `json:"channel"`       // 渠道
	Category    string  `json:"category"`      // 分类
	VIPLevel    string  `json:"vip_level"`     // VIP 等级
	LastLoginAt *string `json:"last_login_at"` // 上次登录时间
	CreatedAt   string  `json:"created_at"`    // 创建时间
	UpdatedAt   string  `json:"updated_at"`    // 更新时间
	DeletedAt   *string `json:"deleted_at"`    // 删除时间
}

// ToUserDTO 将 models.User 转换为 entity.UserDTO
// ToUserDTO 将 models.User 转换为 entity.UserDTO，所有时间字段转换为 ISO 8601 格式
func ToUserDTO(user models.User) UserDTO {
	var lastLoginAt, deletedAt *string

	// 转换 LastLoginAt 为 ISO 8601 格式
	if user.LastLoginAt != nil {
		formatted := user.LastLoginAt.Format(time.RFC3339) // 使用 RFC3339 格式，即 ISO 8601
		lastLoginAt = &formatted
	}

	// 转换 DeletedAt 为 ISO 8601 格式
	if user.DeletedAt.Valid {
		formatted := user.DeletedAt.Time.Format(time.RFC3339)
		deletedAt = &formatted
	}

	return UserDTO{
		ID:          user.ID,
		Account:     user.Account,
		Username:    user.Username,
		Phone:       user.Phone,
		AvatarURL:   user.AvatarURL,
		Status:      user.Status,
		Email:       user.Email,
		Company:     user.Company,
		Department:  user.Department,
		Role:        user.Role,
		Channel:     user.Channel,
		Category:    user.Category,
		VIPLevel:    user.VIPLevel,
		LastLoginAt: lastLoginAt,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   user.UpdatedAt.Format(time.RFC3339),
		DeletedAt:   deletedAt,
	}
}
