package http

// HttpResponse 常规响应实体封装
// swagger:model HttpResponse
type HttpResponse struct {
	/* 业务状态码 */
	Code Code `json:"code"`
	/* 响应描述 */
	Message string `json:"message"`
	/* 响应数据(可以为空) */
	Data interface{} `json:"data,omitempty"`
}

// HttpPageResponse 分页响应实体封装，继承 HttpResponse
// swagger:model HttpPageResponse
type HttpPageResponse struct {
	/* 基础响应：继承 HttpResponse */
	HttpResponse
	/* 响应分页页码 */
	Page int `json:"page"`
	/* 响应分页大小 */
	Size int `json:"size"`
	/* 响应分页总条数 */
	Total int `json:"total"`
}
