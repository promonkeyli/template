package http

type HttpPageRequest struct {
	Page int `form:"page" json:"page" binding:"min=1" example:"1"`          // 当前页码，默认 1
	Size int `form:"size" json:"size" binding:"min=1,max=100" example:"10"` // 每页数量，默认 10，最大 100
}

func (r *HttpPageRequest) GetPage() int {
	if r.Page <= 0 {
		return 1
	}
	return r.Page
}

func (r *HttpPageRequest) GetPageSize() int {
	// 灵活处理：如果不传，默认10；如果传了，限制在 1-100 之间
	switch {
	case r.Size <= 0:
		return 10
	case r.Size > 100:
		return 100
	default:
		return r.Size
	}
}

func (r *HttpPageRequest) GetOffset() int {
	return (r.GetPage() - 1) * r.GetPageSize()
}
