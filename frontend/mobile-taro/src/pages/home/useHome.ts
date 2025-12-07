import { useState } from 'react';
import Taro from '@tarojs/taro';
import { login, logout } from '@/services/user';
import useUserStore from '@/stores/user';

export default function useHomeStore() {
    const [username, setUsername] = useState('15984093508');
    const [password, setPassword] = useState('ly15984093508');
    const [loading, setLoading] = useState(false);

    const { userInfo, setUserInfo, setTokenInfo, clearUser } = useUserStore();

    const handleLogin = async () => {
        if (!username || !password) {
            Taro.showToast({ title: '请输入用户名和密码', icon: 'none' });
            return;
        }

        setLoading(true);
        try {
            const res = await login({ phone: username, password });
            // 假设 res.data 包含 userInfo 和 tokenInfo
            // 这里根据实际接口返回结构进行调整，暂时模拟
            const mockTokenInfo = {
                accessToken: 'mock_access_token',
                refreshToken: 'mock_refresh_token',
                expiresIn: 3600,
            };

            // 注意：这里假设 login 返回的数据结构中直接包含了用户信息
            // 如果实际接口不同，请根据 type.ts 调整
            setUserInfo(res.data.userInfo || { id: 1, username, avatar: '', nickname: '管理员' });
            setTokenInfo(res.data.tokenInfo || mockTokenInfo);

            Taro.showToast({ title: '登录成功', icon: 'success' });
        } catch (error) {
            console.error('登录失败', error);
            // http 拦截器通常会处理错误提示，这里可以根据需要补充
        } finally {
            setLoading(false);
        }
    };

    const handleLogout = async () => {
        try {
            await logout();
        } catch (error) {
            console.error('登出失败', error);
        } finally {
            clearUser();
            Taro.showToast({ title: '已退出登录', icon: 'none' });
        }
    };

    return {
        username,
        setUsername,
        password,
        setPassword,
        loading,
        userInfo,
        handleLogin,
        handleLogout,
    };
}
