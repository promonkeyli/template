package mw

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                 // 允许访问的域名，可以配置多个
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},           // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                                    // 公开的响应头
		AllowCredentials: true,                                                          // 允许包含凭据
		MaxAge:           12 * time.Hour,                                                // 预检请求的缓存时间
	})
}
