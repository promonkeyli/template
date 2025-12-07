import { View, Text, Input, Button, Image } from '@tarojs/components'
// import useIndex from './useIndex'

export default function Index() {

  // const {
  //   username,
  //   setUsername,
  //   password,
  //   setPassword,
  //   loading,
  //   userInfo,
  //   handleLogin,
  //   handleLogout
  // } = useIndex()

  return (
    <View style={{ flex: 1 }}>1212</View>
  )
  // return (
  //   <View className='h-screen w-full bg-gray-100 flex flex-col items-center justify-center p-4'>
  //     <View className='bg-white p-8 rounded-xl shadow-lg w-full'>
  //       <View className='mb-8 text-center'>
  //         <Text className='text-3xl font-bold text-gray-800'>Welcome Back</Text>
  //         <View className='mt-2'>
  //           <Text className='text-gray-500 text-sm'>Taro + Tailwind + Zustand Demo</Text>
  //         </View>
  //       </View>
  //
  //       {userInfo ? (
  //         <View className='flex flex-col items-center'>
  //           <View className='w-20 h-20 bg-blue-100 rounded-full flex items-center justify-center mb-4 text-4xl'>
  //             {userInfo.avatar ? (
  //               <Image src={userInfo.avatar} className='w-full h-full rounded-full' />
  //             ) : (
  //               <Text>👤</Text>
  //             )}
  //           </View>
  //           <Text className='text-xl font-semibold text-gray-800 mb-2'>
  //             {userInfo.nickname || userInfo.username}
  //           </Text>
  //           <Text className='text-gray-500 mb-8 text-sm'>
  //             已登录
  //           </Text>
  //
  //           <Button
  //             className='w-full bg-red-500 active:bg-red-600 text-white font-bold py-3 rounded-lg border-none'
  //             onClick={handleLogout}
  //           >
  //             退出登录
  //           </Button>
  //         </View>
  //       ) : (
  //         <View>
  //           <View className='mb-4 w-full flex flex-col'>
  //             <Text className='text-gray-700 text-sm font-bold mb-2'>用户名</Text>
  //             <Input
  //               className='bg-white border border-gray-200 rounded-md py-2 text-gray-900 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 block w-full transition-all duration-200'
  //               placeholder='请输入用户名'
  //               value={username}
  //               onInput={(e) => setUsername(e.detail.value)}
  //             />
  //           </View>
  //
  //           <View className='mb-6 w-full flex flex-col'>
  //             <Text className='text-gray-700 text-sm font-bold mb-2'>密码</Text>
  //             <Input
  //               className='bg-white border placeholder:mx-2 border-gray-200 rounded-md py-2 text-gray-900 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 block w-full transition-all duration-200'
  //               password
  //               placeholder='请输入密码'
  //               value={password}
  //               onInput={(e) => setPassword(e.detail.value)}
  //             />
  //           </View>
  //
  //           <Button
  //             className={`w-full h-12 text-white font-bold py-1 rounded-lg border-none flex items-center justify-center ${loading ? 'bg-blue-300' : 'bg-blue-600 active:bg-blue-700'
  //               }`}
  //             onClick={handleLogin}
  //             disabled={loading}
  //           >
  //             {loading ? '登录中...' : '立即登录'}
  //           </Button>
  //         </View>
  //       )}
  //     </View>
  //
  //     <View className='mt-8 text-center'>
  //       <Text className='text-gray-400 text-xs'>Powered by Antigravity</Text>
  //     </View>
  //   </View>
  // )
}
