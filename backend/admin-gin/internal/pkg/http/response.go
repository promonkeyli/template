package http

type HttpResponse[T any] struct {
	// code: HTTP 状态码
	Code int `json:"code" example:"200"`
	// message: 响应描述
	Message string `json:"message" example:"操作成功"`
	// data: 响应数据（可以为空）
	Data T `json:"data,omitempty"`
}

// 分页包装结构体
type PageRes[T any] struct {
	List  []T   `json:"list"`               // 列表
	Total int64 `json:"total" example:"50"` // 总数
	Page  int   `json:"page" example:"1"`   // 当前页码
	Size  int   `json:"size" example:"10"`  // 分页大小
}

// 响应为空时的结构体
type Empty struct{}
