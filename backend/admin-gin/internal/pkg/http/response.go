package http

type HttpResponse[T any] struct {
	// code: HTTP 状态码（与 net/http 保持一致）
	Code int `json:"code" example:"200"`
	// message: 响应描述（默认可用 internal/pkg/http.StatusText(code)）
	Message string `json:"message" example:"成功"`
	// data: 响应数据（可以为空）
	Data T `json:"data,omitempty"`
}

type HttpPageResponse[T any] struct {
	// 嵌入 HttpResponse 使用泛型切片
	HttpResponse[[]T]
	// page: 当前页码
	Page int `json:"page" example:"1"`
	// size: 每页数量
	Size int `json:"size" example:"10"`
	// total: 总条数
	Total int64 `json:"total" example:"50"`
}
