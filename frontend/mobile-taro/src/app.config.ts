export default defineAppConfig({
  pages: [
    'pages/home/index',
    'pages/category/index',
    'pages/cart/index',
    'pages/my/index',
    'pages/login/index'
  ],
  window: {
    backgroundTextStyle: 'light',
    navigationBarBackgroundColor: '#fff',
    navigationBarTitleText: 'WeChat',
    navigationBarTextStyle: 'black'
  },
  tabBar: {
    // custom: true, // vite 版本的自定义tabbar 编译现在bug，等待后续官方修复
    color: '#000000',
    selectedColor: '#9BAA6A',
    backgroundColor: '#ffffff',
    list: [
      {
        pagePath: 'pages/home/index',
        text: '首页',
        iconPath: 'assets/tabbar/png/home.png',
        selectedIconPath: 'assets/tabbar/png/home_selected.png',
      },
      {
        pagePath: 'pages/category/index',
        text: '分类',
        iconPath: 'assets/tabbar/png/category.png',
        selectedIconPath: 'assets/tabbar/png/category_selected.png',
      },
      {
        pagePath: 'pages/cart/index',
        text: '购物车',
        iconPath: 'assets/tabbar/png/cart.png',
        selectedIconPath: 'assets/tabbar/png/cart_selected.png',
      },
      {
        pagePath: 'pages/my/index',
        text: '我的',
        iconPath: 'assets/tabbar/png/my.png',
        selectedIconPath: 'assets/tabbar/png/my_selected.png',
      }
    ]
  }
})
