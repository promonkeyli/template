import type { AxiosError, AxiosRequestConfig } from "axios";
import { createAxiosFactory } from "@/services/core/factory.ts";
import { useAuthStore } from "@/stores/auth.ts";

declare const process: { env: Record<string, string> }; // process 类型声明

export const mainInstance = createAxiosFactory({
	baseURL: process.env.VITE_API_BASE_URL,
	token: useAuthStore.getState().token?.access_token,
	whiteList: ["/admin/auth/login", "/admin/auth/register", "/admin/auth/refresh"],
});

// -------------------------- Orval 适配器, 基于自定义的 axios 实例 --------------------------
export const customInstance = <T>(config: AxiosRequestConfig, options?: AxiosRequestConfig): Promise<T> => {
	return mainInstance({
		...config,
		...options,
	});
};
export type ErrorType<Error> = AxiosError<Error>;
export type BodyType<BodyData> = BodyData;
