/**
 * @description: 认证相关的类型定义
 */

/**
 * 手机号密码登录请求参数
 */
export interface PhoneLoginReq {
	phone: string;
	password: string;
}

/**
 * 登录响应数据
 */
export interface PhoneLoginRes {
	access_token: string;
	refresh_token: string;
	expires_at: number;
	uid: string;
}
