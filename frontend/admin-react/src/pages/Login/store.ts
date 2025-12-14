import { useMutation } from '@tanstack/react-query'
import { useNavigate } from '@tanstack/react-router'
import { useAuthStore } from '@/stores/auth'
import type { LoginReq } from '@/services/auth/type'
import { fetchLogin } from '@/services/auth'

export const useLogin = () => {
  const navigate = useNavigate()
  const { setToken, setUserInfo } = useAuthStore()

  return useMutation({
    mutationFn: (credentials: LoginReq) => fetchLogin(credentials),
    onSuccess: (response) => {
      console.log('response', response)
      const data = response.data
      // 更新 auth store
      setToken({
        access_token: data.access_token,
        refresh_token: data.refresh_token,
        expires_at: data.expires_at,
        uid: data.uid
      })
      setUserInfo({
        id: data.uid,
        username: data.uid,
        nickname: '',
        roles: [],
        permissions: []
      })

      // 登录成功后重定向到首页
      navigate({ to: '/admin' })
    },
    onError: (error) => {
      console.error('登录失败:', error)
    }
  })
}
