package middleware

import (
	pkghttp "mall-api/internal/pkg/http"
	"mall-api/internal/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var jt *jwt.JWT

// InitJWT 注入 JWT 引擎供中间件使用
func InitJWT(svc *jwt.JWT) {
	jt = svc
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")

		// 1. 检查是否存在且格式为 "Bearer <token>"
		if tokenHeader == "" {
			pkghttp.Fail(c, http.StatusUnauthorized, "缺少认证令牌")
			return
		}

		parts := strings.SplitN(tokenHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			pkghttp.Fail(c, http.StatusUnauthorized, "令牌格式错误，请使用 Bearer 格式")
			return
		}

		tokenString := parts[1]

		// 2. 验证令牌
		claims, err := jt.ParseToken(tokenString, "access")
		if err != nil {
			// 区分过期和其他错误
			if err == jwt.ErrExpiredToken {
				pkghttp.Fail(c, http.StatusUnauthorized, "登录已过期，请重新登录")
			} else {
				pkghttp.Fail(c, http.StatusUnauthorized, "令牌无效或已失效")
			}
			return
		}

		// 3. 存储结果并放行
		// 存储到上下文供后续 Controller 使用：uid := c.GetUint("uid")
		c.Set("uid", claims.UID)
		c.Next()
	}
}
