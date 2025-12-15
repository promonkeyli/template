/**
 * @description: 基于Axios的请求简易封装
 */
import axios, { type AxiosResponse } from "axios";
import type {
	ApiResponse,
	IAxiosRequestConfig,
	IInternalAxiosRequestConfig,
} from "@/services/types"; // 导入类型扩展
import { useAuthStore } from "@/stores/auth.ts";

const axiosInstance = axios.create({
	baseURL: import.meta.env.VITE_API_BASE_URL,
	timeout: 1000,
	headers: {
		"Content-Type": "application/json;charset=utf-8",
	},
});

// -------------------------- 请求拦截器 --------------------------
axiosInstance.interceptors.request.use(
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
axiosInstance.interceptors.response.use((response: AxiosResponse) => {
	const { status } = response;
	if (status === 200 && response.data.code === 200) {
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
		return axiosInstance.request<ApiResponse<T>>(config) as any;
	},
	get: <T = any>(
		url: string,
		config?: IAxiosRequestConfig,
	): Promise<ApiResponse<T>> => {
		return axiosInstance.get<ApiResponse<T>>(url, config) as any;
	},
	post: <T = any>(
		url: string,
		data?: any,
		config?: IAxiosRequestConfig,
	): Promise<ApiResponse<T>> => {
		return axiosInstance.post<ApiResponse<T>>(url, data, config) as any;
	},
	put: <T = any>(
		url: string,
		data?: any,
		config?: IAxiosRequestConfig,
	): Promise<ApiResponse<T>> => {
		return axiosInstance.put<ApiResponse<T>>(url, data, config) as any;
	},
	delete: <T = any>(
		url: string,
		config?: IAxiosRequestConfig,
	): Promise<ApiResponse<T>> => {
		return axiosInstance.delete<ApiResponse<T>>(url, config) as any;
	},
};

export default request;
