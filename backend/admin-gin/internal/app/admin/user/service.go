package user

import (
	"context"
	"errors"
	"strings"
	"time"

	"mall-api/internal/pkg/uuid"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ValidationError 用于标记“参数/业务校验类错误”，方便 handler 层统一映射为 422。
type ValidationError struct {
	msg string
}

func (e *ValidationError) Error() string { return e.msg }

func newValidationError(msg string) error { return &ValidationError{msg: msg} }

func IsValidationError(err error) bool {
	var ve *ValidationError
	return errors.As(err, &ve)
}

type Service interface {
	// List 分页查询后台用户列表
	List(ctx context.Context, req *ReadReq) ([]ReadRes, int, error)
	// Create 创建后台用户
	Create(ctx context.Context, req *CreateReq) error
	// Update 按 UID 更新后台用户（邮箱/角色/启用状态）
	Update(ctx context.Context, uid string, req *UpdateReq) error
	// Delete 按 UID 删除后台用户（软删除）
	Delete(ctx context.Context, uid string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) List(ctx context.Context, req *ReadReq) ([]ReadRes, int, error) {
	page := req.GetPage()
	size := req.GetPageSize()

	users, total, err := s.repo.List(ctx, page, size, strings.TrimSpace(req.Role), strings.TrimSpace(req.Keyword))
	if err != nil {
		return nil, 0, err
	}

	out := make([]ReadRes, 0, len(users))
	for _, u := range users {
		out = append(out, ReadRes{
			ID:        u.UID,
			Username:  u.Username,
			Email:     u.Email,
			Role:      u.Role,
			IsActive:  u.IsActive,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return out, total, nil
}

func (s *service) Create(ctx context.Context, req *CreateReq) error {
	// 角色校验（保持与 constant.go 一致）
	if !Role(req.Role).IsValid() {
		return newValidationError("role 不合法")
	}

	// 用户名唯一性检查
	exist, err := s.repo.ExistsByUsername(ctx, req.Username)
	if err != nil {
		return err
	}
	if exist {
		return newValidationError("用户名已存在")
	}

	// 邮箱唯一性检查（如果传了 email）
	if strings.TrimSpace(req.Email) != "" {
		existEmail, err := s.repo.ExistsByEmail(ctx, req.Email)
		if err != nil {
			return err
		}
		if existEmail {
			return newValidationError("邮箱已存在")
		}
	}

	uid := uuid.NewUUID()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	u := &User{
		UID:       uid,
		Username:  req.Username,
		Email:     strings.TrimSpace(req.Email),
		Password:  string(hashedPassword),
		Role:      req.Role,
		IsActive:  true,
		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return s.repo.Create(ctx, u)
}

func (s *service) Update(ctx context.Context, uid string, req *UpdateReq) error {
	uid = strings.TrimSpace(uid)
	if uid == "" {
		return newValidationError("uid 不能为空")
	}

	// 查询目标用户
	u, err := s.repo.GetByUID(ctx, uid)
	if err != nil {
		return err
	}
	if u.IsDeleted {
		return gorm.ErrRecordNotFound
	}

	// 更新字段（部分更新）
	updates := map[string]any{}

	if req.Email != "" {
		email := strings.TrimSpace(req.Email)
		// 邮箱唯一性检查
		existEmail, err := s.repo.ExistsByEmailExcludeUID(ctx, email, uid)
		if err != nil {
			return err
		}
		if existEmail {
			return newValidationError("邮箱已存在")
		}
		updates["email"] = email
	}

	if req.Role != "" {
		if !Role(req.Role).IsValid() {
			return newValidationError("role 不合法")
		}
		updates["role"] = req.Role
	}

	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}

	if len(updates) == 0 {
		return nil
	}

	updates["updated_at"] = time.Now()
	return s.repo.UpdateByUID(ctx, uid, updates)
}

func (s *service) Delete(ctx context.Context, uid string) error {
	uid = strings.TrimSpace(uid)
	if uid == "" {
		return newValidationError("uid 不能为空")
	}
	return s.repo.SoftDeleteByUID(ctx, uid)
}
