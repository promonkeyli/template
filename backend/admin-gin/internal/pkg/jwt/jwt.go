package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("无效令牌") // 两种错误 http 业务 code 均返回401
	ErrExpiredToken = errors.New("令牌过期")
)

// Claims JWT声明结构
type Claims struct {
	UID       string `json:"uid"`        // 全局唯一用户标识
	TokenType string `json:"token_type"` // "access" 或 "refresh"
	jwt.RegisteredClaims
}

// TokenPair 包含 access token 和 refresh token
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"` // access token 过期时间戳
}

// JWT 处理引擎
type JWT struct {
	secret        []byte
	issuer        string
	accessExpire  time.Duration
	refreshExpire time.Duration
}

// New 创建 JWT 实例
func New(secret, issuer string, accessExpire, refreshExpire time.Duration) *JWT {
	return &JWT{
		secret:        []byte(secret),
		issuer:        issuer,
		accessExpire:  accessExpire,
		refreshExpire: refreshExpire,
	}
}

// GetAccessExpire 获取 Access Token 过期时间
func (j *JWT) GetAccessExpire() time.Duration {
	return j.accessExpire
}

// GetRefreshExpire 获取 Refresh Token 过期时间
func (j *JWT) GetRefreshExpire() time.Duration {
	return j.refreshExpire
}

// GenerateToken 签发指定类型的 Token (用于刷新场景，可指定自定义有效期)
func (j *JWT) GenerateToken(uid string, tokenType string, expireDuration time.Duration) (string, error) {
	return j.generateToken(uid, tokenType, expireDuration)
}

// GenerateTokenPair 生成一对 Token (access + refresh)
func (j *JWT) GenerateTokenPair(uid string) (*TokenPair, error) {
	accessToken, err := j.generateToken(uid, "access", j.accessExpire)
	if err != nil {
		return nil, err
	}

	refreshToken, err := j.generateToken(uid, "refresh", j.refreshExpire)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(j.accessExpire).Unix(),
	}, nil
}

// ParseToken 解析并验证 JWT token
// 如果 tokenType 不为空，则会额外校验声明中的 token_type 字段是否匹配 ("access" 或 "refresh")
func (j *JWT) ParseToken(tokenString string, tokenType string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	// 校验 Token 类型（如果指定了类型）
	if tokenType != "" && claims.TokenType != tokenType {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// token 生成通用函数
func (j *JWT) generateToken(uid string, tokenType string, expireDuration time.Duration) (string, error) {
	now := time.Now()
	claims := Claims{
		UID:       uid,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(now.Add(expireDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.secret)
}
