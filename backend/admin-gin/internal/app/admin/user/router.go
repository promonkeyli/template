package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, handlers *Handler) {

	// userGroup := r.Group("/user")
	// userGroup.Use(middleware.JWT())
	// {
	// 	// user 列表
	// 	userGroup.GET("", handlers.List)

	// 	// user 新增
	// 	userGroup.POST("", handlers.Create)

	// 	// user 更新（按 UID）
	// 	userGroup.PUT("/:uid", handlers.Update)

	// 	// user 删除（按 UID）
	// 	userGroup.DELETE("/:uid", handlers.Delete)
	// }
}
