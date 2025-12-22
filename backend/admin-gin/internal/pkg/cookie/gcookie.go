// gin 框架中 设置 cookie
package cookie

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CookieConfig struct {
	// Name 是 Cookie 的键名，例如 "refresh_token"
	Name string `mapstructure:"name"`

	// Path 限制了 Cookie 的作用路径。
	// 建议设置为特定的刷新接口路径（如 "/auth/refresh"），防止在请求普通业务接口时携带，减少暴露风险。
	Path string `mapstructure:"path"`

	// Domain 指定了哪些域名可以接收该 Cookie。
	// 开发环境通常设为 "localhost"，生产环境设为你的顶级域名（如 "example.com"）。
	Domain string `mapstructure:"domain"`

	// MaxAge 定义了 Cookie 的有效期，单位为秒。
	// 例如：604800 表示 7 天。设置为 -1 则会立即删除 Cookie。
	MaxAge int64 `mapstructure:"max_age"`

	// Secure 若为 true，则 Cookie 仅通过 HTTPS 加密协议传输。
	// 注意：本地开发若没有配置 SSL/HTTPS，需设为 false，否则浏览器会拒绝存储。
	Secure bool `mapstructure:"secure"`

	// HttpOnly 若为 true，则 JavaScript 脚本无法读取该 Cookie。
	// 这是防御 XSS（跨站脚本攻击）窃取 Token 的核心屏障。
	HttpOnly bool `mapstructure:"http_only"`

	// SameSite 控制 Cookie 的跨站行为，用于防御 CSRF 攻击。
	// 可选值："Lax"（推荐）、"Strict"（最严格）、"None"（必须配合 Secure=true 使用）。
	SameSite string `mapstructure:"same_site"`
}

// 2. CookieManager: 集中管理的操作对象
type CookieManager struct {
	cfg CookieConfig
}

// NewCookieManager: 构造函数
func NewCookieManager(cfg CookieConfig) *CookieManager {
	return &CookieManager{cfg: cfg}
}

// 3. Set: 写入 Cookie (登录或刷新后调用)
func (m *CookieManager) Set(c *gin.Context, value string) {
	// 设置 SameSite (必须单独调用 Gin 的方法)
	// 在跨站请求时，是否允许带上这个 Cookie
	c.SetSameSite(m.parseSameSite())

	c.SetCookie(
		m.cfg.Name,
		value,
		int(m.cfg.MaxAge), // 有效期, 7天
		m.cfg.Path,        // 关键：建议设为 /admin/auth/session/refresh 缩小暴露面
		m.cfg.Domain,      // 域名
		m.cfg.Secure,      // 是否仅 HTTPS
		m.cfg.HttpOnly,    // 是否禁止 JS 读取 (防 XSS)
	)
}

// 4. Get: 获取 Cookie (验证时调用)
func (m *CookieManager) Get(c *gin.Context) (string, error) {
	return c.Cookie(m.cfg.Name)
}

// 5. Remove: 清除 Cookie (登出时调用)
func (m *CookieManager) Remove(c *gin.Context) {
	// 删除的核心是：MaxAge 设为 -1
	// 注意：Path 和 Domain 必须与写入时完全一致
	c.SetCookie(
		m.cfg.Name,
		"",
		-1,
		m.cfg.Path,
		m.cfg.Domain,
		m.cfg.Secure,
		m.cfg.HttpOnly,
	)
}

// 内部辅助函数：转换 SameSite 字符串
func (m *CookieManager) parseSameSite() http.SameSite {
	switch m.cfg.SameSite {
	case "Lax":
		return http.SameSiteLaxMode
	case "Strict":
		return http.SameSiteStrictMode
	case "None":
		return http.SameSiteNoneMode
	default:
		return http.SameSiteDefaultMode
	}
}
