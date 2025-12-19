package http

import "net/http"

func statusTextZH(code int) string {
	switch code {
	// 1xx: 信息响应（请求已接收，继续处理）
	case http.StatusContinue:
		return "继续"
	case http.StatusSwitchingProtocols:
		return "切换协议"
	case http.StatusProcessing:
		return "处理中"
	case http.StatusEarlyHints:
		return "提前提示"

	// 2xx: 成功（请求已成功处理）
	case http.StatusOK:
		return "成功"
	case http.StatusCreated:
		return "已创建"
	case http.StatusAccepted:
		return "已接受"
	case http.StatusNonAuthoritativeInfo:
		return "非权威信息"
	case http.StatusNoContent:
		return "无内容"
	case http.StatusResetContent:
		return "重置内容"
	case http.StatusPartialContent:
		return "部分内容"
	case http.StatusMultiStatus:
		return "多状态"
	case http.StatusAlreadyReported:
		return "已报告"
	case http.StatusIMUsed:
		return "已使用 IM"

	// 3xx: 重定向（需要进一步操作以完成请求）
	case http.StatusMultipleChoices:
		return "多种选择"
	case http.StatusMovedPermanently:
		return "永久重定向"
	case http.StatusFound:
		return "已找到"
	case http.StatusSeeOther:
		return "查看其他位置"
	case http.StatusNotModified:
		return "未修改"
	case http.StatusUseProxy:
		return "使用代理"
	case http.StatusTemporaryRedirect:
		return "临时重定向"
	case http.StatusPermanentRedirect:
		return "永久重定向"

	// 4xx: 客户端错误（请求包含错误或无法完成）
	case http.StatusBadRequest:
		return "请求参数错误"
	case http.StatusUnauthorized:
		return "未授权"
	case http.StatusPaymentRequired:
		return "需要付费"
	case http.StatusForbidden:
		return "禁止访问"
	case http.StatusNotFound:
		return "未找到"
	case http.StatusMethodNotAllowed:
		return "方法不允许"
	case http.StatusNotAcceptable:
		return "不可接受"
	case http.StatusProxyAuthRequired:
		return "需要代理认证"
	case http.StatusRequestTimeout:
		return "请求超时"
	case http.StatusConflict:
		return "冲突"
	case http.StatusGone:
		return "资源已删除"
	case http.StatusLengthRequired:
		return "需要 Content-Length"
	case http.StatusPreconditionFailed:
		return "前置条件失败"
	case http.StatusRequestEntityTooLarge:
		return "请求实体过大"
	case http.StatusRequestURITooLong:
		return "请求 URI 过长"
	case http.StatusUnsupportedMediaType:
		return "不支持的媒体类型"
	case http.StatusRequestedRangeNotSatisfiable:
		return "请求范围不满足"
	case http.StatusExpectationFailed:
		return "期望失败"
	case http.StatusTeapot:
		return "我是茶壶"
	case http.StatusMisdirectedRequest:
		return "错误的请求目标"
	case http.StatusUnprocessableEntity:
		return "无法处理的实体"
	case http.StatusLocked:
		return "已锁定"
	case http.StatusFailedDependency:
		return "依赖失败"
	case http.StatusTooEarly:
		return "请求过早"
	case http.StatusUpgradeRequired:
		return "需要升级协议"
	case http.StatusPreconditionRequired:
		return "需要前置条件"
	case http.StatusTooManyRequests:
		return "请求过多"
	case http.StatusRequestHeaderFieldsTooLarge:
		return "请求头字段过大"
	case http.StatusUnavailableForLegalReasons:
		return "因法律原因不可用"

	// 5xx: 服务器错误（服务器处理请求时发生错误）
	case http.StatusInternalServerError:
		return "服务器内部错误"
	case http.StatusNotImplemented:
		return "未实现"
	case http.StatusBadGateway:
		return "网关错误"
	case http.StatusServiceUnavailable:
		return "服务不可用"
	case http.StatusGatewayTimeout:
		return "网关超时"
	case http.StatusHTTPVersionNotSupported:
		return "HTTP 版本不受支持"
	case http.StatusVariantAlsoNegotiates:
		return "变体也参与协商"
	case http.StatusInsufficientStorage:
		return "存储空间不足"
	case http.StatusLoopDetected:
		return "检测到循环"
	case http.StatusNotExtended:
		return "未扩展"
	case http.StatusNetworkAuthenticationRequired:
		return "需要网络认证"

	default:
		return ""
	}
}

func StatusText(code int) string {
	if zh := statusTextZH(code); zh != "" {
		return zh
	}
	return http.StatusText(code)
}
