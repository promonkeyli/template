package user

import (
	"errors"
	"net/http"
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

// mapServiceErrorCode 将 service 层错误映射为统一的 HTTP 状态码
// 约定：service 返回的 ValidationError => 422（参数/业务校验失败）
func mapServiceErrorCode(err error) int {
	if IsValidationError(err) {
		return http.StatusUnprocessableEntity
	}
	return http.StatusInternalServerError
}

// @Summary		获取用户列表
// @Description	支持分页以及条件查询
// @ID				listUser
// @Security		BearerAuth
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			params	query		listReq											true	"查询参数"
// @Success		200		{object}	pkghttp.HttpResponse[pkghttp.PageRes[listRes]]	"查询成功"
// @Router			/admin/user [get]
// @Router			/admin/user [get]
func (h *Handler) List(c *gin.Context) {
	var req listReq
	if err := c.ShouldBindQuery(&req); err != nil {
		pkghttp.Fail(c, http.StatusBadRequest, "参数错误")
		return
	}

	res, total, err := h.service.List(c.Request.Context(), &req)
	if err != nil {
		pkghttp.Fail(c, mapServiceErrorCode(err), err.Error())
		return
	}

	pkghttp.OKWithPage(c, pkghttp.PageRes[listRes]{
		List:  res,
		Total: int64(total),
		Page:  req.GetPage(),
		Size:  req.GetPageSize(),
	})
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkghttp.Fail(c, http.StatusBadRequest, "参数错误")
		return
	}

	// Role 校验（避免 dto 里写死 oneof 与常量不一致）
	if !Role(req.Role).IsValid() {
		pkghttp.Fail(c, http.StatusUnprocessableEntity, "role 不合法")
		return
	}

	if err := h.service.Create(c.Request.Context(), &req); err != nil {
		pkghttp.Fail(c, mapServiceErrorCode(err), err.Error())
		return
	}

	pkghttp.OK(c, CreateRes{})
}

func (h *Handler) Update(c *gin.Context) {
	uid := strings.TrimSpace(c.Param("uid"))
	if uid == "" {
		pkghttp.Fail(c, http.StatusBadRequest, "uid 不能为空")
		return
	}

	var req UpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkghttp.Fail(c, http.StatusBadRequest, "参数错误")
		return
	}

	if req.Role != "" && !Role(req.Role).IsValid() {
		pkghttp.Fail(c, http.StatusUnprocessableEntity, "role 不合法")
		return
	}

	if err := h.service.Update(c.Request.Context(), uid, &req); err != nil {
		pkghttp.Fail(c, mapServiceErrorCode(err), err.Error())
		return
	}

	pkghttp.OK(c, UpdateRes{})
}

func (h *Handler) Delete(c *gin.Context) {
	uid := strings.TrimSpace(c.Param("uid"))
	if uid == "" {
		pkghttp.Fail(c, http.StatusBadRequest, "uid 不能为空")
		return
	}

	if err := h.service.Delete(c.Request.Context(), uid); err != nil {
		// 如果你用 gorm.ErrRecordNotFound，可以映射成更友好的提示
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pkghttp.Fail(c, http.StatusNotFound, "用户不存在")
			return
		}

		pkghttp.Fail(c, mapServiceErrorCode(err), err.Error())
		return
	}

	pkghttp.OK(c, DeleteRes{})
}
