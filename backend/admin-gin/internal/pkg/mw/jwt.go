package mw

import (
	"mall-api/internal/pkg/http"
	"mall-api/internal/pkg/util"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWT 标准JWT鉴权中间件
// 用法：在需要鉴权的路由分组中使用 group.Use(middleware.JWT())
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// 检查token是否存在且格式正确
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			http.Fail(c, &http.FailOption{
				Code:    http.Unauthorized,
				Message: "缺少或格式错误的令牌",
			})
			c.Abort()
			return
		}

		// 移除 "Bearer " 前缀
		tokenString = tokenString[7:]

		// 使用封装好的ValidateToken函数验证access token
		claims, err := util.ValidateToken(tokenString)
		if err != nil {
			// 区分token过期和其他错误
			if err == util.ErrExpiredToken {
				http.Fail(c, &http.FailOption{
					Code:    http.Unauthorized,
					Message: "令牌已过期",
				})
			} else {
				http.Fail(c, &http.FailOption{
					Code:    http.Unauthorized,
					Message: "无效的令牌",
				})
			}
			c.Abort()
			return
		}

		// 将用户UID存储到上下文中，供后续处理器使用
		c.Set("uid", claims.UID)
		c.Next()
	}
}
