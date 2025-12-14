package auth

import (
	"mall-api/internal/pkg/http"

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
// @Param			data	body		LoginReq	true	"登录信息"
// @Success		200		{object}	http.HttpResponse
// @Failure		400		{object}	http.HttpResponse	"参数错误"
// @Router			/admin/auth/register [post]
func (h *Handler) Register(c *gin.Context) {

	// 1. 读取接口传参
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		http.Fail(c, &http.FailOption{
			Code:    http.Failed,
			Message: "参数错误",
		})
		return
	}

	// 2. 调用 service 层的用户注册
	if err := h.service.Register(&req); err != nil {
		http.Fail(c, &http.FailOption{
			Code:    http.Failed,
			Message: err.Error(),
		})
		return
	}

	// 3. 返回成功
	http.OK(c, &http.OKOption{
		Data:    nil,
		Message: "用户注册成功",
	})

}

// 登陆
func (h *Handler) Login(c *gin.Context) {

	// 1.读取接口传参
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		http.Fail(c, &http.FailOption{
			Code:    http.Failed,
			Message: "参数错误",
		})
		return
	}

	// 2.调用 service 层的 login 业务
	res, err := h.service.Login(&req)
	if err != nil {
		http.Fail(c, &http.FailOption{
			Code:    http.Failed,
			Message: err.Error(),
		})
	} else {
		http.OK(c, &http.OKOption{
			Data:    res,
			Message: "登录成功",
		})
	}
}

// 注销
// func (h *Handler) Logout(c *gin.Context) {

// }
