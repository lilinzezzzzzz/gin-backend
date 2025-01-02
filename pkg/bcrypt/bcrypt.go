package bcrypt

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 对密码进行哈希加密。
func HashPassword(password string) (string, error) {
	// 生成bcrypt哈希，使用默认成本参数
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "HashPassword.bcrypt.GenerateFromPassword")
	}
	return string(hashedBytes), nil
}

// VerifyPassword 验证密码是否匹配。
func VerifyPassword(plainPassword, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return errors.Errorf("verify password err: %v", err)
	}
	return nil
}
