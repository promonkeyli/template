package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	// JWT密钥，建议从配置文件读取
	jwtSecret = []byte("your-secret-key")

	// Token过期时间配置
	AccessTokenExpire  = time.Hour * 2      // Access Token 2小时过期
	RefreshTokenExpire = time.Hour * 24 * 7 // Refresh Token 7天过期

	// 错误定义
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

// Claims JWT声明结构
type Claims struct {
	UID       string `json:"uid"`        // 全局唯一用户标识
	TokenType string `json:"token_type"` // "access" 或 "refresh"
	jwt.RegisteredClaims
}

// TokenPair 包含access token和refresh token
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"` // access token过期时间戳
}

// GenerateAccessToken 生成Access Token
func GenerateAccessToken(uid string) (string, error) {
	return generateToken(uid, "access", AccessTokenExpire)
}

// GenerateRefreshToken 生成Refresh Token
func GenerateRefreshToken(uid string) (string, error) {
	return generateToken(uid, "refresh", RefreshTokenExpire)
}

// GenerateTokenPair 生成一对Token (access + refresh)
func GenerateTokenPair(uid string) (*TokenPair, error) {
	accessToken, err := GenerateAccessToken(uid)
	if err != nil {
		return nil, err
	}

	refreshToken, err := GenerateRefreshToken(uid)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(AccessTokenExpire).Unix(),
	}, nil
}

// generateToken 生成JWT token的通用函数
func generateToken(uid string, tokenType string, expireDuration time.Duration) (string, error) {
	claims := Claims{
		UID:       uid,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析JWT token (通用)
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// ParseAccessToken 解析并验证Access Token
func ParseAccessToken(tokenString string) (*Claims, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != "access" {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// ParseRefreshToken 解析并验证Refresh Token
func ParseRefreshToken(tokenString string) (*Claims, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != "refresh" {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// ValidateToken 验证token是否有效 (鉴权函数)
func ValidateToken(tokenString string) (*Claims, error) {
	return ParseAccessToken(tokenString)
}

// RefreshAccessToken 使用refresh token刷新access token
func RefreshAccessToken(refreshTokenString string) (*TokenPair, error) {
	claims, err := ParseRefreshToken(refreshTokenString)
	if err != nil {
		return nil, err
	}

	// 使用refresh token中的信息生成新的token pair
	return GenerateTokenPair(claims.UID)
}

// SetJWTSecret 设置JWT密钥 (可在初始化时调用)
func SetJWTSecret(secret string) {
	jwtSecret = []byte(secret)
}

// SetTokenExpire 设置Token过期时间 (可在初始化时调用)
func SetTokenExpire(accessExpire, refreshExpire time.Duration) {
	AccessTokenExpire = accessExpire
	RefreshTokenExpire = refreshExpire
}
