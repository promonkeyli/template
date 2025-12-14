package user

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RegisterRouter(r *gin.Engine, d *gorm.DB, rdb *redis.Client) {

	// 数据层：负责数据库交互
	// repo := auth.NewRepository(d)

	// 业务层：负责核心业务逻辑
	// service := NewService(repo)

	// 接入层：负责处理 HTTP 请求
	// handler := NewHandler(service)

	// v1 staff 接口分组
	// asGroup := r.Group("/admin/staff")
	// asGroup.Use(middleware.JWT())
	// {
	// 	asGroup.GET("")    // staff 列表
	// 	asGroup.POST("")   // staff 新增
	// 	asGroup.PUT("")    // staff 更新
	// 	asGroup.DELETE("") // staff 删除
	// }
}
