import {View, Button} from '@tarojs/components'
import useHome from './useHome'
import Taro from "@tarojs/taro";

export default function Home() {

  const { } = useHome()

  const handleLoginClick = () => {
    Taro.navigateTo({
      url: "/pages/login/index",
    })
  }

  return (
    <View className="w-full h-full flex justify-center items-center">
      <Button className="bg-primary text-white px-5" onClick={handleLoginClick}>去登陆</Button>
    </View>
  )
}
