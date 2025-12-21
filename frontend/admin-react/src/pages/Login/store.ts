import { useNavigate } from "@tanstack/react-router";
import { useLogin as useLoginMutation } from "@/services/api/auth/auth";
import type { AuthLoginReq } from "@/services/api/model";
import { useAuthStore } from "@/stores/auth";

export const useLogin = () => {
	const navigate = useNavigate();
	const { setToken, setUserInfo } = useAuthStore();

	return useLoginMutation({
		mutation: {
			onSuccess: (response) => {
				console.log("response", response);
				const data = response.data;

				// 更新 auth store
				if (data) {
					setToken({
						access_token: data.access_token || "",
						refresh_token: data.refresh_token || "",
						expires_at: data.expires_at || 0,
						uid: data.uid || "",
					});
					setUserInfo({
						id: data.uid || "",
						username: data.uid || "",
						nickname: "",
						roles: [],
						permissions: [],
					});

					// 登录成功后重定向到首页
					navigate({ to: "/admin" });
				}
			},
			onError: (error) => {
				console.error("登录失败:", error);
			},
		},
	});
};

// 导出类型以供使用
export type { AuthLoginReq };
