import http from '@/utils/http';
import type { LoginParams, LoginResponse } from './type';

/**
 * 用户登录
 * @param params 登录参数
 * @returns 登录响应数据
 */
export function fetchLogin(params: LoginParams) {
    return http.post<LoginResponse>('/v1/auth/login/phone', params, { isSkipAuth: true });
}

/**
 * 用户登出
 */
export function logout() {
    return http.post<void>('/logout');
}
