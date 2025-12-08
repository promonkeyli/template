/**
 * @description: 认证相关API
 */

import request from "@/services/request";
import type { PhoneLoginReq, PhoneLoginRes } from "./type";

/**
 * 手机号密码登录
 * @param data 登录参数
 * @returns 登录响应数据
 */
export const phoneLogin = (data: PhoneLoginReq) => {
	return request.post<PhoneLoginRes>( "/v1/auth/login/phone", data, { isSkipAuth: true } );
};
