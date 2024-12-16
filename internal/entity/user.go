package entity

// UserEntity 用于传输 User 数据的结构体
type UserEntity struct {
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
