package user

// Define Role type based on string
// 定义 Role 类型，底层是 string，方便数据库存储和 JSON 序列化
type Role string

const (
	// ==========================================================
	// 1. 系统管理层
	// ==========================================================

	// RoleSuperAdmin 超级管理员
	// 权限：上帝视角，拥有所有权限，不可被删除，通常只有 1-2 个
	RoleSuperAdmin Role = "super_admin"

	// RoleAdmin 普通管理员/店长
	// 权限：仅次于超级管理员，管理日常运营，但不能修改系统底层配置
	RoleAdmin Role = "admin"

	// ==========================================================
	// 2. 业务运营层
	// ==========================================================

	// RoleProductManager 商品管理员
	// 权限：商品上架/下架、编辑详情、库存预警、分类管理、品牌管理
	RoleProductManager Role = "product_manager"

	// RoleMarketing 营销/运营专员
	// 权限：装修店铺首页、发布公告、创建优惠券、管理秒杀活动、广告位管理
	RoleMarketing Role = "marketing"

	// ==========================================================
	// 3. 履约与服务层
	// ==========================================================

	// RoleOrderManager 订单/仓储专员
	// 权限：查看待发货订单、打印面单、填写物流号、处理退货入库
	// 注意：通常看不到订单的“成本价”或“利润”
	RoleOrderManager Role = "order_manager"

	// RoleCustomerService 客服专员
	// 权限：查看订单详情、查看物流、回复工单
	// 注意：敏感信息（手机号）需脱敏，通常只有只读权限
	RoleCustomerService Role = "customer_service"

	// ==========================================================
	// 4. 财务层
	// ==========================================================

	// RoleFinance 财务专员
	// 权限：查看营收报表、资金流水、审核退款打款、开发票
	RoleFinance Role = "finance"
)

// RoleLabels 用于前端展示的中文名称映射
// 可以通过 API 返回给前端，用于生成下拉选择框
var RoleLabels = map[Role]string{
	RoleSuperAdmin:      "超级管理员",
	RoleAdmin:           "系统管理员",
	RoleProductManager:  "商品管理员",
	RoleMarketing:       "营销运营",
	RoleOrderManager:    "订单/仓储",
	RoleCustomerService: "客服专员",
	RoleFinance:         "财务专员",
}

// IsValid 校验角色是否合法（防止前端传乱码）
func (r Role) IsValid() bool {
	_, ok := RoleLabels[r]
	return ok
}

// String 实现 Stringer 接口，打印时自动显示字符串
func (r Role) String() string {
	return string(r)
}

// Label 获取中文名称
func (r Role) Label() string {
	if label, ok := RoleLabels[r]; ok {
		return label
	}
	return "未知角色"
}
