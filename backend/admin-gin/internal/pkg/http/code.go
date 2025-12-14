package http

// Code 业务状态码（用于响应体中的 code 字段）
//
// 规范（按你的要求）：
// - HTTP 响应永远返回 200（见 http.OK/http.Fail 的实现）
// - 响应体中的 Code 值与 HTTP Status Code 数值对齐（便于前后端/日志/排障统一语义）
// - 业务代码中禁止直接使用 net/http 的状态码常量，统一引用本文件中定义的 Code 常量
type Code int

// 常用业务状态码（对齐 HTTP Status Code 数值）
const (
	// 2xx Success
	Success Code = 200

	// 4xx Client Errors
	InvalidParam   Code = 400 // 参数错误（解析/缺失/类型不符）
	ValidationFail Code = 422 // 参数校验失败（binding/validator）

	Unauthorized Code = 401 // 未登录/缺少 token/鉴权失败
	Forbidden    Code = 403 // 已登录但无权限

	NotFound      Code = 404 // 资源不存在
	Conflict      Code = 409 // 资源冲突（状态冲突/并发更新冲突等）
	AlreadyExists Code = 409 // 资源已存在（唯一键冲突等，与 Conflict 同值）

	// 5xx Server Errors
	InternalError  Code = 500 // 服务内部错误（兜底）
	NotImplemented Code = 501 // 功能未实现

	// Token 细分（由于本规范要求 code 与 HTTP 状态码一致，这里保留语义常量但与 Unauthorized 同值）
	// 如果未来你愿意增加 biz_code（或 details），可以在不破坏现有 code 的情况下表达更细粒度原因。
	TokenExpired Code = 401
	TokenInvalid Code = 401

	// 兼容旧名称（如果历史代码里用到了 Failed，保留为 500 的别名）
	Failed Code = 500
)

// codeMessages 业务状态码对应的默认消息
var codeMessages = map[Code]string{
	Success: "成功",

	InvalidParam:   "参数错误",
	ValidationFail: "参数校验失败",

	Unauthorized: "未授权",
	Forbidden:    "无权限",

	NotFound: "资源不存在",

	Conflict: "资源冲突",

	InternalError:  "服务内部错误",
	NotImplemented: "功能未实现",
}

// Message 获取业务状态码对应的默认消息
func (c Code) Message() string {
	if msg, exists := codeMessages[c]; exists {
		return msg
	}
	return "未知错误"
}
