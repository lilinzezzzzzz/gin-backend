// Package xsignature
package xsignature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/setting"
	"innoversepm-backend/pkg/logger"
	"sort"
	"strconv"
	"strings"
	"time"
)

// SignatureSrv 提供HMAC签名与验证相关的功能
type SignatureSrv struct {
	secretKey          []byte
	hashAlgorithm      string
	timestampTolerance int64
}

// NewSignatureSrv 创建一个HMACSigner实例
func NewSignatureSrv(cfg *setting.AppConfig) *SignatureSrv {
	return &SignatureSrv{
		secretKey:          []byte(cfg.App.SecretKey),
		hashAlgorithm:      "sha256",
		timestampTolerance: 18000,
	}
}

// GenerateSignature 生成签名
func (s *SignatureSrv) GenerateSignature(ctx *gin.Context, timestamp, nonce string) (string, error) {
	// 对数据进行排序，确保签名一致性
	data := map[string]string{
		"timestamp": timestamp,
		"nonce":     nonce,
	}
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for i, k := range keys {
		if i > 0 {
			sb.WriteString("&")
		}
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(data[k])
	}

	message := []byte(sb.String())

	// 使用HMAC + SHA256进行签名
	mac := hmac.New(sha256.New, s.secretKey)
	mac.Write(message)
	signature := hex.EncodeToString(mac.Sum(nil))

	return signature, nil
}

// VerifySignature 验证签名
func (s *SignatureSrv) VerifySignature(ctx *gin.Context, timestamp, signature, nonce string) bool {
	expectedSignature, err := s.GenerateSignature(ctx, timestamp, nonce)
	if err != nil {
		logger.Logger.Printf("VerifySignature generate expected signature error: %v", err)
		return false
	}

	// 使用hmac.Equal进行常量时间比较
	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}

// IsTimestampValid 验证时间戳是否有效
func (s *SignatureSrv) IsTimestampValid(ctx *gin.Context, requestTimeStr string) (bool, error) {
	requestTime, err := strconv.ParseInt(requestTimeStr, 10, 64)
	if err != nil {
		logger.Logger.Printf("IsTimestampValid failed: %v", err)
		return false, err
	}

	currentTime := time.Now().Unix()
	if (currentTime - requestTime) > s.timestampTolerance {
		logger.Logger.Printf("invalid timestamp, request_time: %d, current_time: %d", requestTime, currentTime)
		return false, nil
	}

	return true, nil
}

// VerifySignatureData 将原先独立的VerifySignatureFunc逻辑整合到HMACSigner的方法中
func (s *SignatureSrv) VerifySignatureData(ctx *gin.Context, xSignature, xTimestamp, xNonce string) (bool, error) {
	// 检查时间戳，防止重放攻击
	valid, err := s.IsTimestampValid(ctx, xTimestamp)
	if err != nil {
		return false, err
	}
	if !valid {
		logger.Logger.Printf("invalid timestamp: %s", xTimestamp)
		return false, nil
	}

	// 检查签名是否有效
	if !s.VerifySignature(ctx, xTimestamp, xSignature, xNonce) {
		logger.Logger.Printf("invalid signature: timestamp: %s, nonce: %s, signature: %s", xTimestamp, xNonce, xSignature)
		return false, nil
	}

	return true, nil
}
