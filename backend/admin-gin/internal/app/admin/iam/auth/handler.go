package auth

import (
	"net/http"

	"mall-api/internal/pkg/cookie"
	pkghttp "mall-api/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	se service
	cm *cookie.CookieManager
}

func newHandler(se service, cm *cookie.CookieManager) *handler {
	return &handler{se: se, cm: cm}
}

// @Summary		用户登录
// @Description	用户名/密码登录
// @ID				login
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			data	body		loginReq						true	"登录参数"
// @Success		200		{object}	pkghttp.HttpResponse[loginRes]	"登录成功"
// @Router			/admin/auth/login [post]
func (h *handler) login(c *gin.Context) {

	// 1.读取接口传参
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkghttp.Fail(c, http.StatusBadRequest)
		return
	}

	// 2.调用 service 层的 login 业务
	res, err := h.se.login(c.Request.Context(), &req)
	if err != nil {
		pkghttp.Fail(c, http.StatusUnauthorized, err.Error())
		return
	}

	// 3. gin 设置 refresh cookie
	h.cm.Set(c, res.RefreshToken)

	pkghttp.OK(c, res)
}

// @Summary		用户注册
// @Description	用户名/密码进行注册
// @ID				register
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			data	body		registerReq					true	"注册参数"
// @Success		200		{object}	pkghttp.HttpResponse[Empty]	"注册成功"
// @Router			/admin/auth/register [post]
func (h *handler) register(c *gin.Context) {

	// 1. 读取接口传参
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkghttp.Fail(c, http.StatusBadRequest)
		return
	}

	// 2. 调用 service 层的用户注册
	if err := h.se.register(&req); err != nil {
		pkghttp.Fail(c, http.StatusConflict, err.Error())
		return
	}

	// 3. 返回成功
	pkghttp.OK(c, pkghttp.Empty{})

}

// @Summary		用户注销
// @Description	用户注销,同时移除刷新令牌
// @Security		BearerAuth
// @ID				logout
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			Cookie	header		string						true	"格式: refresh_token=xxx"
// @Success		200		{object}	pkghttp.HttpResponse[Empty]	"注销成功"
// @Router			/admin/auth/session/logout [post]
func (h *handler) logout(c *gin.Context) {

	// 1.参数 refresh_token, 从 cookie 中取值
	refreshToken, err := h.cm.Get(c)
	if err != nil {
		pkghttp.Fail(c, http.StatusBadRequest)
		return
	}

	// 2. 调用 service 层注销业务
	code, err := h.se.logout(c.Request.Context(), refreshToken)
	if err != nil {
		pkghttp.Fail(c, code, err.Error())
		return
	}

	// 3. 清除 refresh_token cookie
	h.cm.Remove(c)

	// 4. 成功
	pkghttp.OK(c, pkghttp.Empty{})

}

// @Summary		刷新令牌
// @Description	用于短期令牌 Access_Token 续期
// @ID				refreshToken
// @Security		CookieAuth
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			Cookie	header		string							true	"格式: refresh_token=xxx"
// @Success		200		{object}	pkghttp.HttpResponse[loginRes]	"刷新成功"
// @Router			/admin/auth/session/refresh [post]
func (h *handler) refresh(c *gin.Context) {

	// 1.参数 refresh_token, 从 cookie 中取值
	refreshToken, err := h.cm.Get(c)
	if err != nil {
		pkghttp.Fail(c, http.StatusBadRequest)
		return
	}

	// 2.调用 service 层的 refresh 业务
	res, err := h.se.refresh(c.Request.Context(), refreshToken)
	if err != nil {
		pkghttp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 3.返回 token 对
	pkghttp.OK(c, res)
}
