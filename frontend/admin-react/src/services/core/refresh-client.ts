import type { AxiosRequestConfig } from "axios";
import { createAxiosFactory } from "@/services/core/factory";

declare const process: { env: Record<string, string> };

/**
 * @description 刷新专用 Axios 实例 (Clean Instance)
 * @features
 * 1. 没有任何 Token 注入逻辑 (防止闭包/死循环)
 * 2. 没有任何 401 重试逻辑
 * 3. 仅用于 refresh-token 接口
 */
export const refreshInstance = createAxiosFactory({
    baseURL: process.env.VITE_API_BASE_URL,
    interceptors: {
        response: {
            success: (res) => {
                if (res.status === 200 && res.data.code === 401) {
                    // 业务码401，跳转登录页
                    window.location.href = "/login";
                }
                return res.data;
            },
            error: (error) => {
                return Promise.reject(error);
            },
        },
    },
});

/**
 * @description Orval 适配器：刷新专用
 * @usage 在 orval.config.ts 中针对 refreshToken 接口配置此 mutator
 */
export const refreshClient = <T>(
    config: AxiosRequestConfig,
    options?: AxiosRequestConfig
): Promise<T> => {
    return refreshInstance({
        ...config,
        ...options,
    });
};