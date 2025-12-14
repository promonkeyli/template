package auth

import (
	"errors"
	"mall-api/internal/app/admin/staff"
	"mall-api/internal/pkg/util"
	"time"

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
	if exist, err := s.repo.FindStaffIsExist(req.Username); err != nil {
		return err
	} else if exist {
		return errors.New("用户已经存在")
	}

	// 2. 生成全局唯一 UID
	uid := util.MewUUID()

	// 3. 加密密码：password []byte
	// 你要加密的原始密码，需要转成字节切片。
	// 这里写 []byte(req.Password) 就是把用户输入的字符串密码转换成字节切片。
	// cost int：加密强度（迭代次数，越高越安全，但计算越慢）bcrypt.DefaultCost 是官方推荐默认值，通常是 10。
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 4. 构建 staff 数据
	staffData := &staff.Staff{
		UID:       uid,
		IsActive:  true,
		Username:  req.Username,
		IsDeleted: false,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.repo.CreateStaff(staffData); err != nil {
		return err
	}
	return nil
}

func (s *service) Login(req *LoginReq) (LoginRes, error) {
	// 1. 查找用户
	targetStaff, err1 := s.repo.FindStaffByName(req.Username)
	if err1 != nil {
		return LoginRes{}, err1
	}
	// 2. 校验用户密码是否正确（先解密数据库的密码，再对比）
	err2 := bcrypt.CompareHashAndPassword([]byte(targetStaff.Password), []byte(req.Password))
	if err2 != nil {
		return LoginRes{}, errors.New("密码不正确")
	}
	// 3. 构建token
	token_pair, err3 := util.GenerateTokenPair(targetStaff.UID)
	if err3 != nil {
		return LoginRes{}, err3
	} else {
		return LoginRes{
			UID:          targetStaff.UID,
			AccessToken:  token_pair.AccessToken,
			RefreshToken: token_pair.RefreshToken,
			ExpiresAt:    token_pair.ExpiresAt,
		}, nil
	}
}
