import axios, { type AxiosRequestConfig, type AxiosResponse } from "axios";
import type { IInternalAxiosRequestConfig } from "@/services/types";
import { useAuthStore } from "@/stores/auth";

declare const process: { env: Record<string, string> }; // process 类型声明

const BASE_URL = process.env.VITE_API_BASE_URL || "";

export const axiosInstance = axios.create({
	baseURL: BASE_URL,
	timeout: 3000,
	headers: {
		"Content-Type": "application/json;charset=utf-8",
	},
});

// -------------------------- 请求拦截器 --------------------------
axiosInstance.interceptors.request.use(
	(config: IInternalAxiosRequestConfig) => {
		// 你的自定义属性处理
		if (config.isSkipAuth) {
			return config;
		}

		// Zustand 在组件外使用的方式
		const token = useAuthStore.getState().token;
		if (token) {
			config.headers.Authorization = `Bearer ${token.access_token}`;
		}
		return config;
	},
	(error) => {
		console.log("请求参数错误", error);
		return Promise.reject(error);
	},
);

// -------------------------- 响应拦截器 --------------------------
axiosInstance.interceptors.response.use(
	(response: AxiosResponse) => {
		const { status } = response;
		// 这里你之前的逻辑是：只要 HTTP 200 且 业务 code 200 就返回 data
		// 注意：这意味着 customInstance 拿到的已经是 response.data (Body) 了
		if (status === 200 && response.data.code === 200) {
			return response.data;
		} else {
			console.log("请求失败", response.data);
			return Promise.reject(response.data);
		}
	},
	(error) => {
		// 建议增加 HTTP 状态码处理 (401, 403, 500等)
		return Promise.reject(error);
	},
);

// -------------------------- Orval 核心适配器 --------------------------

/**
 * Orval 调用的自定义请求函数
 * Orval 会把 swagger 中的 url, method, params 等参数传给 config
 * 把额外的 options (如 signal) 传给 options
 */
export const customInstance = <T>(
	config: AxiosRequestConfig,
	options?: AxiosRequestConfig,
): Promise<T> => {
	const source = axios.CancelToken.source();

	const promise = axiosInstance({
		...config,
		...options,
		cancelToken: source.token,
	}).then((res) => res); // 拦截器里已经返回了 response.data，所以这里直接返回 res 即可

	// @ts-ignore 添加 cancel 方法，以支持 React Query 的取消请求功能
	promise.cancel = () => {
		source.cancel("Query was cancelled");
	};

	// 强制转换为 T，因为拦截器已经解包了 Response
	return promise as Promise<T>;
};

// 你可以保留这个类型导出，或者让 Orval 自动生成
export type ErrorType<Error> = AxiosResponse<Error>;
