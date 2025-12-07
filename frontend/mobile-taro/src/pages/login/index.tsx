import { View, Button, Input, Image } from '@tarojs/components'
import { useLogin } from './useLogin'
import logo from '@/assets/images/logo.jpg'

const Login = () => {
  const {
    phoneNumber,
    password,
    handlePhoneNumberChange,
    handlePasswordChange,
    handleLogin,
    handleGetPhoneNumber,
  } = useLogin()

  return (
    <View className="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-4">
      <View className="mb-8">
        <Image src={logo} className="w-24 h-24 rounded-full" />
      </View>

      <View className="w-full max-w-sm bg-white p-6 rounded-lg shadow-md">
        <View className="text-2xl font-bold text-center mb-6">商城登录</View>

        <View className="mb-4">
          <Input
            className="w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            type="number"
            placeholder="请输入手机号"
            value={phoneNumber}
            onInput={handlePhoneNumberChange}
            maxLength={11}
          />
        </View>

        <View className="mb-6">
          <Input
            className="w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            type="password"
            placeholder="请输入密码"
            value={password}
            onInput={handlePasswordChange}
          />
        </View>

        <Button
          className="w-full bg-blue-500 text-white p-3 rounded-md text-lg hover:bg-blue-600 transition-colors duration-200"
          onClick={handleLogin}
        >
          登录
        </Button>

        <Button
          openType="getPhoneNumber"
          onGetPhoneNumber={handleGetPhoneNumber}
          className="w-full bg-green-500 text-white p-3 rounded-md text-lg mt-4 hover:bg-green-600 transition-colors duration-200"
        >
          微信手机号一键登录
        </Button>
      </View>
    </View>
  )
}

export default Login
