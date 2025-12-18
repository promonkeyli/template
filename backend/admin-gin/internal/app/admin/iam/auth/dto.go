package auth

type LoginReq struct {
	Username string `json:"username" binding:"required" example:"admin"`  // 用户名
	Password string `json:"password" binding:"required" example:"123456"` // 密码
}

type LoginRes struct {
	UID          string `json:"uid"`           // 用户UID
	AccessToken  string `json:"access_token"`  // 访问令牌
	RefreshToken string `json:"refresh_token"` // 刷新令牌
	ExpiresAt    int64  `json:"expires_at"`    // 过期时间
}

type RegisterReq struct {
	Username string `json:"username" binding:"required" example:"admin"`  // 用户名
	Password string `json:"password" binding:"required" example:"123456"` // 密码
}

type RefreshReq struct {
	RefreshToken string `json:"refresh_token" binding:"required"` // 刷新令牌
}
