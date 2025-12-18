package auth

import (
	"context"
	"errors"
	"mall-api/internal/pkg/jwt"
	"mall-api/internal/pkg/uuid"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(req *RegisterReq) error                                 // 注册
	Login(ctx context.Context, req *LoginReq) (*LoginRes, error)     // 登录
	Refresh(ctx context.Context, req *RefreshReq) (*LoginRes, error) // 刷新 token
}

type service struct {
	repo Repository
	jt   *jwt.JWT
}

func NewService(repo Repository, jt *jwt.JWT) Service {
	return &service{
		repo: repo,
		jt:   jt,
	}
}

// 注册
func (s *service) Register(req *RegisterReq) error {

	// 1. 检查用户是否已存在
	if exist, err := s.repo.FindUserIsExist(req.Username); err != nil {
		return err
	} else if exist {
		return errors.New("用户已经存在")
	}

	// 2. 生成全局唯一 UID
	uid := uuid.NewUUID()

	// 3. 加密密码（bcrypt hash）
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 4. 构建 auth 领域的最小账号信息（与 user 模块解耦）
	account := &account{
		UID:      uid,
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     "admin",
		IsActive: true,
	}

	if err := s.repo.CreateUser(account); err != nil {
		return err
	}
	return nil
}

// 登陆
func (s *service) Login(ctx context.Context, req *LoginReq) (*LoginRes, error) {
	// 1. 查找用户
	account, err := s.repo.FindUserByName(req.Username)
	if err != nil {
		return nil, err
	}

	// 2. 校验用户密码是否正确（对比 hash 与明文）
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("密码不正确")
	}

	// 3. 构建 token
	tokenPair, err := s.jt.GenerateTokenPair(account.UID)
	if err != nil {
		return nil, err
	}

	// 4. 将 refresh token 存储到 redis
	if err := s.repo.SetRefreshToken(ctx, account.UID, tokenPair.RefreshToken, s.jt.GetRefreshExpire()); err != nil {
		return nil, err
	}

	return &LoginRes{
		UID:          account.UID,
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresAt:    tokenPair.ExpiresAt,
	}, nil
}

// 刷新token
func (s *service) Refresh(ctx context.Context, req *RefreshReq) (*LoginRes, error) {

	// 1. 解析 refresh_token，静态校验
	claims, err := s.jt.ParseToken(req.RefreshToken, "refresh")
	if err != nil {
		return nil, err
	}

	// 2. 计算剩余有效期 (实现绝对过期时间，防止无限续期)
	remaining := time.Until(claims.ExpiresAt.Time)
	if remaining <= 0 {
		return nil, jwt.ErrExpiredToken
	}

	// 3. 从jwt中提取UID
	uid := claims.UID

	// 4. redis 值校验
	savedRefreshToken, err2 := s.repo.GetRefreshToken(ctx, uid)
	if err2 != nil {
		return nil, err2
	}

	// 5. redis 取出来的token和前端传的 refresh token进行对比
	if savedRefreshToken != req.RefreshToken {
		return nil, errors.New("刷新token无效")
	}

	// 6. 签发新的 Token 对
	// Access Token 总是给满额有效期
	newAccess, err := s.jt.GenerateToken(uid, "access", s.jt.GetAccessExpire())
	if err != nil {
		return nil, err
	}
	// Refresh Token 继承旧 Token 的剩余寿命
	newRefresh, err := s.jt.GenerateToken(uid, "refresh", remaining)
	if err != nil {
		return nil, err
	}

	// 7. 更新 Redis
	if err := s.repo.SetRefreshToken(ctx, uid, newRefresh, remaining); err != nil {
		return nil, err
	}

	// 8. 返回 签发的token信息
	return &LoginRes{
		UID:          uid,
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
		ExpiresAt:    time.Now().Add(s.jt.GetAccessExpire()).Unix(),
	}, nil
}
