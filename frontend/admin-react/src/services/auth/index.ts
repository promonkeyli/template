/**
 * @description: 认证相关API
 */

import request from "@/services/request";
import type { LoginReq, LoginRes } from "./type";

/**
 * 用户名密码登录
 * @param data 登录参数
 * @returns 登录响应数据
 */
export const fetchLogin = (data: LoginReq) => {
	return request.post<LoginRes>("/admin/auth/login", data, { isSkipAuth: true });
};
