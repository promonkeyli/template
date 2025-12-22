import type { AxiosError, AxiosRequestConfig, InternalAxiosRequestConfig } from "axios";
import { createAxiosFactory } from "@/services/core/factory";
import { useAuthStore } from "@/stores/auth";
import { refreshToken } from "@/services/api/auth/auth";

// 1. 规范环境变量读取 (Vite 标准)
const BASE_URL = import.meta.env.VITE_API_BASE_URL;

// 2. 定义业务状态码 (根据你后端的实际定义修改)
const ResultEnum = {
	SUCCESS: 200,
	EXPIRE: 401, // Token 过期
	ERROR: 500,
};

// 3. 定义白名单
const WHITE_LIST = ["/admin/auth/login"];

// 4. 并发控制变量
let isRefreshing = false;
let requestsQueue: Array<(token: string) => void> = [];

/**
 * @description 主业务 Axios 实例
 */
export const httpInstance = createAxiosFactory({
	baseURL: BASE_URL,
	interceptors: {
		// ---------------- 请求拦截器 ----------------
		request: (config) => {
			const url = config.url || "";
			// 白名单检查
			if (WHITE_LIST.some((path) => url.includes(path))) {
				return config;
			}

			// 实时获取 Token
			const token = useAuthStore.getState().token?.access_token;
			if (token) {
				config.headers.Authorization = `Bearer ${token}`;
			}
			return config;
		},

		// ---------------- 响应拦截器 ----------------
		response: {
			// 🔥 核心修改：逻辑移到 success，因为后端返回 HTTP 200
			success: async (response) => {
				const { data, config } = response;
				// 类型断言，确保能访问 _retry 属性
				const originalRequest = config as InternalAxiosRequestConfig & { _retry?: boolean };

				// A. 业务成功
				if (data.code === ResultEnum.SUCCESS) {
					return data; // 根据需要返回 data 或 data.data
				}

				// B. Token 过期 (业务码 401)
				if (data.code === ResultEnum.EXPIRE && !originalRequest._retry) {
					// 1. 如果正在刷新，加入队列等待
					if (isRefreshing) {
						return new Promise((resolve) => {
							requestsQueue.push((token) => {
								originalRequest.headers.Authorization = `Bearer ${token}`;
								resolve(httpInstance(originalRequest));
							});
						});
					}

					// 2. 开启刷新锁
					originalRequest._retry = true;
					isRefreshing = true;

					try {
						// 调用 refresh-client 刷新
						// 假设 executeRefreshToken 返回的数据结构就是 Token 对象
						const newTokenData: any = await refreshToken();

						// 更新 Store
						useAuthStore.getState().setToken(newTokenData);

						// 唤醒队列中的请求
						requestsQueue.forEach((cb) => cb(newTokenData.access_token));
						requestsQueue = [];

						// 重试当前请求
						originalRequest.headers.Authorization = `Bearer ${newTokenData.access_token}`;
						return httpInstance(originalRequest);
					} catch (refreshErr) {
						// 刷新失败，清空队列并登出
						requestsQueue = [];
						useAuthStore.getState().logout();
						window.location.href = "/login";

						return Promise.reject(refreshErr);
					} finally {
						isRefreshing = false;
					}
				}

				// C. 其他业务错误 (虽然 HTTP 200，但业务失败)
				// 手动 reject，这样 Orval 的 onError 或 try/catch 才能捕获到
				return Promise.reject(data);
			},

			// 🔥 网络层面的错误 (HTTP Status != 2xx)
			// 例如：超时、断网、502 Bad Gateway
			error: (error) => {
				console.error("网络请求异常:", error);
				return Promise.reject(error);
			},
		},
	},
});

/**
 * @description Orval 适配器
 */
export const httpClient = <T>(
	config: AxiosRequestConfig,
	options?: AxiosRequestConfig
): Promise<T> => {
	return httpInstance({
		...config,
		...options,
	});
};

// ---------------- 类型定义导出 ----------------
export type ErrorType<Error> = AxiosError<Error>;
export type BodyType<BodyData> = BodyData;