package http

import "net/http"

type Code int

const (
	// 2xx Success
	Success Code = http.StatusOK // 200

	// 4xx Client Errors
	InvalidParam   Code = http.StatusBadRequest          // 400: 参数错误（解析/类型不符）
	Unauthorized   Code = http.StatusUnauthorized        // 401: 未登录或 Token 失效
	Forbidden      Code = http.StatusForbidden           // 403: 已登录但权限不足
	NotFound       Code = http.StatusNotFound            // 404: 资源不存在
	Conflict       Code = http.StatusConflict            // 409: 资源冲突（如唯一键重复）
	ValidationFail Code = http.StatusUnprocessableEntity // 422: 参数格式校验失败（Validator）

	// 5xx Server Errors
	InternalError  Code = http.StatusInternalServerError // 500: 服务器内部错误
	NotImplemented Code = http.StatusNotImplemented      // 501: 功能尚未实现

	// 语义化别名（保持与 net/http 一致的值）
	AlreadyExists Code = http.StatusConflict            // 409
	TokenExpired  Code = http.StatusUnauthorized        // 401
	TokenInvalid  Code = http.StatusUnauthorized        // 401
	Failed        Code = http.StatusInternalServerError // 500
)

// code => desc
var codeMessages = map[Code]string{
	Success: "请求成功",

	InvalidParam:   "请求参数格式错误",
	Unauthorized:   "鉴权失败，请重新登录",
	Forbidden:      "权限不足，拒绝访问",
	NotFound:       "请求的资源不存在",
	Conflict:       "资源已存在或发生冲突",
	ValidationFail: "字段校验失败",

	InternalError:  "服务器内部错误",
	NotImplemented: "功能开发中，敬请期待",
}

// 中文描述映射 func
func (c Code) GetMsg() string {
	if msg, ok := codeMessages[c]; ok {
		return msg
	}
	// 如果找不到自定义描述，返回标准 HTTP 文本（英文）作为兜底
	return http.StatusText(int(c))
}
