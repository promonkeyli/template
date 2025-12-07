import {useEffect, useRef, useState} from 'react'
import {useRequest} from "ahooks";
import {fetchLogin} from "@/services/auth";
import Taro from "@tarojs/taro";
import useUserStore from "@/stores/user";

export const useLogin = () => {
  const [phoneNumber, setPhoneNumber] = useState('15984093508')
  const [password, setPassword] = useState('ly15984093508')

  const timerRef = useRef<NodeJS.Timeout | null>(null)

  const { loading, run: loginRun } = useRequest(fetchLogin, {
    manual: true,
    onSuccess: (res) => {
      // 存储 token 信息
      useUserStore.getState().setTokenInfo(res.data)
      // 跳转主页
      Taro.showToast({
        title: "登陆成功",
        icon: "none",
        duration: 2000,
      })
      timerRef.current = setTimeout(() => {
        Taro.switchTab({
          url: "/pages/home/index",
        })
      }, 1000)
    },
    onError: (err) => {
      Taro.showToast({
        title: err.message,
        icon: 'error',
      })
    }
  });

  useEffect(() => {
    return () => {
      if (timerRef.current) {
        clearTimeout(timerRef.current)
        timerRef.current = null
      }
    }
  }, [])

  const handlePhoneNumberChange = (e) => {
    setPhoneNumber(e.detail.value)
  }

  const handlePasswordChange = (e) => {
    setPassword(e.detail.value)
  }

  const handleLogin = () => {
    loginRun({ phone: phoneNumber, password })
  }

  return {
    phoneNumber,
    password,
    loading,
    handlePhoneNumberChange,
    handlePasswordChange,
    handleLogin,
  }
}
