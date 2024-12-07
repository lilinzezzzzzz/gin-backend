// Package xjwt
package xjwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"innoversepm-backend/internal/setting"
	"innoversepm-backend/pkg/logger"
	"strings"
	"time"
)

// Claims 定义JWT的自定义声明
type Claims struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

// JwtService 提供JWT相关的功能
type JwtService struct {
	secretKey     string
	jwtAlgorithm  string
	expireMinutes int
}

// NewJWTService 创建一个新的JWTService实例
func NewJWTService(cfg *setting.AppConfig) *JwtService {
	return &JwtService{
		secretKey:     cfg.App.SecretKey,
		jwtAlgorithm:  "HS256",
		expireMinutes: 180,
	}
}

// CreateToken 创建一个新的JWT令牌
func (j *JwtService) CreateToken(c *gin.Context, userID int, userName string) (string, error) {
	expirationTime := time.Now().UTC().Add(time.Duration(j.expireMinutes) * time.Minute)

	claims := &Claims{
		UserID:   userID,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(j.jwtAlgorithm), claims)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken 验证JWT令牌并返回用户ID
func (j *JwtService) VerifyToken(c *gin.Context, tokenStr string) (int, bool) {
	if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
		logger.Logger.Warn("Token verification failed: token is empty or does not start with Bearer")
		return 0, false
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			logger.Logger.Warn("Token verification failed: token expired")
			return 0, false
		}
		logger.Logger.Warn("Token verification failed: invalid token")
		return 0, false
	}

	if !token.Valid {
		logger.Logger.Warn("Token verification failed: token is invalid")
		return 0, false
	}

	return claims.UserID, true
}
