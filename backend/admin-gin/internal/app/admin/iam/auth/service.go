package auth

import (
	"context"
	"errors"
	"mall-api/internal/pkg/util"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(req *RegisterReq) error                                 // 注册
	Login(req *LoginReq) (LoginRes, error)                           // 登录
	Refresh(ctx context.Context, req *RefreshReq) (*LoginRes, error) // 刷新 token
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Register(req *RegisterReq) error {

	// 1. 检查用户是否已存在
	if exist, err := s.repo.FindUserIsExist(req.Username); err != nil {
		return err
	} else if exist {
		return errors.New("用户已经存在")
	}

	// 2. 生成全局唯一 UID
	uid := util.NewUUID()

	// 3. 加密密码（bcrypt hash）
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 4. 构建 auth 领域的最小账号信息（与 user 模块解耦）
	account := &Account{
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

func (s *service) Login(req *LoginReq) (LoginRes, error) {
	// 1. 查找用户
	account, err := s.repo.FindUserByName(req.Username)
	if err != nil {
		return LoginRes{}, err
	}

	// 2. 校验用户密码是否正确（对比 hash 与明文）
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		return LoginRes{}, errors.New("密码不正确")
	}

	// 3. 构建 token
	tokenPair, err := util.GenerateTokenPair(account.UID)
	if err != nil {
		return LoginRes{}, err
	}

	return LoginRes{
		UID:          account.UID,
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresAt:    tokenPair.ExpiresAt,
	}, nil
}

func (s *service) Refresh(ctx context.Context, req *RefreshReq) (*LoginRes, error) {

	// 1. 解析 refresh_token，静态校验
	claims, err := util.ParseRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	// 2. 从jwt中提取UID
	uid := claims.UID

	// 3. redis 值校验
	savedRefreshToken, err2 := s.repo.GetRefreshToken(ctx, uid)
	if err2 != nil {
		return nil, err2
	}

	// 4. redis 取出来的token和前端传的 refresh token进行对比
	if savedRefreshToken != req.RefreshToken {
		return nil, errors.New("刷新token无效")
	}

	// 5.签发新的token对
	tokenPair, err3 := util.GenerateTokenPair(uid)
	if err3 != nil {
		return nil, err3
	}

	// 6. 新的refresh token存在redis
	if err := s.repo.SetRefreshToken(ctx, uid, tokenPair.RefreshToken, 7*24*time.Hour); err != nil {
		return nil, err
	}

	// 7. 返回 签发的token信息
	return &LoginRes{
		UID:          uid,
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
		ExpiresAt:    tokenPair.ExpiresAt,
	}, nil
}
