package auth

import (
	"net/http"

	pkghttp "mall-api/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// @Summary		用户注册
// @Description	使用用户名密码进行注册
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			data	body		RegisterReq					true	"注册信息"
// @Success		200		{object}	pkghttp.HttpResponse[any]	"成功"
// @Failure		200		{object}	pkghttp.HttpResponse[any]	"失败"
// @Router			/admin/auth/register [post]
func (h *Handler) Register(c *gin.Context) {

	// 1. 读取接口传参
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkghttp.Fail(c, http.StatusBadRequest, "参数错误")
		return
	}

	// 2. 调用 service 层的用户注册
	if err := h.service.Register(&req); err != nil {
		pkghttp.Fail(c, http.StatusConflict, err.Error())
		return
	}

	// 3. 返回成功
	pkghttp.OK(c, gin.H{
		"message": "用户注册成功",
	})

}

// @Summary		用户登录
// @Description	使用用户名密码登录，成功后返回 token 对
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			data	body		LoginReq						true	"登录信息"
// @Success		200		{object}	pkghttp.HttpResponse[LoginRes]	"成功"
// @Failure		200		{object}	pkghttp.HttpResponse[any]		"失败"
// @Router			/admin/auth/login [post]
func (h *Handler) Login(c *gin.Context) {

	// 1.读取接口传参
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkghttp.Fail(c, http.StatusBadRequest, "参数错误")
		return
	}

	// 2.调用 service 层的 login 业务
	res, err := h.service.Login(c.Request.Context(), &req)
	if err != nil {
		pkghttp.Fail(c, http.StatusUnauthorized, err.Error())
		return
	}

	pkghttp.OK(c, res)
}

// 注销
// func (h *Handler) Logout(c *gin.Context) {

// }

// @Summary		刷新 Token
// @Description	使用 refresh token 刷新访问 token，成功后返回新的 token 对
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			data	body		RefreshReq						true	"刷新 token 请求"
// @Success		200		{object}	pkghttp.HttpResponse[LoginRes]	"成功"
// @Failure		200		{object}	pkghttp.HttpResponse[any]		"失败"
// @Router			/admin/auth/refresh [post]
func (h *Handler) RefreshToken(c *gin.Context) {

	// 1.参数校验
	var req RefreshReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkghttp.Fail(c, http.StatusBadRequest, "参数错误")
		return
	}

	// 2.调用 service 层的 refresh 业务
	res, err := h.service.Refresh(c.Request.Context(), &req)
	if err != nil {
		pkghttp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 3.返回 token 对
	pkghttp.OK(c, res)
}
