package user

import (
	"errors"
	"strings"

	pkghttp "mall-api/internal/pkg/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// mapServiceErrorCode 将 service 层错误映射为统一的业务 code
// 约定：service 返回的 ValidationError => 422（参数/业务校验失败）
func mapServiceErrorCode(err error) pkghttp.Code {
	if IsValidationError(err) {
		return pkghttp.ValidationFail
	}
	return pkghttp.InternalError
}

// @Summary		管理员-用户列表
// @Description	分页获取后台用户列表，支持按角色与关键字筛选（keyword 可匹配 uid/username/email）
// @Tags			AdminUser
// @Accept			json
// @Produce		json
// @Param			page	query		int		false	"页码"	minimum(1)
// @Param			size	query		int		false	"每页数量"	minimum(1)	maximum(100)
// @Param			role	query		string	false	"角色"
// @Param			keyword	query		string	false	"关键字(uid/username/email)"
// @Success		200		{object}	pkghttp.HttpPageResponse
// @Failure		200		{object}	pkghttp.HttpResponse
// @Router			/admin/user [get]
// @Security		BearerAuth
func (h *Handler) List(c *gin.Context) {
	var req ReadReq
	if err := c.ShouldBindQuery(&req); err != nil {
		pkghttp.Fail(c, &pkghttp.FailOption{
			Code:    pkghttp.InvalidParam,
			Message: "参数错误",
		})
		return
	}

	res, total, err := h.service.List(c.Request.Context(), &req)
	if err != nil {
		pkghttp.Fail(c, &pkghttp.FailOption{
			Code:    mapServiceErrorCode(err),
			Message: err.Error(),
		})
		return
	}

	pkghttp.OKWithPage(c, &pkghttp.PageOption{
		Data:    res,
		Message: "查询成功",
		Page:    req.GetPage(),
		Size:    req.GetPageSize(),
		Total:   total,
	})
}

// @Summary		管理员-创建用户
// @Description	创建后台用户（密码会在服务层进行 bcrypt 加密）
// @Tags			AdminUser
// @Accept			json
// @Produce		json
// @Param			data	body		CreateReq	true	"创建用户请求"
// @Success		200		{object}	pkghttp.HttpResponse
// @Failure		200		{object}	pkghttp.HttpResponse
// @Router			/admin/user [post]
// @Security		BearerAuth
func (h *Handler) Create(c *gin.Context) {
	var req CreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkghttp.Fail(c, &pkghttp.FailOption{
			Code:    pkghttp.InvalidParam,
			Message: "参数错误",
		})
		return
	}

	// Role 校验（避免 dto 里写死 oneof 与常量不一致）
	if !Role(req.Role).IsValid() {
		pkghttp.Fail(c, &pkghttp.FailOption{
			Code:    pkghttp.ValidationFail,
			Message: "role 不合法",
		})
		return
	}

	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		pkghttp.Fail(c, &pkghttp.FailOption{
			Code:    mapServiceErrorCode(err),
			Message: err.Error(),
		})
		return
	}

	pkghttp.OK(c, &pkghttp.OKOption{
		Data:    CreateRes{},
		Message: "创建成功",
	})
}

// @Summary		管理员-更新用户
// @Description	按 UID 更新后台用户（邮箱/角色/启用状态）
// @Tags			AdminUser
// @Accept			json
// @Produce		json
// @Param			uid		path		string		true	"用户 UID"
// @Param			data	body		UpdateReq	true	"更新用户请求"
// @Success		200		{object}	pkghttp.HttpResponse
// @Failure		200		{object}	pkghttp.HttpResponse
// @Router			/admin/user/{uid} [put]
// @Security		BearerAuth
func (h *Handler) Update(c *gin.Context) {
	uid := strings.TrimSpace(c.Param("uid"))
	if uid == "" {
		pkghttp.Fail(c, &pkghttp.FailOption{
			Code:    pkghttp.InvalidParam,
			Message: "uid 不能为空",
		})
		return
	}

	var req UpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkghttp.Fail(c, &pkghttp.FailOption{
			Code:    pkghttp.InvalidParam,
			Message: "参数错误",
		})
		return
	}

	if req.Role != "" && !Role(req.Role).IsValid() {
		pkghttp.Fail(c, &pkghttp.FailOption{
			Code:    pkghttp.ValidationFail,
			Message: "role 不合法",
		})
		return
	}

	if err := h.service.Update(c.Request.Context(), uid, &req); err != nil {
		pkghttp.Fail(c, &pkghttp.FailOption{
			Code:    mapServiceErrorCode(err),
			Message: err.Error(),
		})
		return
	}

	pkghttp.OK(c, &pkghttp.OKOption{
		Data:    UpdateRes{},
		Message: "更新成功",
	})
}

// @Summary		管理员-删除用户
// @Description	按 UID 删除后台用户（建议实现为软删除：is_deleted=true）
// @Tags			AdminUser
// @Accept			json
// @Produce		json
// @Param			uid	path		string	true	"用户 UID"
// @Success		200	{object}	pkghttp.HttpResponse
// @Failure		200	{object}	pkghttp.HttpResponse
// @Router			/admin/user/{uid} [delete]
// @Security		BearerAuth
func (h *Handler) Delete(c *gin.Context) {
	uid := strings.TrimSpace(c.Param("uid"))
	if uid == "" {
		pkghttp.Fail(c, &pkghttp.FailOption{
			Code:    pkghttp.InvalidParam,
			Message: "uid 不能为空",
		})
		return
	}

	if err := h.service.Delete(c.Request.Context(), uid); err != nil {
		// 如果你用 gorm.ErrRecordNotFound，可以映射成更友好的提示
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pkghttp.Fail(c, &pkghttp.FailOption{
				Code:    pkghttp.NotFound,
				Message: "用户不存在",
			})
			return
		}

		pkghttp.Fail(c, &pkghttp.FailOption{
			Code:    mapServiceErrorCode(err),
			Message: err.Error(),
		})
		return
	}

	pkghttp.OK(c, &pkghttp.OKOption{
		Data:    DeleteRes{},
		Message: "删除成功",
	})
}
