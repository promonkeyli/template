# logger 使用示例

本文档演示 `pkg/logger` 中以下三个文件在真实 Web 项目中的使用方式：

- `config.go`
- `context.go`
- `fields.go`

---

## 一、config.go 使用示例（统一初始化配置）

### 场景

- 启动时根据 **Gin mode** 初始化 logger
- 控制日志级别、环境、服务名

### 代码示例（main.go）

```go
func main() {
	gin.SetMode(gin.ReleaseMode)

	logger.Init(logger.Config{
		Service: "admin-api",
		Env:     "prod",
		Level:   slog.LevelInfo,
	})

	r := gin.New()
	r.Use(middleware.Gin())

	r.Run()
}
```

### 价值

- 日志初始化**参数集中**
- 后续可从环境变量或配置文件读取
- `Config` 成为基础设施稳定入口

---

## 二、context.go 使用示例（跨层传递 logger）

### 场景

- Handler → Service → Repository
- 每一层都能拿到**同一个 request logger**
- 自动带 `request_id`

### Gin 中间件（注入 logger）

```go
func Gin() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := logger.Default().With(
			logger.FieldRequestID, uuid.NewString(),
		)

		ctx := logger.WithContext(c.Request.Context(), l)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
```

### Handler 层

```go
func (h *Handler) CreateUser(c *gin.Context) {
	log := logger.FromContext(c.Request.Context())

	log.Info("create user request received")

	_ = h.service.CreateUser(c.Request.Context())
}
```

### Service 层

```go
func (s *Service) CreateUser(ctx context.Context) error {
	log := logger.FromContext(ctx)

	log.Debug("creating user in service layer")

	return nil
}
```

### Repository 层

```go
func (r *Repo) InsertUser(ctx context.Context) error {
	log := logger.FromContext(ctx).With("table", "users")

	log.Info("insert user into database")

	return nil
}
```

### 价值

- 不依赖 Gin
- request_id 自动贯穿全链路
- 体现 context.go 核心价值

---

## 三、fields.go 使用示例（字段统一规范）

### 场景

- 日志需要被 ELK / Loki / ClickHouse 分析
- 字段名必须稳定、统一

### fields.go

```go
const (
	FieldRequestID = "request_id"
	FieldUserID    = "user_id"
	FieldOrderID   = "order_id"
)
```

### 使用示例

```go
log.Info(
	"user created",
	logger.FieldUserID, userID,
	logger.FieldRequestID, reqID,
)
```

### 价值

- 统一字段名
- 避免多人开发字段命名冲突
- 日志分析工具可直接按字段查询