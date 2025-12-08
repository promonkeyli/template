import type { AxiosRequestConfig, InternalAxiosRequestConfig } from "axios";

/**
 * @description: 扩展Axios配置，添加自定义字段（用于API调用）
 */
export interface IAxiosRequestConfig extends AxiosRequestConfig {
	/** 是否跳过鉴权: 默认为false，需要鉴权 */
	isSkipAuth?: boolean;
}

/**
 * @description: 扩展Axios内部配置（用于拦截器）
 */
export interface IInternalAxiosRequestConfig extends InternalAxiosRequestConfig {
	/** 是否跳过鉴权: 默认为false，需要鉴权 */
	isSkipAuth?: boolean;
}

/**
 * @description: 请求响应类型
 */
export interface ApiResponse<T = any> {
	code: number;
	message: string;
	data: T;
}
