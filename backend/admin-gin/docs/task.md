# 商城后台开发任务规划（Tasks）

---

## 🟢 进行中（In Progress）

### Staff 管理模块
- [ ] staff 表结构定义（双 ID）
- [ ] staff CRUD 接口
- [ ] staff 列表分页
- [ ] staff 新增 / 编辑 / 删除
- [ ] 权限点：staff:list / staff:create / staff:update / staff:delete

---

## 🔵 下一个任务（Next）

### 商品（Product）模块
- [ ] product 表结构
- [ ] 商品编号（product_sn）规范
- [ ] 商品CRUD
- [ ] 商品图片上传（OSS）
- [ ] 前端商品列表页

### 权限系统（RBAC）
- [ ] role 表结构
- [ ] permission 表结构
- [ ] staff_role / role_permission 中间表
- [ ] 登录时加载权限
- [ ] 后端路由拦截

---

## 🟡 待办（Todo）

### 订单（Order）模块
- [ ] 订单列表
- [ ] 订单详情
- [ ] 订单状态流转

### 用户（User）模块（前台会员）
- [ ] user 表（双 ID）
- [ ] 用户登录 / 注册接口

---

## 🔴 已完成（Done）
（完成后从上面移动到这里）

### 工程化/基础设施
- [x] 引入 google/wire，集中装配依赖（DB/Redis/模块 Handler/路由注册）
- [x] admin 路由注册从“模块内 New”调整为“外部注入 handler”（RegisterRouter(r, handler)）

### Admin 用户模块（后台管理员 user）
- [x] user 列表分页接口（支持 role / keyword 过滤）
- [x] user 创建接口（bcrypt 加密密码）
- [x] user 更新接口（按 uid，支持 email / role / is_active）
- [x] user 删除接口（按 uid，软删除 is_deleted=true）
- [x] 为 user 模块接口补充 Swagger 注释

### Auth 模块
- [x] auth 与 user 解耦（共用同表，但 auth 不直接依赖 user.User 模型）

- [ ] 示例：初始化项目结构
- [ ] 示例：Gin + GORM 初始化

---

## 📝 备注（Notes）
- 所有重要业务表采用双 ID（自增 id + uid）
- 所有中间表不使用双 ID
- 文档持续更新，不拆为多个文件
