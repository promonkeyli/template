/**
 * @description: axios 工厂函数，按配置生成 axios 实例
 */

import axios, {
	type AxiosInstance,
	type AxiosRequestConfig,
	type AxiosResponse,
	type InternalAxiosRequestConfig,
} from "axios";

// 定义拦截器回调类型
type RequestInterceptor = (config: InternalAxiosRequestConfig) => InternalAxiosRequestConfig | Promise<InternalAxiosRequestConfig>;
type ResponseSuccessInterceptor = (response: AxiosResponse) => any; // 允许返回解包后的数据
type ResponseErrorInterceptor = (error: any) => any;

export interface CreateAFOptions extends AxiosRequestConfig {
	// 允许调用者传入自定义拦截器
	interceptors?: {
		request?: RequestInterceptor;
		requestError?: ResponseErrorInterceptor;
		response?: {
			success?: ResponseSuccessInterceptor;
			error?: ResponseErrorInterceptor;
		};
	};
}

/**
 * @description 通用 Axios 实例工厂
 * 只负责创建实例和注册拦截器，不包含具体业务逻辑
 */
export function createAxiosFactory(options: CreateAFOptions = {}): AxiosInstance {
	const { interceptors, ...axiosConfig } = options;

	// 1. 基础配置
	const instance = axios.create({
		timeout: 60000,
		headers: { "Content-Type": "application/json" },
		...axiosConfig,
	});

	// 2. 注册请求拦截器
	if (interceptors?.request) {
		instance.interceptors.request.use(
			interceptors.request,
			interceptors.requestError || ((error) => Promise.reject(error))
		);
	}

	// 3. 注册响应拦截器
	if (interceptors?.response) {
		instance.interceptors.response.use(
			interceptors.response.success || ((res) => res), // 默认原样返回
			interceptors.response.error || ((err) => Promise.reject(err)) // 默认直接抛出
		);
	}

	return instance;
}
