package auth

// swagger:model RegisterReq
type RegisterReq struct {
	Username string `json:"username" binding:"required" example:"admin"`  // 用户名
	Password string `json:"password" binding:"required" example:"123456"` // 密码
}

// swagger:model LoginReq
type LoginReq struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

// swagger:model LoginRes
type LoginRes struct {
	UID          string `json:"uid"`           // 用户UID
	AccessToken  string `json:"access_token"`  // 访问令牌
	RefreshToken string `json:"refresh_token"` // 刷新令牌
	ExpiresAt    int64  `json:"expires_at"`    // 过期时间
}
