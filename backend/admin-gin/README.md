# Mall API

基于 Go 和 Gin 构建的企业级 RESTful API 商城应用后端。

## 技术栈

*   **语言**: Go 1.25+
*   **Web 框架**: [Gin](https://github.com/gin-gonic/gin) - 高性能 HTTP Web 框架
*   **ORM**: [GORM](https://gorm.io/) - Go 语言优秀的 ORM 库
*   **数据库**: PostgreSQL
*   **认证**: JWT (Bearer Token)
*   **缓存/KV存储**: [go-redis](https://github.com/redis/go-redis) - 官方推荐的 Redis 客户端
*   **依赖注入**: [google/wire](https://github.com/google/wire) (编译期依赖注入，用于集中装配模块依赖)
*   **日志**: [log/slog](https://pkg.go.dev/log/slog) - Go 1.21+ 标准库结构化日志
*   **操作系统/工具**: Linux/macOS
*   **文档**: Swagger / OpenAPI (通过 `swaggo`)

## 项目结构（实际目录与路由）

本项目采用“按业务模块聚合”的分层结构，数据流向保持单向依赖：`Router -> Handler -> Service -> Repository -> Database`。

当前 admin 模块主要目录如下（与代码一致）：

```bash
mall-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── app/
│   │   └── admin/
│   │       ├── iam/
│   │       │   └── auth/
│   │       │       ├── dto.go
│   │       │       ├── handler.go
│   │       │       ├── model.go          # auth 自己的 Account + GORM 映射模型（与 user 解耦，但共用同表）
│   │       │       ├── repository.go
│   │       │       ├── service.go
│   │       │       └── router.go         # RegisterRouter(r, handler)
│   │       ├── user/
│   │       │   ├── constant.go
│   │       │   ├── dto.go
│   │       │   ├── handler.go            # user 常用 CRUD + swagger 注释
│   │       │   ├── model.go
│   │       │   ├── repository.go
│   │       │   ├── service.go
│   │       │   └── router.go             # RegisterRouter(r, handler)
│   │       └── wire/
│   │           ├── wire.go               #go:build wireinject（注入器声明）
│   │           └── wire_gen.go           # 生成文件（默认编译使用）
│   ├── database/
│   │   ├── postgre.go                    # InitDB（单数表名 SingularTable=true）
│   │   └── redis.go                      # InitRedis
│   ├── router/
│   │   └── router.go                     # 全局路由入口（admin 路由注册）
│   ├── wire/
│   │   ├── wire.go                       #go:build wireinject（全应用注入器声明）
│   │   └── wire_gen.go                   # 生成文件
│   └── pkg/
│       ├── http/                         # 统一响应、分页请求
│       └── mw/                           # JWT / cors / log / recover
└── api/
    └── openapi/                          # swaggo 生成的 swagger 文档输出目录
```

> 说明：`wire.go` 文件带有 `wireinject` build tag，默认会被编辑器/gopls 排除；查看实现请打开对应的 `wire_gen.go`。

## Admin 用户模块（/admin/user）接口

该模块用于后台“管理员用户”管理（与登录认证用户表共用）。

统一鉴权：所有 `/admin/user` 路由均需要 `Authorization: Bearer <access_token>`。

### 1) 获取用户列表（分页）

- **GET** `/admin/user`
- Query:
  - `page` (int, default 1)
  - `size` (int, default 10, max 100)
  - `role` (string, optional)
  - `keyword` (string, optional，匹配 uid/username/email)

### 2) 创建用户

- **POST** `/admin/user`
- Body: `CreateReq`
  - `username` (required)
  - `password` (required, 明文传入，服务端 bcrypt)
  - `email` (optional)
  - `role` (required，服务端按 `user.Role.IsValid()` 校验)

### 3) 更新用户

- **PUT** `/admin/user/{uid}`
- Body: `UpdateReq`
  - `email` (optional)
  - `role` (optional，服务端按 `user.Role.IsValid()` 校验)
  - `is_active` (optional, *bool，区分“不修改/修改为 false”)

### 4) 删除用户（软删除）

- **DELETE** `/admin/user/{uid}`
- 行为：设置 `is_deleted=true`（软删除）

## 开发最佳实践

### 1. 命名规范

*   **Go 语言命名**: 遵循 `CamelCase`（驼峰命名法）。
    *   **导出变量/函数**: 首字母大写 (e.g., `UserRepository`, `CreateUser`).
    *   **私有变量/函数**: 首字母小写 (e.g., `db`, `parseToken`).
    *   **接口**: 通常以 `er` 结尾，或明确表示其行为 (e.g., `Reader`, `Service`).
    *   **目录命名**: 推荐使用 **全小写**，尽量使用 **单个单词** (e.g., `user`, `config`)。
    *   **文件命名**: 推荐使用 **单个单词** 或 **下划线连接多个单词** (snake_case) (e.g., `handler.go`, `user_service.go`)。

### 2. RESTful API 规范

本项目严格遵循 RESTful 风格，仅使用以下四种 HTTP 方法。

*   **URL 设计**: 使用名词表达资源，路径 **全小写** (e.g., `/admin/user`)。
*   **方法使用**:

    | 方法 | 描述 | 适用场景 | 示例 |
    | :--- | :--- | :--- | :--- |
    | **GET** | 查询资源 | 获取列表、获取详情 | `GET /admin/user` |
    | **POST** | 创建资源 | 新增一条数据 | `POST /admin/user` |
    | **PUT** | 更新资源 | 修改现有数据 (全量或部分更新) | `PUT /admin/user/:uid` |
    | **DELETE** | 删除资源 | 删除一条数据 | `DELETE /admin/user/:uid` |

> **注意**: 本项目统一使用 `PUT` 处理更新操作。

### 3. 分层架构与数据流

本项目采用严格的分层架构，数据流向单向依赖：`Router -> Handler -> Service -> Repository -> Database`。

#### 3.1 各层职责划分

*   **Handler 层 (接口层)**:
    *   **职责**: 解析 HTTP 请求参数 (DTO)，进行基础参数校验，调用 Service 层方法，封装统一响应格式。
    *   **禁止**: 禁止包含复杂业务逻辑，禁止直接访问数据库。
    *   **依赖**: 依赖 `Service` 接口。

*   **Service 层 (业务层)**:
    *   **职责**: 核心业务逻辑实现，事务控制，数据组装，编排 Repository 层接口。
    *   **设计**: 定义并实现 Service 接口，保证可测试性。
    *   **依赖**: 依赖 `Repository` 接口。

*   **Repository 层 (仓储层)**:
    *   **职责**: 负责数据的持久化操作 (CRUD)，屏蔽底层存储细节 (GORM/SQL/Redis)。
    *   **设计**: 实现 Domain 层定义的 Repository 接口。
    *   **禁止**: 禁止包含业务逻辑。

*   **Model (数据模型)**:
    *   **职责**: 定义与数据库表一一对应的结构体 (Struct)，包含 GORM 标签。
    *   **位置**: 通常位于模块目录 `internal/app/admin/{module}/model.go`。

*   **DTO (数据传输对象)**:
    *   **职责**: 定义 API 输入参数 (Request) 和输出响应 (Response) 的结构。
    *   **目的**: 将 HTTP 接口契约与底层数据库模型解耦，防止 API 变更影响数据库，反之亦然。

### 4. 统一响应格式

所有 API 接口返回的数据结构保持一致，使用 `internal/pkg/http` 包中的辅助函数 `http.OK`（不分页）、`http.OKWithPage`（分页）以及 `http.Fail`（失败）进行处理。

#### 4.1 基础结构（不分页）

适用于单个资源查询、创建、更新、删除等操作。

```go
// 示例代码
http.OK(c, &http.OKOption{
    Data: user,
})
```

**响应示例**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "uid": "u123",
    "nickname": "TopG"
  }
}
```

#### 4.2 分页结构

适用于列表查询，包含分页元数据。

```go
// 示例代码
http.OKWithPage(c, &http.PageOption{
    Data:  users,
    Page:  1,
    Size:  10,
    Total: 100,
})
```

**响应示例**:
```json
{
  "code": 0,
  "message": "success",
  "data": [
    { "uid": "u1", "nickname": "User1" },
    { "uid": "u2", "nickname": "User2" }
  ],
  "page": 1,
  "size": 10,
  "total": 100
}
```

#### 4.3 错误响应

适用于业务逻辑错误或系统异常。

```go
// 示例代码
http.Fail(c, &http.FailOption{
    Code:    http.Unauthorized,
    Message: "Token 已过期或无效",
})
```

**响应示例**:
```json
{
  "code": 40001,
  "message": "Token 已过期或无效",
  "data": null
}
```

## 数据库最佳实践

*   **命名规范**:
    *   **表名**: 使用 **单数名词** + **全小写** + **下划线分隔**。例如: `user`, `product`, `user_profile`。
    *   **列名**: 使用 **全小写** + **下划线分隔** (snake_case)。例如: `created_at`, `user_id`, `status`。
    *   **主键**: 统一命名为 `id` (bigint/uuid)，业务主键可命名为 `uid` 或 `order_no`。
*   **外键**: 尽量在应用层维护关联关系，高并发场景下减少物理外键约束。

## Redis 最佳实践

### 1. 引入与初始化

Redis 连接初始化位于 `internal/database/redis.go`，配置位于 `internal/config/redis.go`。推荐使用 `go-redis` 库进行操作。

*   **配置**: 使用 `RedisConfig` 结构体管理地址、密码和 DB。
*   **连接**: 在 `main.go` 中初始化，并将其注入到 `Router` 中，最终传递给 `Service` 或 `Repository` 层使用。

### 2. 使用场景

Redis 应当仅用于缓存、临时数据存储或特定的高性能需求场景，**严禁**作为持久化主数据存储（除非有充分理由）。

*   **缓存 (Caching)**: 缓存热点数据（如用户信息、商品详情），减少数据库压力。必须设置过期时间 (TTL)。
*   **会话管理 (Session)**: 存储短期有效的 Token 白名单/黑名单。
*   **分布式锁 (Distributed Lock)**: 处理高并发下的资源竞争。
*   **限流 (Rate Limiting)**: 防止恶意刷接口。
*   **排行榜 (Leaderboard)**: 利用 ZSet 实现高性能排行榜。

### 3. Key 命名规范

类似于数据库表名，Redis Key 也需要规范命名以防止冲突和方便管理。

*   **格式**: `项目名:模块名:业务名:唯一标识` (使用冒号分隔)
*   **示例**:
    *   `mall:auth:token:u123` (用户 u123 的 Token)
    *   `mall:prod:detail:p888` (商品 p888 的详情缓存)
    *   `mall:sms:code:13800000000` (手机验证码)

### 4. 代码示例

在 Repository 层中使用 Redis：

```go
func (r *userRepository) GetUserCache(ctx context.Context, uid string) (*model.User, error) {
    key := fmt.Sprintf("mall:user:info:%s", uid)
    val, err := r.rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        return nil, nil // 缓存未命中
    } else if err != nil {
        return nil, err // 其它错误
    }
    
    // 反序列化
    var user model.User
    json.Unmarshal([]byte(val), &user)
    return &user, nil
}
```

## Git 协作规范

本项目采用 **Conventional Commits** 规范提交代码，格式为 `<type>(<scope>): <subject>`。

### Commit Types

*   `feat`: ✨ 新增功能 (A new feature)
*   `fix`: 🐛 修复 Bug (A bug fix)
*   `docs`: 📚 文档变更 (Documentation only changes)
*   `style`: 💎 代码格式调整 (Changes that do not affect the meaning of the code)
*   `refactor`: ♻️ 代码重构 (A code change that neither fixes a bug nor adds a feature)
*   `perf`: 🚀 性能优化 (A code change that improves performance)
*   `test`: 🚨 测试相关 (Adding missing tests or correcting existing tests)
*   `chore`: 🔧 构建/工具变动 (Changes to the build process or auxiliary tools)

### Subject 规范

1.  使用 **祈使句**，**一般现在时** (e.g., "change" not "changed" nor "changes")。
2.  首字母 **不要大写**。
3.  结尾 **不要句号** (.)。
4.  语言: 推荐使用 **英文**。

**示例**:
> `feat(user): implement login api`
> `fix(db): resolve connection timeout`
> `docs(readme): update project structure`

## 运行与配置

### 前置要求

*   Go 1.25+
*   PostgreSQL
*   Redis

### 启动项目

1.  **安装依赖**:
    ```bash
    go mod tidy     # 整理依赖 (添加缺失模块，移除未使用模块)
    go mod download # 下载依赖到本地缓存
    ```

2.  **运行应用**:
    使用 Makefile 快捷命令启动服务器：
    ```bash
    make run
    ```

3.  **生成文档**:
    更新并生成 Swagger API 文档：
    ```bash
    make swag
    ```
