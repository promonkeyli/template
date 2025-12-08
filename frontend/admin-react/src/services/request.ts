/**
 * @description: 基于Axios的请求简易封装
 */

import axios, { type AxiosResponse } from "axios";
import { useAuthStore } from "@/stores/auth.ts";
import type { ApiResponse, IAxiosRequestConfig, IInternalAxiosRequestConfig } from "@/services/types"; // 导入类型扩展

// axios 实例
const requestInstance = axios.create({
	baseURL: import.meta.env.VITE_API_BASE_URL,
	timeout: 1000,
	headers: {
		"Content-Type": "application/json;charset=utf-8",
	},
});

// -------------------------- 请求拦截器 --------------------------
requestInstance.interceptors.request.use(
	(config: IInternalAxiosRequestConfig) => {
		// 判断是否跳过鉴权，默认为false（需要鉴权）
		if (config.isSkipAuth) {
			return config;
		}

	// 添加token
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
requestInstance.interceptors.response.use((response: AxiosResponse) => {
	const { status } = response;
	if (status === 200 && response.data.code === 0) {
		return response.data; // 直接返回业务数据
	} else {
		console.log("请求失败", response.data);
		return Promise.reject(response.data);
	}	
});

/**
 * @description: 请求函数: 为了ts考虑, 提供更具体的类型
 */
const request = {
	request: <T = any>(config: IAxiosRequestConfig): Promise<ApiResponse<T>> => {
		return requestInstance.request<ApiResponse<T>>(config) as any;
	},
	get: <T = any>(url: string, config?: IAxiosRequestConfig): Promise<ApiResponse<T>> => {
		return requestInstance.get<ApiResponse<T>>(url, config) as any;
	},
	post: <T = any>(url: string, data?: any, config?: IAxiosRequestConfig): Promise<ApiResponse<T>> => {
		return requestInstance.post<ApiResponse<T>>(url, data, config) as any;
	},
	put: <T = any>(url: string, data?: any, config?: IAxiosRequestConfig): Promise<ApiResponse<T>> => {
		return requestInstance.put<ApiResponse<T>>(url, data, config) as any;
	},
	delete: <T = any>(url: string, config?: IAxiosRequestConfig): Promise<ApiResponse<T>> => {
		return requestInstance.delete<ApiResponse<T>>(url, config) as any;
	}
}

export default request;
