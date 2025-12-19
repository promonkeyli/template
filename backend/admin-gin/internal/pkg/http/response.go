package http

type HttpResponse[T any] struct {
	// code: HTTP 状态码
	Code int `json:"code" example:"200"`
	// message: 响应描述
	Message string `json:"message" example:"操作成功"`
	// data: 响应数据（可以为空）
	Data T `json:"data,omitempty"`
}

type HttpPageResponse[T any] struct {
	HttpResponse[[]T]
	// page: 当前页码
	Page int `json:"page" example:"1"`
	// size: 每页数量
	Size int `json:"size" example:"10"`
	// total: 总条数
	Total int64 `json:"total" example:"50"`
}

// 响应为空时的结构体
type Empty struct{}
