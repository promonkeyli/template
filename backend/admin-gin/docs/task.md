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

- [ ] 示例：初始化项目结构
- [ ] 示例：Gin + GORM 初始化

---

## 📝 备注（Notes）
- 所有重要业务表采用双 ID（自增 id + uid）
- 所有中间表不使用双 ID
- 文档持续更新，不拆为多个文件
