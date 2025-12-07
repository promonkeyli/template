import { View, Button, Input, Image } from '@tarojs/components'
import { useLogin } from './useLogin'
import logo from '@/assets/images/logo.jpg'

const Login = () => {

  const {
    phoneNumber,
    password,
    loading,
    handlePhoneNumberChange,
    handlePasswordChange,
    handleLogin,
  } = useLogin()

  return (
    <View className="flex flex-col items-center justify-center px-8 pt-20">
      <View className="my-8">
        <Image src={logo} className="w-24 h-24 rounded-full" />
      </View>

      <View className="w-full bg-white px-5">

        <View className="mb-4">
          <Input
            className="w-full py-1 border border-gray-300 focus:outline-none focus:ring-2"
            type="number"
            placeholder="请输入手机号"
            value={phoneNumber}
            onInput={handlePhoneNumberChange}
            maxlength={11}
          />
        </View>

        <View className="mb-4">
          <Input
            type="safe-password"
            password
            className="w-full py-1 border border-gray-300 focus:outline-none focus:ring-2"
            placeholder="请输入密码"
            value={password}
            onInput={handlePasswordChange}
          />
        </View>

        <Button
          className="w-full bg-primary text-white py-2 rounded-sm text-base"
          disabled={loading}
          onClick={handleLogin}
        >
          登录
        </Button>
      </View>
    </View>
  )
}

export default Login
