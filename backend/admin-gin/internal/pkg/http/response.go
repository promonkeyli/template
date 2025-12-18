package http

type HttpResponse[T any] struct {
	/* 业务状态码 */
	Code Code `json:"code" example:"200"`
	/* 响应描述 */
	Message string `json:"message" example:"操作成功"`
	/* 响应数据(可以为空) */
	Data T `json:"data,omitempty"`
}

type HttpPageResponse[T any] struct {
	/* 嵌入 HttpResponse 使用泛型切片 */
	HttpResponse[[]T]
	/* 页码 */
	Page int `json:"page" example:"1"`
	/* 页数 */
	Size int `json:"size" example:"10"`
	/* 总条数 */
	Total int64 `json:"total" example:"50"`
}
