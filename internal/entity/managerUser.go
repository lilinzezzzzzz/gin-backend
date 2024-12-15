package entity

import "time"

// ManageUserDTO 用于传输 ManageUser 数据的结构体
type ManageUserDTO struct {
	ID        uint      `json:"id"`         // 用户 ID
	Account   string    `json:"account"`    // 账号
	Username  string    `json:"username"`   // 用户名
	Password  string    `json:"password"`   // 密码（通常不返回给前端）
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	DeletedAt time.Time `json:"deleted_at"` // 删除时间
}
