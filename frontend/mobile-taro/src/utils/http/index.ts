import Taro from '@tarojs/taro';
import axios, { AxiosResponse } from 'axios';
import taroAdapter from './adapter';
import useUserStore from '@/stores/user';
// import { TokenInfo } from '@/stores/user/type';
import type { ApiResponse, IAxiosRequestConfig, IInternalAxiosRequestConfig } from "@/utils/http/type"; // 导入类型扩展

declare module 'axios' {
  export interface AxiosRequestConfig<D = any> {
    isSkipAuth?: boolean;
  }
  export interface InternalAxiosRequestConfig<D = any> {
    isSkipAuth?: boolean;
  }
}

const instance = axios.create({
  baseURL: process.env.TARO_APP_BASE_URL || '', // 从编译时注入的常量获取 baseURL
  timeout: 10000,
  adapter: taroAdapter,
});

// 是否正在刷新 token
// let isRefreshing = false;
// // 重试队列，每一项是一个 resolve 函数
// let requests: ((token: string) => void)[] = [];
//

// 请求拦截器
instance.interceptors.request.use(
  (config: IInternalAxiosRequestConfig) => {
    console.log('config', config);
    if (!config.isSkipAuth) {
      const tokenInfo = useUserStore.getState().tokenInfo as any;
      if (tokenInfo?.accessToken) {
        config.headers = config.headers || {};
        config.headers.Authorization = `Bearer ${tokenInfo.accessToken}`;
      }
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
instance.interceptors.response.use(
  (response: AxiosResponse) => {
    const { code } = response.data as ApiResponse;
    console.log('【原始响应】', response);

    // http code 200 ， 业务 code 0 同时满足表示，是成功的响应，否则全是失败
    if (response.status === 200 && code === 0) {
      return response.data
    } else {
      // TODO 根据 code 处理双token过期
      return Promise.reject(response.data)
    }

    // // 处理业务错误
    // if (response.status !== 401) {
    //   Taro.showToast({
    //     title: message || '请求失败',
    //     icon: 'none',
    //   });
    //   return Promise.reject(new Error(message || 'Error'));
    // }
    //
    // // 处理 401 Token 过期
    // const config = response.config as InternalAxiosRequestConfig;
    //
    // if (!isRefreshing) {
    //   isRefreshing = true;
    //   const refreshToken = useUserStore.getState().tokenInfo?.refreshToken;
    //
    //   if (!refreshToken) {
    //     return handleLoginExpired();
    //   }
    //
    //   // 尝试刷新 Token
    //   return refreshTokenFunc(refreshToken)
    //     .then((newTokenInfo) => {
    //       useUserStore.getState().setTokenInfo(newTokenInfo);
    //       config.headers = config.headers || {};
    //       config.headers.Authorization = `Bearer ${newTokenInfo.accessToken}`;
    //
    //       // 执行队列中的请求
    //       requests.forEach((cb) => cb(newTokenInfo.accessToken));
    //       requests = [];
    //
    //       return instance(config);
    //     })
    //     .catch(() => {
    //       return handleLoginExpired();
    //     })
    //     .finally(() => {
    //       isRefreshing = false;
    //     });
    // } else {
    //   // 正在刷新，将请求加入队列
    //   return new Promise((resolve) => {
    //     requests.push((token) => {
    //       config.headers.Authorization = `Bearer ${token}`;
    //       resolve(instance(config));
    //     });
    //   });
    // }
  },
  (error) => {
    // 处理网络错误等
    Taro.showToast({
      title: error.message || '网络异常',
      icon: 'none',
    });
    return Promise.reject(error);
  }
);

// // 模拟刷新 Token 的方法，实际应调用后端接口
// async function refreshTokenFunc(_refreshToken: string): Promise<TokenInfo> {
//   // 这里应该发送请求到后端刷新 token
//   // const res = await instance.post('/refreshToken', { refreshToken });
//   // return res.data.data;
//
//   // 模拟
//   return new Promise((resolve) => {
//     setTimeout(() => {
//       resolve({
//         accessToken: 'new_mock_access_token',
//         refreshToken: 'new_mock_refresh_token',
//         expiresIn: 7200,
//         refreshExpiresIn: 2592000
//       });
//     }, 1000);
//   })
// }
//
// function handleLoginExpired() {
//   useUserStore.getState().clearUser();
//   Taro.showToast({
//     title: '登录已过期，请重新登录',
//     icon: 'none',
//   });
//   // 跳转登录页
//   Taro.navigateTo({ url: '/pages/login/index' });
//   return Promise.reject(new Error('Login expired'));
// }

/**
 * @description: 请求函数: 为了ts考虑, 提供更具体的类型
 */
const request = {
	request: <T = any>(config: IAxiosRequestConfig): Promise<ApiResponse<T>> => {
    return instance.request<ApiResponse<T>>(config) as any;
	},
	get: <T = any>(url: string, config?: IAxiosRequestConfig): Promise<ApiResponse<T>> => {
		return instance.get<ApiResponse<T>>(url, config) as any;
	},
	post: <T = any>(url: string, data?: any, config?: IAxiosRequestConfig): Promise<ApiResponse<T>> => {
		return instance.post<ApiResponse<T>>(url, data, config) as any;
	},
	put: <T = any>(url: string, data?: any, config?: IAxiosRequestConfig): Promise<ApiResponse<T>> => {
		return instance.put<ApiResponse<T>>(url, data, config) as any;
	},
	delete: <T = any>(url: string, config?: IAxiosRequestConfig): Promise<ApiResponse<T>> => {
		return instance.delete<ApiResponse<T>>(url, config) as any;
	}
}

export default request;
