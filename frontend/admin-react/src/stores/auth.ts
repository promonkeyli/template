/**
 * @description: 权限认证 store
 */
import { create } from "zustand";
import { createJSONStorage, persist } from "zustand/middleware";

// 用户信息类型定义
export interface UserInfo {
	id: string;
	username: string;
	nickname: string;
	roles: string[]; // 角色列表
	permissions?: string[]; // 权限列表（可选）
	[key: string]: any; // 允许扩展其他字段
}

// Token 相关类型
export interface TokenInfo {
	access_token: string;
	refresh_token: string;
	expires_at: number;
	uid: string;
}

// Store 状态类型
interface AuthState {
	// 状态
	token: TokenInfo | null;
	userInfo: UserInfo | null;
	isAuthenticated: boolean; // 是否已认证

	// 方法
	setToken: (token: TokenInfo) => void;
	setUserInfo: (user: UserInfo) => void;
	refreshAccessToken: (newToken: string) => void; // 单独刷新访问令牌
	logout: () => void; // 登出（清空状态）
	clearUserInfo: () => void; // 仅清空用户信息（保留token，如切换账号场景）
}

// 创建 Auth Store，使用 persist 中间件持久化到 localStorage
export const useAuthStore = create<AuthState>()(
	persist(
		(set, get) => ({
			// 初始状态
			token: null,
			userInfo: null,
			isAuthenticated: false,

		/**
		 * 设置令牌信息
		 * @param token - 令牌对象
		 */
		setToken: (token) => {
			set({
				token,
				isAuthenticated: !!token.access_token, // 有 access_token 则视为已认证
			});
		},

			/**
			 * 设置用户信息
			 * @param user - 用户信息对象
			 */
			setUserInfo: (user) => {
				set({ userInfo: user });
			},

		/**
		 * 单独刷新访问令牌（适用于令牌过期场景）
		 * @param newToken - 新的 access_token
		 */
		refreshAccessToken: (newToken) => {
			const currentToken = get().token;
			if (currentToken) {
				set({
					token: { ...currentToken, access_token: newToken },
					isAuthenticated: true,
				});
			}
		},

			/**
			 * 登出：清空所有认证相关状态
			 */
			logout: () => {
				set({
					token: null,
					userInfo: null,
					isAuthenticated: false,
				});
			},

			/**
			 * 清空用户信息（保留 token，用于切换账号等场景）
			 */
			clearUserInfo: () => {
				set({ userInfo: null });
			},
		}),
		{
			name: "auth-storage",
			storage: createJSONStorage(() => localStorage),
		},
	),
);
