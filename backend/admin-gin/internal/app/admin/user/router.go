package user

import (
	"mall-api/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine, handler *Handler) {

	userGroup := r.Group("/admin/user")
	userGroup.Use(middleware.JWT())
	{
		// user 列表
		userGroup.GET("", handler.List)

		// user 新增
		userGroup.POST("", handler.Create)

		// user 更新（按 UID）
		userGroup.PUT("/:uid", handler.Update)

		// user 删除（按 UID）
		userGroup.DELETE("/:uid", handler.Delete)
	}
}
