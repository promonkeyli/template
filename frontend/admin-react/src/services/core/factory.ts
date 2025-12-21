/**
 * @description: axios 工厂函数，按配置生成 axios 实例
 */

import axios, { type AxiosInstance, type AxiosRequestConfig } from "axios";

export interface CreateAFOptions extends AxiosRequestConfig {
	// 请求白名单，不需要鉴权的
	whiteList?: string[];
	// token
	token?: string;
}

/**
 * @description Axios 实例工厂函数
 */
export function createAxiosFactory(options: CreateAFOptions = {}): AxiosInstance {
	const { whiteList = [], token, ...axiosConfig } = options;

	// 1. 基础配置
	const instance = axios.create({
		timeout: 60000, // axios timeout 中 使用ms，1000ms === 1s，默认设置1min
		headers: { "Content-Type": "application/json" },
		...axiosConfig, // 合并传入的 baseURL 等配置
	});

	// 2. 公共的请求拦截器配置,主要是 筛选请求白名单
	instance.interceptors.request.use(
		(request) => {
			if (request?.url && whiteList.includes(request?.url)) {
				return request;
			}
			if (token) {
				// 标准 Authorization 授权字段 添加
				request.headers.Authorization = `Bearer ${token}`;
			}
			return request;
		},
		(error) => {
			console.log("请求参数错误", error);
			return Promise.reject(error);
		},
	);

	// 3. 公共的响应拦截器配置
	instance.interceptors.response.use(
		(response) => {
			const { status } = response;
			// 只要 HTTP 200 且 业务 code 200 就返回 data
			if (status === 200 && response.data.code === 200) {
				return response.data;
			} else {
				console.log("请求失败", response.data);
				return Promise.reject(response.data);
			}
		},
		(error) => {
			return Promise.reject(error);
		},
	);

	return instance;
}
