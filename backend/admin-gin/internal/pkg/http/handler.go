package http

// NewResponse 创建标准响应对象
// 如果 message 为空字符串，则使用 code 的默认消息
// 使用示例：
//   - network.NewResponse(network.Success, "", data)  // 使用默认消息 "成功"
//   - network.NewResponse(network.Success, "登录成功", data)  // 使用自定义消息
func NewResponse(code Code, message string, data interface{}) HttpResponse {
	// 如果没有提供自定义消息，使用默认消息
	if message == "" {
		message = code.Message()
	}

	return HttpResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewPageResponse 创建分页响应对象
func NewPageResponse(code Code, message string, data interface{}, page, size, total int) HttpPageResponse {
	// 如果没有提供自定义消息，使用默认消息
	if message == "" {
		message = code.Message()
	}

	return HttpPageResponse{
		HttpResponse: HttpResponse{
			Code:    code,
			Message: message,
			Data:    data,
		},
		Page:  page,
		Size:  size,
		Total: total,
	}
}
