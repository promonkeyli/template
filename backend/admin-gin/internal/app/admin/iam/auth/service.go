package auth

import (
	"errors"
	"mall-api/internal/pkg/util"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(req *RegisterReq) error       // 注册
	Login(req *LoginReq) (LoginRes, error) // 登录
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
	uid := util.MewUUID()

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
