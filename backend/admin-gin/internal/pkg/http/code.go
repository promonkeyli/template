package http

// Code 业务状态码
type Code int

// 业务状态码定义
const (
	Success Code = 0 // 成功

	Failed       Code = 2000 // 失败
	Unauthorized Code = 2001 // 未授权
	NotFound     Code = 2002 // 资源不存在
)

// codeMessages 业务状态码对应的默认消息
var codeMessages = map[Code]string{
	Success: "成功",

	Failed:       "操作失败",
	Unauthorized: "未授权",
	NotFound:     "资源不存在",
}

// Message 获取业务状态码对应的默认消息
func (c Code) Message() string {
	if msg, exists := codeMessages[c]; exists {
		return msg
	}
	return "未知错误"
}
