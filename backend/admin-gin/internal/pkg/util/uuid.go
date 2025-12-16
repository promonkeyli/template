package util

import (
	"github.com/google/uuid"
)

// GenerateUID 生成全局唯一的用户标识符
// 使用 UUID v4 算法,返回32位不带横线的字符串
func NewUUID() string {
	return uuid.New().String()[:32]
}
