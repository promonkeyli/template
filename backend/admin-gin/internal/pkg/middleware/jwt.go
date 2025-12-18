package middleware

import (
	"mall-api/internal/pkg/http"
	"mall-api/internal/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")

		// 1. 检查是否存在且格式为 "Bearer <token>"
		if tokenHeader == "" {
			http.Fail(c, http.Unauthorized, "缺少认证令牌")
			return
		}

		parts := strings.SplitN(tokenHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			http.Fail(c, http.TokenInvalid, "令牌格式错误，请使用 Bearer 格式")
			return
		}

		tokenString := parts[1]

		// 2. 验证令牌
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			// 区分过期和其他错误
			// 假设你之前定义的 http.TokenExpired = 401, http.TokenInvalid = 401
			if err == jwt.ErrExpiredToken {
				http.Fail(c, http.TokenExpired, "登录已过期，请重新登录")
			} else {
				http.Fail(c, http.TokenInvalid, "令牌无效或已失效")
			}
			return
		}

		// 3. 存储结果并放行
		// 存储到上下文供后续 Controller 使用：uid := c.GetUint("uid")
		c.Set("uid", claims.UID)
		c.Next()
	}
}
