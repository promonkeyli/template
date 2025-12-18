package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 普通成功响应
func OK[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, HttpResponse[T]{
		Code:    Success,
		Message: Success.GetMsg(),
		Data:    data,
	})
}

// 分页成功响应
func OKWithPage[T any](c *gin.Context, data []T, total int64, page, size int) {
	c.JSON(http.StatusOK, HttpPageResponse[T]{
		HttpResponse: HttpResponse[[]T]{
			Code:    Success,
			Message: Success.GetMsg(),
			Data:    data,
		},
		Total: total,
		Page:  page,
		Size:  size,
	})
}

// 失败响应
func Fail(c *gin.Context, code Code, msg ...string) {
	displayMsg := code.GetMsg()
	if len(msg) > 0 && msg[0] != "" {
		displayMsg = msg[0]
	}

	// 失败后，handler 中断，不再走下面的业务
	c.AbortWithStatusJSON(http.StatusOK, HttpResponse[any]{
		Code:    code,
		Message: displayMsg,
		Data:    nil,
	})
}
