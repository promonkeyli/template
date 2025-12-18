package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 普通成功响应
func OK[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, HttpResponse[T]{
		Code:    http.StatusOK,
		Message: StatusText(http.StatusOK),
		Data:    data,
	})
}

// 分页成功响应
func OKWithPage[T any](c *gin.Context, data []T, total int64, page, size int) {
	c.JSON(http.StatusOK, HttpPageResponse[T]{
		HttpResponse: HttpResponse[[]T]{
			Code:    http.StatusOK,
			Message: StatusText(http.StatusOK),
			Data:    data,
		},
		Total: total,
		Page:  page,
		Size:  size,
	})
}

// 失败响应（code 使用标准 HTTP 状态码）
func Fail(c *gin.Context, code int, msg ...string) {
	displayMsg := StatusText(code)
	if len(msg) > 0 && msg[0] != "" {
		displayMsg = msg[0]
	}

	// 失败后，handler 中断，不再走下面的业务
	// 注意：这里保持原行为，HTTP 响应状态仍返回 200，真实状态放在 JSON 的 code 字段中
	c.AbortWithStatusJSON(http.StatusOK, HttpResponse[any]{
		Code:    code,
		Message: displayMsg,
		Data:    nil,
	})
}
