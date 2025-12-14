package user

import (
	"mall-api/internal/pkg/http"
	"time"
)

// 【新增】请求体
// swagger:model CreateReq
type CreateReq struct {

	// 必填，且通常有长度限制
	Username string `json:"username" binding:"required,min=3,max=64"`

	// 必填，创建时传入明文密码
	Password string `json:"password" binding:"required,min=6,max=32"`

	// 选填，但如果有值必须符合邮箱格式
	Email string `json:"email" binding:"omitempty,email"`

	// 角色：不要写死 oneof，使用 user.Role(req.Role).IsValid() 统一校验（在 handler/service 层做）
	Role string `json:"role" binding:"required"`
}

// 【新增】响应体
// swagger:model CreateRes
type CreateRes struct{}

// 【删除】请求体
// swagger:model DeleteReq
type DeleteReq struct{}

// 【删除】响应体
// swagger:model DeleteRes
type DeleteRes struct{}

// 【修改】请求体
// swagger:model UpdateReq
type UpdateReq struct {
	// 允许修改邮箱
	Email string `json:"email" binding:"omitempty,email"`

	// 允许修改角色：不要写死 oneof，使用 user.Role(req.Role).IsValid() 统一校验（在 handler/service 层做）
	Role string `json:"role" binding:"omitempty"`

	// 使用指针，以便区分 "不修改" 和 "修改为禁用(false)"
	IsActive *bool `json:"is_active"`
}

// 【修改】响应体
// swagger:model UpdateRes
type UpdateRes struct{}

// 【读取】请求体
// swagger:model ReadReq
type ReadReq struct {

	// 分页请求结构体复用
	http.PageReq

	// 角色：枚举
	Role string `form:"role"`

	// 关键字：多字段综合搜索， email/username/uid
	Keyword string `form:"keyword"`
}

// 【读取】响应体
// swagger:model ReadRes
type ReadRes struct {

	/** ID (对应数据库的 UID) - 前端通常习惯叫 id */
	ID string `json:"id"`

	/** 登录用户名 */
	Username string `json:"username"`

	/** 邮箱 */
	Email string `json:"email"`

	/** 角色标识 */
	Role string `json:"role"`

	/** 账号状态 */
	IsActive bool `json:"is_active"`

	/** 创建时间 (JSON会自动格式化为 RFC3339 字符串) */
	CreatedAt time.Time `json:"created_at"`

	/** 更新时间 */
	UpdatedAt time.Time `json:"updated_at"`
}
