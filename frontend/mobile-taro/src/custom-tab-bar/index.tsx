import Taro from '@tarojs/taro'
import { CoverView, CoverImage } from '@tarojs/components'
import { useMemo } from 'react'

import useTabBarStore from '@/stores/tab-bar'
import './index.css'

const CustomTabBar = () => {
  const list = useMemo(() => [
    {
      pagePath: 'pages/index/index',
      text: '首页',
      iconPath: '@/assets/tabbar/home.svg',
      selectedIconPath: '@/assets/tabbar/home_selected.svg'
    },
    {
      pagePath: 'pages/category/index',
      text: '分类',
      iconPath: '@/assets/tabbar/category.svg',
      selectedIconPath: '@/assets/tabbar/category_selected.svg'
    },
    {
      pagePath: 'pages/cart/index',
      text: '购物车',
      iconPath: '@/assets/tabbar/cart.svg',
      selectedIconPath: '@/assets/tabbar/cart_selected.svg'
    },
    {
      pagePath: 'pages/my/index',
      text: '我的',
      iconPath: '@/assets/tabbar/me.svg',
      selectedIconPath: '@/assets/tabbar/me_selected.svg'
    }
  ], [])

  // 使用全局状态管理选中项
  const { selected, setSelected } = useTabBarStore()

  console.log("custom tabBar selected")

  const switchTab = (item, index) => {
    const url = '/' + item.pagePath
    Taro.switchTab({ url })
    setSelected(index)
  }

  return (
    <CoverView className='tab-bar'>
      {list.map((item, index) => {
        const isSelected = selected === index
        return (
          <CoverView key={item.text} className='tab-bar-item' onClick={() => switchTab(item, index)}>
            <CoverImage
              className='tab-bar-icon'
              src={isSelected ? item.selectedIconPath : item.iconPath}
            />
            <CoverView className={`tab-bar-text ${isSelected ? 'tab-bar-text-selected' : ''}`}>
              {item.text}
            </CoverView>
          </CoverView>
        )
      })}
    </CoverView>
  )
}

export default CustomTabBar


