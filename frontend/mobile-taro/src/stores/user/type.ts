/**
 * 用户 Store 相关的 TypeScript 类型定义
 */
import {LoginResponse} from "@/services/auth/type";

export interface UserState {
  /**
   * 用户信息
   */
  userInfo: any | null;
  /**
   * Token信息
   */
  tokenInfo: LoginResponse | null;
  /**
   * 设置用户信息
   * @param userInfo 用户信息
   */
  setUserInfo: (userInfo: any) => void;
  /**
   * 设置Token信息
   * @param tokenInfo
   */
  setTokenInfo: (tokenInfo: LoginResponse) => void;
  /**
   * 清除用户状态（用于退出登录）
   */
  clearUser: () => void;
}
