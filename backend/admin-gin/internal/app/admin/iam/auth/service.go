package auth

import (
	"context"
	"errors"
	"mall-api/internal/pkg/jwt"
	"mall-api/internal/pkg/uuid"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type service interface {
	register(req *registerReq) error                                     // 注册
	login(ctx context.Context, req *loginReq) (*loginRes, error)         // 登录
	refresh(ctx context.Context, refreshToken string) (*loginRes, error) // 刷新 token
	logout(ctx context.Context, refreshToken string) (int, error)        // 注销
}

type svc struct {
	repo repository
	jt   *jwt.JWT
}

func newService(repo repository, jt *jwt.JWT) service {
	return &svc{
		repo: repo,
		jt:   jt,
	}
}

// 登陆
func (s *svc) login(ctx context.Context, req *loginReq) (*loginRes, error) {
	// 1. 查找用户
	account, err := s.repo.findUserByName(req.Username)
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
	if err := s.repo.setRefreshToken(ctx, account.UID, tokenPair.RefreshToken, s.jt.GetRefreshExpire()); err != nil {
		return nil, err
	}

	return &loginRes{
		UID:          account.UID,
		AccessToken:  tokenPair.AccessToken,
		ExpiresAt:    tokenPair.ExpiresAt,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

// 注册
func (s *svc) register(req *registerReq) error {

	// 1. 检查用户是否已存在
	if exist, err := s.repo.findUserIsExist(req.Username); err != nil {
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

	if err := s.repo.createUser(account); err != nil {
		return err
	}
	return nil
}

// 注销
func (s *svc) logout(ctx context.Context, refreshToken string) (int, error) {

	// 1. 解析获取 UID
	claims, err := s.jt.ParseToken(refreshToken, "refresh")
	if err != nil {
		return http.StatusUnauthorized, err
	}

	// 2. 从jwt中提取UID
	uid := claims.UID

	// 3. 获取 redis 中的值
	redisRefreshToken, err := s.repo.getRefreshToken(ctx, uid)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	// 4. 与请求返回的 refresh token 比对
	if redisRefreshToken != refreshToken {
		return http.StatusUnauthorized, err
	}

	// 5. 从 redis 中删除 refresh token
	if err := s.repo.delRefreshToken(ctx, uid); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// 刷新token
func (s *svc) refresh(ctx context.Context, refreshToken string) (*loginRes, error) {

	// 1. 解析 refresh_token，静态校验
	claims, err := s.jt.ParseToken(refreshToken, "refresh")
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
	savedRefreshToken, err2 := s.repo.getRefreshToken(ctx, uid)
	if err2 != nil {
		return nil, err2
	}

	// 5. redis 取出来的token和前端传的 refresh token进行对比
	if savedRefreshToken != refreshToken {
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
	if err := s.repo.setRefreshToken(ctx, uid, newRefresh, remaining); err != nil {
		return nil, err
	}

	// 8. 返回 签发的token信息
	return &loginRes{
		UID:         uid,
		AccessToken: newAccess,
		ExpiresAt:   time.Now().Add(s.jt.GetAccessExpire()).Unix(),
	}, nil
}
