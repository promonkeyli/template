package http

// PageReq 通用分页请求参数
// 建议设置最大每页数量，防止前端恶意传大数把数据库查挂
type PageReq struct {
	Page int `form:"page" json:"page" binding:"min=1"`         // 当前页码，默认 1
	Size int `form:"size" json:"size" binding:"min=1,max=100"` // 每页数量，默认 10，最大 100
}

// GetPage 获取页码，提供默认值防御
func (r *PageReq) GetPage() int {
	if r.Page <= 0 {
		return 1
	}
	return r.Page
}

// GetPageSize 获取每页数量，提供默认值防御
func (r *PageReq) GetPageSize() int {
	if r.Size <= 0 {
		return 10 // 默认每页 10 条
	}
	if r.Size > 100 {
		return 100 // 限制最大 100 条
	}
	return r.Size
}

// GetOffset 计算数据库查询的偏移量 (用于 GORM 的 Offset)
func (r *PageReq) GetOffset() int {
	return (r.GetPage() - 1) * r.GetPageSize()
}
