/**
 * @description: 认证相关的类型定义
 */

/**
 * 用户名密码登录请求参数
 */
export interface LoginReq {
	username: string;
	password: string;
}

/**
 * 登录响应数据
 */
export interface LoginRes {
	access_token: string;
	refresh_token: string;
	expires_at: number;
	uid: string;
}
