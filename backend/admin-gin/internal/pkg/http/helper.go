package http

import (
	"mall-api/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OKOption 成功响应选项
type OKOption struct {
	Message string      // 响应消息，为空则使用默认消息 "成功"
	Data    interface{} // 响应数据
}

// PageOption 分页响应选项
type PageOption struct {
	Message string      // 响应消息，为空则使用默认消息 "成功"
	Data    interface{} // 响应数据
	Page    int         // 当前页码
	Size    int         // 每页数量
	Total   int         // 总记录数
}

// FailOption 失败响应选项
type FailOption struct {
	Code    Code   // 业务状态码
	Message string // 错误消息，必填
}

// OK 发送成功响应（不分页）
// 使用示例：
//   - network.OK(c, &OKOption{Data: user})                           // 使用默认消息
//   - network.OK(c, &OKOption{Message: "登录成功", Data: token})      // 自定义消息
func OK(c *gin.Context, opt *OKOption) {
	c.JSON(http.StatusOK, NewResponse(Success, opt.Message, opt.Data))
}

// OKWithPage 发送成功分页响应
// 使用示例：
//   - network.OKWithPage(c, &PageOption{Data: items, Page: 1, Size: 10, Total: 100})                    // 使用默认消息
//   - network.OKWithPage(c, &PageOption{Message: "查询成功", Data: items, Page: 1, Size: 10, Total: 100}) // 自定义消息
func OKWithPage(c *gin.Context, opt *PageOption) {
	c.JSON(http.StatusOK, NewPageResponse(Success, opt.Message, opt.Data, opt.Page, opt.Size, opt.Total))
}

// Fail 发送失败响应并自动记录错误日志
// 使用示例：
//   - network.Fail(c, &FailOption{Code: network.Failed, Message: "参数错误"})
//   - network.Fail(c, &FailOption{Code: network.Unauthorized, Message: err.Error()})
func Fail(c *gin.Context, opt *FailOption) {
	// 自动记录错误日志
	logger.Log.Error(opt.Message,
		"path", c.Request.URL.Path,
		"method", c.Request.Method,
		"code", opt.Code,
		"client_ip", c.ClientIP(),
	)

	c.JSON(http.StatusOK, NewResponse(opt.Code, opt.Message, nil))
}
