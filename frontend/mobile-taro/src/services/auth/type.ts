/**
 * 登录请求参数
 */
export interface LoginParams {
    phone: string;
    password: string;
}

/**
 * 登录响应数据
 */
export interface LoginResponse {
  access_token: string;
  expires_at: number;
  refresh_token: string;
  uid: string;
}
