package auth

type loginReq struct {
	Username string `json:"username" binding:"required" example:"admin"`  // 用户名
	Password string `json:"password" binding:"required" example:"123456"` // 密码
}

type loginRes struct {
	UID          string `json:"uid"`           // 用户UID
	AccessToken  string `json:"access_token"`  // 访问令牌: 15分钟过期
	RefreshToken string `json:"refresh_token"` // 刷新令牌：7天过期
	ExpiresAt    int64  `json:"expires_at"`    // 过期时间：访问令牌 Access_token 过期时间(秒)
}

type registerReq struct {
	loginReq
}

type refreshReq struct {
	RefreshToken string `json:"refresh_token" binding:"required"` // 刷新令牌
}

type logoutReq struct {
	refreshReq
}
