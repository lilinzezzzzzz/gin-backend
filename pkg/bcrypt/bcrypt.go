package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 对密码进行哈希加密。
func HashPassword(password string) (string, error) {
	// 生成bcrypt哈希，使用默认成本参数
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// VerifyPassword 验证密码是否匹配。
func VerifyPassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
