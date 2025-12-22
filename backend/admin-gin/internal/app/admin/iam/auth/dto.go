package auth

type loginReq struct {
	Username string `json:"username" binding:"required" example:"admin"`  // 用户名
	Password string `json:"password" binding:"required" example:"123456"` // 密码
}

type loginRes struct {
	UID          string `json:"uid"`                    // 用户UID
	AccessToken  string `json:"access_token"`           // 访问令牌: 15分钟过期
	ExpiresAt    int64  `json:"expires_at"`             // 过期时间：访问令牌 Access_token 过期时间(秒)
	RefreshToken string `json:"-" swaggerignore:"true"` // 刷新令牌不返回前端,JSON 转换也不转换该字段，该字段只在/admin/auth/refresh 接口cookie中携带，还需要配置必要的安全设置
}

type registerReq struct {
	loginReq
}
