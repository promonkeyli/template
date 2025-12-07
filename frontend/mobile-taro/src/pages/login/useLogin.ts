import Taro from '@tarojs/taro'
import { useState } from 'react'

export const useLogin = () => {
  const [phoneNumber, setPhoneNumber] = useState('')
  const [password, setPassword] = useState('')

  const handleGetPhoneNumber = (e) => {
    if (e.detail.errMsg === 'getPhoneNumber:ok') {
      // 用户同意授权
      console.log('用户同意授权', e.detail)
      // 在这里调用后端接口，发送 code 和 encryptedData, iv
      // Taro.login({
      //   success: (res) => {
      //     if (res.code) {
      //       // 发送 res.code 到后台换取 openId, sessionKey, unionId
      //       // 再将 e.detail.encryptedData, e.detail.iv 发送给后端解密
      //     }
      //   }
      // })
    } else {
      // 用户拒绝授权
      console.log('用户拒绝授权')
    }
  }

  const handlePhoneNumberChange = (e) => {
    setPhoneNumber(e.detail.value)
  }

  const handlePasswordChange = (e) => {
    setPassword(e.detail.value)
  }

  const handleLogin = () => {
    console.log('Login attempt with:', { phoneNumber, password })
    // Here you would typically call an API to authenticate the user
    // For now, just logging the values
  }

  return {
    phoneNumber,
    password,
    handlePhoneNumberChange,
    handlePasswordChange,
    handleLogin,
    handleGetPhoneNumber,
  }
}
